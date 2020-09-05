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

// RemoveDuplicateArray ([]int{1, 2, 3, 4, 1, 2, 3})
func RemoveDuplicateArray(slc []int) []int {
	result := []int{}
	tempMap := map[int]byte{} // 存放不重复主键
	for _, e := range slc {
		l := len(tempMap)
		tempMap[e] = 0
		if len(tempMap) != l { // 加入map后，map长度变化，则元素不重复
			result = append(result, e)
		}
	}
	return result
}

// ArrayStringToInt64 :string Array 转为 int64 array: ["1","2"]->[1,2]
func ArrayStringToInt64(ls []string) []int64 {
	o := make([]int64, 0)
	for _, v := range ls {
		o = append(o, StringToInt64(v, 0))
	}
	return o
}

// ArrayInt64ToString :int64 array 转为 string Array [1,2] -> ["1","2"]
func ArrayInt64ToString(ls []int64) []string {
	o := make([]string, 0)
	for _, v := range ls {
		o = append(o, fmt.Sprintf("%d", v))
	}
	return o
}
