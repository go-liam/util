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
func TestStringToInt64Array(t *testing.T) {
	ls := StringToInt64Array("[1,2,3,4]")
	log.Printf("ls=%+v\n", ls)
	//ls=[1 2 3 4]
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

func TestStringIsEmptyValue(t *testing.T) {
	println(StringIsEmptyValue("", "xxxxx"))
}

func TestStringIsEqual(t *testing.T) {
	println(StringIsEqual("a123", "a123"))
	println(StringIsEqual("a123", "a12"))
}

func TestString_StringLeft(t *testing.T) {
	println(StringLeft("12345aSdFg", 100))
	println(StringLeft("一二三四五12345a2sDfg", 7))
	println(StringLeft("一1二2三3", 7))
}

func TestInt64StringToArrayString(t *testing.T) {
	v := Int64StringToArrayString("[1,2,3]")
	log.Printf("v=%+v\n", v)
	//v=[1 2 3]
}

func TestArrayStringToInt64String(t *testing.T) {
	v := ArrayStringToInt64String([]string{"1", "2"})
	println("v=%+v\n", v)
	// v=[1,2]

}
