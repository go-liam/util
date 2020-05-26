package conv

import "testing"

func TestTimeNowInt64(t *testing.T) {
	println(TimeNowInt64())
}

func TestTimeTimeNowFormat(t *testing.T) {
	v1 := TimeNowFormat("2006-01-02 15:04:05")
	println(v1)
}
