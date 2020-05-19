package conv

import (
	"fmt"
	"strings"
)

// ArrayToString :
func ArrayToString(list []int64, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(list), " ", delim, -1), "[]")
}

// ArrayIntToString :
func ArrayIntToString(list []int, delim string) string {
	return strings.Trim(strings.Replace(fmt.Sprint(list), " ", delim, -1), "[]")
}

func ArrayStringContains(list []string, item string) bool {
	if len(list) <= 0 {
		return false
	}
	for _, v := range list {
		if strings.Compare(v, item) == 0 {
			return true
		}
	}
	return false
}

// ArrayInt64Join :合并多个 arry
func ArrayInt64Join(arr1 []int64, arr2 []int64) []int64 {
	return append(arr1, arr2...)
}

// ArrayIntJoin :
func ArrayIntJoin(arr1 []int, arr2 []int) []int {
	return append(arr1, arr2...)
}

// ArrayStringJoin :
func ArrayStringJoin(arr1 []string, arr2 []string) []string {
	return append(arr1, arr2...)
}

// NilArrayStringToChange :
func NilArrayStringToChange(info []string) []string {
	if info == nil {
		return make([]string, 0)
	}
	return info
}

// NilArrayIntToChange :
func NilArrayIntToChange(info []int) []int {
	if info == nil {
		return make([]int, 0)
	}
	return info
}

// ArrayInt64Contains :array 是否存在某值
func ArrayInt64Contains(list []int64, item int64) bool {
	if len(list) <= 0 {
		return false
	}
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}

// ArrayIntContains :
func ArrayIntContains(list []int, item int) bool {
	if len(list) <= 0 {
		return false
	}
	for _, v := range list {
		if v == item {
			return true
		}
	}
	return false
}
