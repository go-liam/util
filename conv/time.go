package conv

import "time"

// TimeNowInt64 :1592793395
func TimeNowInt64() int64 {
	timeLocal, _ := time.LoadLocation("Asia/Chongqing")
	time.Local = timeLocal
	t := time.Now().Local()
	return t.Unix()
}

// TimeNowFormat :2020-06-22 10:36:09
func TimeNowFormat(format string) string {
	if format == "" {
		format = "2006-01-02 15:04:05"
	}
	timeLocal, _ := time.LoadLocation("Asia/Chongqing")
	time.Local = timeLocal
	curNow := time.Now().Local()
	return curNow.Format(format)
}

// TimeTodayInt64 :
func TimeTodayInt64() int64 {
	return TimesTampToday()
}

// TimesTampToday :1592755200
func TimesTampToday() int64 {
	timeLocal, _ := time.LoadLocation("Asia/Chongqing")
	time.Local = timeLocal
	t := time.Now().Local()
	//t := time.Now()
	zeroTm := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix()
	return zeroTm
}

// TimesTampTomorrow :1592841600
func TimesTampTomorrow() int64 {
	timeLocal, _ := time.LoadLocation("Asia/Chongqing")
	time.Local = timeLocal
	t := time.Now().Local()
	//t := time.Now()
	t2 := t.AddDate(0, 0, 1)
	zeroTm := time.Date(t2.Year(), t2.Month(), t2.Day(), 0, 0, 0, 0, t.Location()).Unix()
	return zeroTm
}
