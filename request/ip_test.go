package request

import (
	"net/http"
	"testing"
)

func TestClientIP(t *testing.T) {
	r := new(http.Request)
	r.Header = http.Header{}
	println("ip:", ClientIP(r))
}
