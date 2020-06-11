package conv

import (
	"log"
	"testing"
)

type ModelUser struct {
	UserID   int64
	UserName string
}

/*
=== RUN   TestName_StructToJsonString
{"i":110,"p":"","n":"name","e":"","t":false,"c":0}
*/
func TestName_StructToJsonString(t *testing.T) {
	us := new(ModelUser)
	us.UserID = 110
	us.UserName = "name"
	v1 := StructToJsonString(us)
	println(v1)
}

func TestName_JsonStringToStruct(t *testing.T) {
	st := `{"i":110,"p":"pp","n":"name","e":"eee","t":false,"c":0}`
	us := new(ModelUser)
	JsonStringToStruct(st, us)
	log.Printf("v= %+v \n", us)
}

func TestJson_ChangeStringToInterface(t *testing.T) {
	st := `{"i":110,"p":"pp","n":"name","e":"eee","t":false,"c":0}`
	v := StringToInterface(st)
	log.Printf("v= %+v \n", v)
}
