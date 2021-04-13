package conv

import "testing"

func TestName_GenerateRangeNum(t *testing.T) {
	v1 := RandomNumber(0, 1)
	println(v1) // [0,1)不包含1
}
