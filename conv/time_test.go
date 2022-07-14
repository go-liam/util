package conv

import (
	"fmt"
	"testing"
)

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

func TestXxx_DateSubByDay(t *testing.T) {
	day := DateSubByDay("2022-06-01", "2022-06-11")
	//assert.Equal(t, day, 10)
	println(day) //10
}

func TestXxx_DateAddDay(t *testing.T) {
	st := DateAddDay("2022-06-24", 1)
	fmt.Println(":st:", st) //
	//assert.Equal(t, st, "2022-06-25")
	st2 := DateAddDay("2022-06-24", -1)
	fmt.Println(":st2:", st2) //
	//assert.Equal(t, st2, "2022-06-23")
}
