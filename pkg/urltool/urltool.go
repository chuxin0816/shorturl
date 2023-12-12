package urltool

import (
	"net/url"
	"path"
	"strings"
)

func GetBaseURL(rawURL string) (string, error) {
	rawURL = strings.TrimPrefix(rawURL, "http://")
	rawURL = strings.TrimPrefix(rawURL, "https://")
	myUrl, err := url.Parse(rawURL)
	if err != nil {
		return "", err
	}
	return path.Base(myUrl.Path), nil
}
