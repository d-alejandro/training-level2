package getter

import (
	"bytes"
	"fmt"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

const TagA = "a"
const TagIMG = "img"
const TagSCRIPT = "script"
const TagLINK = "link"

type WebGetter struct {
	levelMaxFlag     int
	fileWriter       *FileWriter
	helper           *Helper
	urlWithSuffix    string
	urlWithoutSuffix string
	currentUrl       string
	currentLevel     int
	rootPath         string
	linkMap          map[string]string
	linkSavedMap     map[string]struct{}
	resourceMap      map[string]string
}

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

func (receiver *WebGetter) Execute(url string) error {
	receiver.urlWithoutSuffix = strings.TrimSuffix(url, "/")
	receiver.urlWithSuffix = receiver.helper.AddUrlSuffix(url)

	receiver.linkMap[receiver.urlWithSuffix] = ""

	for receiver.currentLevel = 0; receiver.currentLevel < receiver.levelMaxFlag; receiver.currentLevel++ {
		fmt.Printf("level %d: start saving web content\n", receiver.currentLevel+1)

		linkMap := receiver.linkMap
		receiver.linkMap = make(map[string]string)

		for link, path := range linkMap {
			if _, isExist := receiver.linkSavedMap[link]; isExist {
				continue
			}

			receiver.currentUrl = link

			buffer, err := receiver.getResponseAndProcessContent()
			if err != nil {
				return err
			}

			if receiver.currentLevel == 0 {
				path = receiver.rootPath
			}

			receiver.fileWriter.WriteContent(path, buffer.String())

			receiver.linkSavedMap[receiver.currentUrl] = struct{}{}
		}

		fmt.Printf("level %d: saving web content completed\n", receiver.currentLevel+1)
	}

	fmt.Println("start loading web resources")

	for resourceUrl, resourcePath := range receiver.resourceMap {
		receiver.fileWriter.WriteResourceFile(resourceUrl, resourcePath)
	}

	fmt.Println("loading of web resources completed")
	fmt.Println("start loading CSS web resources")

	receiver.fileWriter.WriteCSSResources()

	fmt.Println("loading of CSS web resources completed")

	return nil
}

func (receiver *WebGetter) getResponseAndProcessContent() (*bytes.Buffer, error) {
	response, errGet := http.Get(receiver.currentUrl)
	if errGet != nil {
		return nil, errGet
	}

	if receiver.currentLevel == 0 {
		receiver.rootPath = receiver.helper.AddUrlSuffix(response.Request.URL.Host)
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
		receiver.processHtmlElementNode(node)
	}
	for child := node.FirstChild; child != nil; child = child.NextSibling {
		receiver.processNodes(child)
	}
}

func (receiver *WebGetter) processHtmlElementNode(node *html.Node) {
	switch node.Data {
	case TagA:
		for key, attribute := range node.Attr {
			if attribute.Key == "href" && strings.HasPrefix(attribute.Val, receiver.urlWithoutSuffix) {
				attributeValue := receiver.helper.AddUrlSuffix(attribute.Val)
				attributeValueTrimmed := strings.TrimPrefix(attributeValue, receiver.urlWithSuffix)

				attribute.Val = receiver.helper.ConvertPreviousLink(
					attributeValueTrimmed+"index.html",
					receiver.currentLevel,
				)
				node.Attr[key] = attribute

				if _, isExist := receiver.linkSavedMap[attributeValue]; !isExist {
					receiver.linkMap[attributeValue] = receiver.rootPath + attributeValueTrimmed
				}
				break
			}
		}
	case TagIMG:
		for key, attribute := range node.Attr {
			if attribute.Key == "src" {
				value := receiver.helper.ReplaceUrlToPath(attribute.Val)
				receiver.resourceMap[attribute.Val] = receiver.rootPath + value

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
			isCss          bool
		)

		for index, attribute := range node.Attr {
			if attribute.Key == "rel" {
				if attribute.Val == "stylesheet" ||
					attribute.Val == "icon" ||
					attribute.Val == "apple-touch-icon" ||
					attribute.Val == "EditURI" {
					relKeyExist = true
					isCss = attribute.Val == "stylesheet"
					continue
				}
				break
			} else if attribute.Key == "href" {
				attributeIndex, htmlAttribute = index, attribute
				hrefKeyExist = true
			}
		}

		if relKeyExist && hrefKeyExist {
			if isCss && !strings.HasSuffix(htmlAttribute.Val, ".css") {
				htmlAttribute.Val = htmlAttribute.Val + ".css"
			}
			receiver.processAttributeValue(attributeIndex, htmlAttribute, node)
		}
	}
}

func (receiver *WebGetter) processAttributeValue(attributeIndex int, attribute html.Attribute, node *html.Node) {
	value := receiver.helper.ReplaceUrlToPath(attribute.Val)
	value = strings.TrimPrefix(value, receiver.rootPath)
	value = strings.TrimLeft(value, "/")

	modifiedUrl := receiver.helper.ModifyUrl(attribute.Val, receiver.urlWithSuffix)
	receiver.resourceMap[modifiedUrl] = receiver.rootPath + value

	attribute.Val = receiver.helper.ConvertPreviousLink(value, receiver.currentLevel)
	attribute.Val = strings.ReplaceAll(attribute.Val, "%", "%25")
	attribute.Val = strings.ReplaceAll(attribute.Val, "?", "%3F")
	node.Attr[attributeIndex] = attribute
}
