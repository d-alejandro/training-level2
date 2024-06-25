package getter

import (
	"bytes"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

const TagA = "a"
const TagIMG = "img"
const TagSCRIPT = "script"

type WebGetter struct {
	levelMaxFlag     int
	fileWriter       *FileWriter
	urlWithSuffix    string
	urlWithoutSuffix string
	currentUrl       string
	currentPath      string
	currentLevel     int
	rootPatch        string
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
		linkMap:      linkMap,
		linkSavedMap: linkSavedMap,
		resourceMap:  imageMap,
	}
}

func (receiver *WebGetter) Execute(url string) error {
	receiver.urlWithoutSuffix = strings.TrimSuffix(url, "/")

	receiver.urlWithSuffix = receiver.addUrlSuffix(url)
	receiver.linkMap[receiver.urlWithSuffix] = ""

	for receiver.currentLevel = 0; receiver.currentLevel < receiver.levelMaxFlag; receiver.currentLevel++ {
		linkMap := receiver.linkMap
		receiver.linkMap = make(map[string]string)

		for link, path := range linkMap {
			if _, isExist := receiver.linkSavedMap[link]; isExist {
				continue
			}

			receiver.currentUrl = link
			receiver.currentPath = path

			if err := receiver.getResponseProcessAndSaveContent(); err != nil {
				return err
			}

			receiver.linkSavedMap[receiver.currentUrl] = struct{}{}
		}
	}

	for resourceUrl, resourcePath := range receiver.resourceMap {
		receiver.fileWriter.WriteResourceFile(resourceUrl, receiver.rootPatch+resourcePath)
	}

	return nil
}

func (receiver *WebGetter) getResponseProcessAndSaveContent() error {
	response, errGet := http.Get(receiver.currentUrl)
	if errGet != nil {
		return errGet
	}

	if receiver.currentLevel == 0 {
		receiver.rootPatch = receiver.addUrlSuffix(response.Request.URL.Host)
	}

	node, errParse := html.Parse(response.Body)
	if errParse != nil {
		return errParse
	}

	receiver.processNodes(node)

	buffer := &bytes.Buffer{}
	if err := html.Render(buffer, node); err != nil {
		return err
	}

	patch := receiver.rootPatch + receiver.currentPath
	receiver.fileWriter.WriteContent(patch, buffer.String())

	if err := response.Body.Close(); err != nil {
		return err
	}

	return nil
}

func (receiver *WebGetter) addUrlSuffix(url string) string {
	if strings.HasSuffix(url, "/") {
		return url
	}
	return url + "/"
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
				attributeValue := receiver.addUrlSuffix(attribute.Val)
				attributeValueTrimmed := strings.TrimPrefix(attributeValue, receiver.urlWithSuffix)

				attribute.Val = receiver.convertPreviousLink(attributeValueTrimmed + "index.html")
				node.Attr[key] = attribute

				if _, isExist := receiver.linkSavedMap[attributeValue]; !isExist {
					receiver.linkMap[attributeValue] = attributeValueTrimmed
				}
				break
			}
		}
	case TagIMG:
		for key, attribute := range node.Attr {
			if attribute.Key == "src" {
				value := receiver.replaceUrlToPath(attribute.Val)
				receiver.resourceMap[attribute.Val] = value

				attribute.Val = receiver.convertPreviousLink(value)
				node.Attr[key] = attribute

				break
			}
		}
	case TagSCRIPT:
		for key, attribute := range node.Attr {
			if attribute.Key == "src" {
				value := receiver.replaceUrlToPath(attribute.Val)
				value = receiver.removeRootPatch(value)
				value = strings.TrimPrefix(value, "/")

				modifiedUrl := receiver.modifyUrl(attribute.Val)
				receiver.resourceMap[modifiedUrl] = value

				attribute.Val = receiver.convertPreviousLink(value)
				node.Attr[key] = attribute

				break
			}
		}
	}
}

func (receiver *WebGetter) convertPreviousLink(link string) string {
	if receiver.currentLevel > 0 {
		backLink := strings.Repeat("../", receiver.currentLevel)
		return backLink + link
	}
	return link
}

func (receiver *WebGetter) replaceUrlToPath(url string) string {
	if strings.HasPrefix(url, "https://") {
		return strings.TrimPrefix(url, "https://")
	}
	return strings.TrimPrefix(url, "http://")
}

func (receiver *WebGetter) removeRootPatch(url string) string {
	return strings.TrimPrefix(url, receiver.rootPatch)
}

func (receiver *WebGetter) modifyUrl(url string) string {
	if strings.HasPrefix(url, "http") {
		return url
	}

	url = strings.TrimPrefix(url, "/")
	return receiver.urlWithSuffix + url
}
