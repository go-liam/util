package tshttp

import (
	"strings"
	"testing"
)

func TestHTTPIndex(t *testing.T) {
	sv := new(HTTPTest)
	url := "http://baidu.com"
	data := `
		{
			"ID":1,
			"versionId":101
		}
		`
	reader := strings.NewReader(data)
	code, back, _ := sv.HTTPData(MethodHTTPGET, url, reader, "", "")
	println("code:", code)
	println(back)
}
