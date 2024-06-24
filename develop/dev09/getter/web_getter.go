package getter

import (
	"bytes"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

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
	linkSavedMap     map[string]string
}

func NewWebGetter(levelMaxFlag int) *WebGetter {
	linkMap := make(map[string]string)
	linkSavedMap := make(map[string]string)

	return &WebGetter{
		levelMaxFlag: levelMaxFlag,
		fileWriter:   NewFileWriter(),
		linkMap:      linkMap,
		linkSavedMap: linkSavedMap,
	}
}

func (receiver *WebGetter) Execute(url string) error {
	receiver.urlWithoutSuffix = strings.TrimSuffix(url, "/")

	receiver.urlWithSuffix = receiver.addUrlSuffix(url)
	receiver.linkMap[receiver.urlWithSuffix] = ""

	for receiver.currentLevel = 0; receiver.currentLevel <= 2; receiver.currentLevel++ {
		linkMap := receiver.linkMap
		receiver.linkMap = make(map[string]string)

		for link, path := range linkMap {
			if _, isExist := receiver.linkSavedMap[link]; isExist {
				continue
			}

			receiver.currentUrl = link
			receiver.currentPath = path

			if err := receiver.get(); err != nil {
				return err
			}

			receiver.linkSavedMap[receiver.currentUrl] = receiver.currentPath
		}
	}

	return nil
}

func (receiver *WebGetter) get() error {
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
	const TagA = "a"

	switch node.Data {
	case TagA:
		for key, attribute := range node.Attr {
			if attribute.Key == "href" && strings.HasPrefix(attribute.Val, receiver.urlWithoutSuffix) {
				attributeValue := receiver.addUrlSuffix(attribute.Val)
				attributeValueTrimmed := strings.TrimPrefix(attributeValue, receiver.urlWithSuffix)

				attribute.Val = attributeValueTrimmed + "index.html"
				node.Attr[key] = attribute

				if _, isExist := receiver.linkSavedMap[attributeValue]; !isExist {
					receiver.linkMap[attributeValue] = attributeValueTrimmed
				}
				break
			}
		}
	}
}