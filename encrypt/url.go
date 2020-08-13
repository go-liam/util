package encrypt

import (
	"net/url"
	"strings"
)

// EncodeURIComponent :
func EncodeURIComponent(str string) string {
	r := url.QueryEscape(str)
	r = strings.Replace(r, "+", "%20", -1)
	return r
}
