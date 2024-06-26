package getter

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

/*
TagA tag <a>
*/
const TagA = "a"

/*
TagIMG tag <img>
*/
const TagIMG = "img"

/*
TagSCRIPT tag <script>
*/
const TagSCRIPT = "script"

/*
TagLINK tag <link>
*/
const TagLINK = "link"

/*
WebGetter structure
*/
type WebGetter struct {
	levelMaxFlag     int
	fileWriter       *FileWriter
	helper           *Helper
	urlWithSuffix    string
	urlWithoutSuffix string
	currentURL       string
	currentLevel     int
	rootPath         string
	linkMap          map[string]string
	linkSavedMap     map[string]struct{}
	resourceMap      map[string]string
}

/*
NewWebGetter constructor
*/
func NewWebGetter(levelMaxFlag int) *WebGetter {
	linkMap := make(map[string]string)
	linkSavedMap := make(map[string]struct{})
	imageMap := make(map[string]string)

	return &WebGetter{
		levelMaxFlag: levelMaxFlag,
		fileWriter:   NewFileWriter(),
		helper:       NewHelper(),
		linkMap:      linkMap,
		linkSavedMap: linkSavedMap,
		resourceMap:  imageMap,
	}
}

/*
Execute method
*/
func (receiver *WebGetter) Execute(url string) error {
	receiver.urlWithoutSuffix = strings.TrimSuffix(url, "/")
	receiver.urlWithSuffix = receiver.helper.AddURLSuffix(url)

	receiver.linkMap[receiver.urlWithSuffix] = ""

	for receiver.currentLevel = 0; receiver.currentLevel < receiver.levelMaxFlag; receiver.currentLevel++ {
		fmt.Printf("level %d: start saving web content\n", receiver.currentLevel+1)

		linkMap := receiver.linkMap
		receiver.linkMap = make(map[string]string)

		for link, path := range linkMap {
			if _, isExist := receiver.linkSavedMap[link]; isExist {
				continue
			}

			receiver.currentURL = link

			buffer, err := receiver.getResponseAndProcessContent()
			if err != nil {
				return err
			}

			if receiver.currentLevel == 0 {
				path = receiver.rootPath
			}

			receiver.fileWriter.WriteContent(path, buffer.String())

			receiver.linkSavedMap[receiver.currentURL] = struct{}{}
		}

		fmt.Printf("level %d: saving web content completed\n", receiver.currentLevel+1)
	}

	fmt.Println("start loading web resources")

	for resourceURL, resourcePath := range receiver.resourceMap {
		receiver.fileWriter.WriteResourceFile(resourceURL, resourcePath)
	}

	fmt.Println("loading of web resources completed")
	fmt.Println("start loading CSS web resources")

	receiver.fileWriter.WriteCSSResources()

	fmt.Println("loading of CSS web resources completed")

	return nil
}

func (receiver *WebGetter) getResponseAndProcessContent() (*bytes.Buffer, error) {
	response, errGet := http.Get(receiver.currentURL)
	if errGet != nil {
		return nil, errGet
	}

	if receiver.currentLevel == 0 {
		receiver.rootPath = receiver.helper.AddURLSuffix(response.Request.URL.Host)
	}

	node, errParse := html.Parse(response.Body)
	if errParse != nil {
		return nil, errParse
	}

	receiver.processNodes(node)

	buffer := &bytes.Buffer{}
	if err := html.Render(buffer, node); err != nil {
		return nil, err
	}

	if err := response.Body.Close(); err != nil {
		return nil, err
	}

	return buffer, nil
}

func (receiver *WebGetter) processNodes(node *html.Node) {
	if node.Type == html.ElementNode {
		receiver.processHTMLElementNode(node)
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		receiver.processNodes(child)
	}
}

