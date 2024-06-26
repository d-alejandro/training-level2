package getter

import "strings"

/*
Helper structure
*/
type Helper struct {
}

/*
NewHelper constructor
*/
func NewHelper() *Helper {
	return new(Helper)
}

/*
AddURLSuffix method
*/
func (receiver *Helper) AddURLSuffix(url string) string {
	if strings.HasSuffix(url, "/") {
		return url
	}
	return url + "/"
}

/*
ConvertPreviousLink method
*/
func (receiver *Helper) ConvertPreviousLink(link string, currentLevel int) string {
	if currentLevel > 0 {
		backLink := strings.Repeat("../", currentLevel)
		return backLink + link
	}
	return link
}

/*
ReplaceURLToPath method
*/
func (receiver *Helper) ReplaceURLToPath(url string) string {
	if strings.HasPrefix(url, "https://") {
		return strings.TrimPrefix(url, "https://")
	}
	return strings.TrimPrefix(url, "http://")
}

/*
ModifyURL method
*/
func (receiver *Helper) ModifyURL(url, urlWithSuffix string) string {
	if strings.HasPrefix(url, "http") {
		return url
	}

	url = strings.TrimLeft(url, "./")
	return urlWithSuffix + url
}
