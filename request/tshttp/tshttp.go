package tshttp

import (
	"errors"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

// MethodHTTPGET :
const (
	MethodHTTPGET     = "GET"
	MethodHTTPPOST    = "POST"
	MethodHTTPPATCH   = "PATCH"
	MethodHTTPOPTIONS = "OPTIONS"
	MethodHTTPPUT     = "PUT"
	MethodHTTPDELETE  = "DELETE"
)

// HTTPTest :
type HTTPTest struct {
	TokenKey     string // Token
	AppSecretKey string // AppSecret
}

// HTTPData :
func (sv *HTTPTest) HTTPData(method string, url string, body io.Reader, utoken string, appSecret string) (int, string, error) {
	if sv.TokenKey == "" {
		sv.TokenKey = "Token"
	}
	if sv.AppSecretKey == "" {
		sv.AppSecretKey = "AppSecret"
	}
	return sv.httpData(method, url, body, utoken, appSecret)
}

// httpData :测试工具
func (sv *HTTPTest) httpData(method string, url string, body io.Reader, token string, appSecret string) (int, string, error) {
	req, _ := http.NewRequest(method, url, body)
	req.Header.Add("accept", "application/json")
	req.Header.Add("content-type", "application/json")
	req.Header.Add(sv.TokenKey, token)
	req.Header.Add(sv.AppSecretKey, appSecret)
	res, err := (&http.Client{}).Do(req)
	if err != nil {
		return 0, "", fmt.Errorf("[ERROR]send failed, error: %v", err.Error())
	}
	defer func() { _ = res.Body.Close() }()
	result, err := ioutil.ReadAll(res.Body)
	if res.StatusCode != 200 {
		log.Printf("[ERROR] RouteTestTool %v--%v \n", res.StatusCode, url)
		return res.StatusCode, "", errors.New("httpStatusFail")
	}
	return res.StatusCode, string(result), nil
}
