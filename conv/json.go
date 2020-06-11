package conv

import (
	"encoding/json"
	"log"
)

// StructToJsonString :
func StructToJsonString(info interface{}) string {
	if info == nil {
		return ""
	}
	jsons, errs := json.Marshal(info) //转换成JSON返回的是byte[]
	if errs != nil {
		log.Println("[ERROR]", errs.Error())
		//return ""
	}
	return string(jsons) //byte[]转换成string
}

// JsonStringToStruct :
func JsonStringToStruct(jsonSt string, obj interface{}) {
	if jsonSt == "" || obj == nil {
		return
	}
	errs := json.Unmarshal([]byte(jsonSt), obj)
	if errs != nil {
		log.Println("[ERROR]", errs.Error())
	}
}

// StringToInterface :
func StringToInterface(st string) interface{} {
	var obj interface{}
	JsonStringToStruct(st, &obj)
	return obj
}
