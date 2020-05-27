package conv

import "time"

// TimeNowInt64 :
func TimeNowInt64() int64 {
	return time.Now().Unix()
}

// TimeNowFormat :
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
	t := time.Now()
	zeroTm := time.Date(t.Year(), t.Month(), t.Day(), 0, 0, 0, 0, t.Location()).Unix()
	return zeroTm
}
