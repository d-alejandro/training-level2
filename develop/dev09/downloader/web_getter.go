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
	rootUrl      string
	currentLevel int
	linkMap      map[string]struct{}
}

func NewWebGetter(levelMaxFlag int) *WebGetter {
	linkMap := make(map[string]struct{})

	return &WebGetter{
		levelMaxFlag: levelMaxFlag,
		fileWriter:   NewFileWriter(),
		linkMap:      linkMap,
	}
}

func (receiver *WebGetter) Get(url string) error {
	receiver.rootUrl = receiver.addUrlSuffix(url)

	response, errGet := http.Get(receiver.rootUrl)
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

	host := response.Request.URL.Host
	receiver.fileWriter.WriteContent(host, buffer.String())

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

				if strings.HasPrefix(attribute.Val, receiver.rootUrl) {
					attribute.Val = strings.TrimPrefix(attribute.Val, receiver.rootUrl) + "index.html"
					node.Attr[key] = attribute

					receiver.linkMap[attribute.Val] = struct{}{}
				}
				break
			}
		}
	}
}
