package check

import "github.com/go-liam/util/conv"

func PhoneSimple(phone string) bool {
	l := 11
	if len(phone) != l {
		return false
	}
	i := conv.StringToInt64(phone, 0)
	return i > 10000000000
}