func (receiver *WebGetter) processHTMLElementNode(node *html.Node) {
	switch node.Data {
	case TagA:
		for key, attribute := range node.Attr {
			if attribute.Key == "href" {
				if strings.HasPrefix(attribute.Val, receiver.urlWithoutSuffix) {
					attributeValue := receiver.helper.AddURLSuffix(attribute.Val)
					attributeValueTrimmed := strings.TrimPrefix(attributeValue, receiver.urlWithSuffix)

					attribute.Val = receiver.helper.ConvertPreviousLink(
						attributeValueTrimmed+"index.html",
						receiver.currentLevel,
					)
					node.Attr[key] = attribute

					if _, isExist := receiver.linkSavedMap[attributeValue]; !isExist {
						receiver.linkMap[attributeValue] = receiver.rootPath + attributeValueTrimmed
					}
				} else if strings.HasPrefix(attribute.Val, ".") {
					attributeValue := receiver.helper.AddURLSuffix(attribute.Val)
					attributeValue = strings.TrimLeft(attributeValue, "./")

					attribute.Val = receiver.helper.ConvertPreviousLink(
						attributeValue+"index.html",
						receiver.currentLevel,
					)
					node.Attr[key] = attribute

					if _, isExist := receiver.linkSavedMap[attributeValue]; !isExist {
						receiver.linkMap[receiver.urlWithSuffix+attributeValue] = receiver.rootPath + attributeValue
					}
				}
				break
			}
		}
	case TagIMG:
		for key, attribute := range node.Attr {
			if attribute.Key == "src" {
				value := receiver.helper.ReplaceURLToPath(attribute.Val)
				value = strings.TrimPrefix(value, receiver.rootPath)
				value = strings.TrimLeft(value, "./")

				modifiedURL := receiver.helper.ModifyURL(attribute.Val, receiver.urlWithSuffix)
				receiver.resourceMap[modifiedURL] = receiver.rootPath + value

				attribute.Val = receiver.helper.ConvertPreviousLink(value, receiver.currentLevel)
				node.Attr[key] = attribute

				break
			}
		}
	case TagSCRIPT:
		for key, attribute := range node.Attr {
			if attribute.Key == "src" {
				receiver.processAttributeValue(key, attribute, node)
				break
			}
		}
	case TagLINK:
		var (
			attributeIndex int
			htmlAttribute  html.Attribute
			relKeyExist    bool
			hrefKeyExist   bool
			isCSS          bool
		)

		for index, attribute := range node.Attr {
			if attribute.Key == "rel" {
				if attribute.Val == "stylesheet" ||
					attribute.Val == "icon" ||
					attribute.Val == "apple-touch-icon" ||
					attribute.Val == "EditURI" {
					relKeyExist = true
					isCSS = attribute.Val == "stylesheet"
					continue
				}
				break
			} else if attribute.Key == "href" {
				attributeIndex, htmlAttribute = index, attribute
				hrefKeyExist = true
			}
		}

		if relKeyExist && hrefKeyExist {
			if isCSS && !strings.HasSuffix(htmlAttribute.Val, ".css") {
				htmlAttribute.Val = htmlAttribute.Val + ".css"
			}
			receiver.processAttributeValue(attributeIndex, htmlAttribute, node)
		}
	}
}

func (receiver *WebGetter) processAttributeValue(attributeIndex int, attribute html.Attribute, node *html.Node) {
	value := receiver.helper.ReplaceURLToPath(attribute.Val)
	value = strings.TrimPrefix(value, receiver.rootPath)
	value = strings.TrimLeft(value, "./")

	modifiedURL := receiver.helper.ModifyURL(attribute.Val, receiver.urlWithSuffix)
	receiver.resourceMap[modifiedURL] = receiver.rootPath + value

	attribute.Val = receiver.helper.ConvertPreviousLink(value, receiver.currentLevel)
	attribute.Val = strings.ReplaceAll(attribute.Val, "%", "%25")
	attribute.Val = strings.ReplaceAll(attribute.Val, "?", "%3F")
	node.Attr[attributeIndex] = attribute
}
