package conv

import (
	"fmt"
	"log"
	"testing"
)

func Test_ArrayToString(t *testing.T) {
	list := []int64{1, 2, 3, 4}
	//list := []int64{1}
	rs := ArrayToString(list, ",")
	fmt.Println(rs)
	// 1,2,3,4
}

func Test_ArrayIntToString(t *testing.T) {
	list := []int{1, 2, 3, 4}
	//list := []int64{1}
	rs := ArrayIntToString(list, ",")
	fmt.Println(rs)
	// 1,2,3,4
}


func TestNilArrayStringTChange(t *testing.T) {
	NilArrayStringToChange(nil)
}

func TestNilArrayIntTChange(t *testing.T) {
	NilArrayIntToChange(nil)
}

func TestArrayInt64Contains(t *testing.T) {
	ArrayInt64Contains(nil,1)
}

func TestArrayIntContains(t *testing.T) {
	println( ArrayIntContains(nil,2))
}

func TestArrayInt64Join(t *testing.T) {
	a1 := []int64{1,2,3}
	a2 := []int64{1,2,3}
	join := ArrayInt64Join(a1,a2)
	log.Printf("v=%+v\n ",join)
}


func TestArrayIntJoin(t *testing.T) {
	a1 := []int{1,2,3}
	a2 := []int{1,2,3,4}
	join := ArrayIntJoin(a1,a2)
	log.Printf("v=%+v\n ",join)
}

func TestArrayStringJoin(t *testing.T) {
	a1 := []string{"1","2","3" }
	a2 := []string{"1","2","3"}
	join := ArrayStringJoin(a1,a2)
	log.Printf("v=%+v\n ",join)
}

func TestArrayStringContains(t *testing.T) {
	a1 := []string{"1","2","3" }
	a2 := []string{"1","2","3"}
	join := ArrayStringJoin(a1,a2)
	tag := "300"
	f := ArrayStringContains(join,tag)
	log.Printf("v=%+v\n ",f)
}
