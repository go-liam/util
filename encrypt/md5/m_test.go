package md5

import "testing"

func TestGetMd5(t *testing.T) {
	key := "abc123$$##@@!!**xziKwm095"
	v := GetMD5Hash(key + "123456")
	println("v=", v)
	v = GetMD5Hash(key + "123456789")
	println("v=", v)
}
