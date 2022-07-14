package conv

import (
	"fmt"
	"strconv"
)

func InterfaceToBoolToInt(val interface{}, defaultInt int) int {
	if val == nil {
		return defaultInt
	}
	s := fmt.Sprintf("%v", val)
	if s == "1" || s == "true" || s == "True" {
		return 1
	}
	return 0
}

func InterfaceToBool(val interface{}, defaultInt bool) bool {
	if val == nil {
		return defaultInt
	}
	s := fmt.Sprintf("%v", val)
	if s == "1" || s == "true" || s == "True" {
		return true
	}
	return false
}

func InterfaceToFloat32(val interface{}, defaultInt float32) float32 {
	if val == nil {
		return defaultInt
	}
	s := fmt.Sprintf("%v", val)
	if s1, err := strconv.ParseFloat(s, 32); err == nil {
		return float32(s1)
	}
	return 0
}

func InterfaceToFloat64(val interface{}, defaultInt float64) float64 {
	if val == nil {
		return defaultInt
	}
	s := fmt.Sprintf("%v", val)
	if s1, err := strconv.ParseFloat(s, 32); err == nil {
		return s1
	}
	return 0
}

func InterfaceToInt(val interface{}, defaultInt int) int {
	if val == nil {
		return defaultInt
	}
	s := fmt.Sprintf("%v", val)
	ii := StringToInt(s, defaultInt)
	return ii
}

func InterfaceToInt64(val interface{}, defaultInt int64) int64 {
	if val == nil {
		return defaultInt
	}
	s := fmt.Sprintf("%v", val)
	ii := StringToInt64(s, defaultInt)
	return ii
}

func InterfaceToString(val interface{}) string {
	if val == nil {
		return ""
	}
	s := fmt.Sprintf("%v", val)
	return s
}

func InterfaceToStringLeft(val interface{}, num int) string {
	if val == nil {
		return ""
	}
	s := fmt.Sprintf("%v", val)
	return StringLeft(s, num)
}
