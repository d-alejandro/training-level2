package downloader

import (
	"bytes"
	"golang.org/x/net/html"
	"net/http"
	"strings"
)

type WebGetter struct {
	levelMaxFlag int
	fileWriter   *FileWriter
	host         string
	currentUrl   string
	currentPath  string
	currentLevel int
	linkMap      map[string]string
}

func NewWebGetter(levelMaxFlag int) *WebGetter {
	linkMap := make(map[string]string)

	return &WebGetter{
		levelMaxFlag: levelMaxFlag,
		fileWriter:   NewFileWriter(),
		linkMap:      linkMap,
	}
}

func (receiver *WebGetter) Execute(url string) error {
	receiver.currentUrl = receiver.addUrlSuffix(url)

	if err := receiver.get(); err != nil {
		return err
	}

	for receiver.currentLevel = 1; receiver.currentLevel <= receiver.levelMaxFlag; receiver.currentLevel++ {
		linkMap := receiver.linkMap
		receiver.linkMap = make(map[string]string)

		for link, path := range linkMap {
			receiver.currentUrl = link
			receiver.currentPath = path

			if err := receiver.get(); err != nil {
				return err
			}
		}
	}

	return nil
}

func (receiver *WebGetter) get() error {
	response, errGet := http.Get(receiver.currentUrl)
	if errGet != nil {
		return errGet
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

	if receiver.currentLevel == 0 {
		receiver.host = response.Request.URL.Host
	}

	path := receiver.host + "/" + receiver.currentPath

	receiver.fileWriter.WriteContent(path, buffer.String())

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
	case "a":
		for key, attribute := range node.Attr {
			if attribute.Key == "href" {
				attribute.Val = receiver.addUrlSuffix(attribute.Val)

				if strings.HasPrefix(attribute.Val, receiver.currentUrl) {
					attributeValue := attribute.Val
					attributeValueTrimmed := strings.TrimPrefix(attributeValue, receiver.currentUrl)

					attribute.Val = attributeValueTrimmed + "index.html"
					node.Attr[key] = attribute

					if attributeValue != receiver.currentUrl {
						receiver.linkMap[attributeValue] = attributeValueTrimmed
					}
				}
				break
			}
		}
	}
}
