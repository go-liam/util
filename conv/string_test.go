package conv

import (
	"fmt"
	"log"
	"testing"
)

func TestFunc_StringArrayToString(t *testing.T) {
	v := StringArrayToString("[a,b,c]")
	println(v)
}

func TestFunc_StringToStringArray(t *testing.T) {
	v := StringToStringArray("[a,b,c]")
	fmt.Printf("%+v\n", v)
}

func TestStringToInt(t *testing.T) {
	v := StringToInt("1110", -1)
	println("v=", v)
}

func TestStringToIntArray(t *testing.T) {
	ls := StringToIntArray("[1,2,3,4]")
	log.Printf("ls=%+v\n", ls)
}

func TestStringToInt64(t *testing.T) {
	println(StringToInt64("123", 0))
}

func TestInt64ToString(t *testing.T) {
	log.Printf("ls=%+v\n", Int64ToString(123456789012345678))
}

func TestIntToString(t *testing.T) {
	log.Printf("ls=%+v\n", Int64ToString(1234))
}
