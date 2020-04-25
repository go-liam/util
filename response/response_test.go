package response

import (
	"log"
	"testing"
)

func TestAPIResponseData(t *testing.T) {
	resp := APIResponseData(0,"OK",nil)
	log.Printf("v1=%+v\n",resp)
	resp2 := APIResponseData(0,"","hello")
	log.Printf("v2=%+v\n",resp2)
}
