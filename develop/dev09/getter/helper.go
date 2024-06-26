package getter

import "strings"

type Helper struct {
}

func NewHelper() *Helper {
	return new(Helper)
}

func (receiver *Helper) AddUrlSuffix(url string) string {
	if strings.HasSuffix(url, "/") {
		return url
	}
	return url + "/"
}

func (receiver *Helper) ConvertPreviousLink(link string, currentLevel int) string {
	if currentLevel > 0 {
		backLink := strings.Repeat("../", currentLevel)
		return backLink + link
	}
	return link
}

func (receiver *Helper) ReplaceUrlToPath(url string) string {
	if strings.HasPrefix(url, "https://") {
		return strings.TrimPrefix(url, "https://")
	}
	return strings.TrimPrefix(url, "http://")
}

func (receiver *Helper) ModifyUrl(url, urlWithSuffix string) string {
	if strings.HasPrefix(url, "http") {
		return url
	}

	url = strings.TrimLeft(url, "./")
	return urlWithSuffix + url
}
