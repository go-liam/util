package conv

import "testing"

func TestTimeNowInt64(t *testing.T) {
	println(TimeNowInt64())
}

func TestTimeTimeNowFormat(t *testing.T) {
	v1 := TimeNowFormat("2006-01-02 15:04:05")
	println(v1)
	v2 := TimeNowFormat("")
	println(v2)
	v3 := TimeNowFormat("2006年01月02日 15:04")
	println(v3)
}

func TestTimeToday(t *testing.T) {
	println("v=", TimeTodayInt64())
}

// 1592755200
func TestTimesTampToday(t *testing.T) {
	println(TimesTampToday())
}

// 1592841600
func TestTimesTampTomorrow(t *testing.T) {
	println(TimesTampTomorrow())
}
