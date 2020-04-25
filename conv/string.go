package conv

import (
	"fmt"
	"strconv"
	"strings"
)

// StringArrayToString :
func StringArrayToString(rs interface{}) string {
	v := fmt.Sprintf("%v", rs)
	return strings.Trim(strings.Replace(v, " ", ",", -1), "")
}

// StringToStringArray : [a,b,c]
func StringToStringArray(rs string) []string {
	if rs == "" {
		return []string{}
	}
	s0 := strings.Trim(rs, "[]")
	s := strings.Split(s0, ",")
	return s
}
// StringToIntArray : "[1,2,3,4]"
func StringToIntArray(rs string) []int {
	if rs == "" {
		return []int{}
	}
	s0 := strings.Trim(rs, "[]")
	s := strings.Split(s0, ",")
	var ls []int
	for _,v:= range s{
		ls = append(ls,StringToInt(v,0))
	}
	return ls
}


// StringToInt :
func StringToInt(st string, defaultInt int) int {
	if st == "" {
		return defaultInt
	}
	vInt, err := strconv.Atoi(st)
	if err != nil {
		println("[ERROR]string到int:", err)
		return defaultInt
	}
	return vInt
}

// StringToInt64 :
func StringToInt64(st string, defaultInt int64) int64 {
	if st == "" {
		return defaultInt
	}
	vInt, err := strconv.ParseInt(st, 10, 64)
	if err != nil {
		println("[ERROR]string到int64:", err)
		return defaultInt
	}
	return vInt
}