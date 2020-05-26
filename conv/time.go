package conv

import "time"

// TimeNowInt64 :
func TimeNowInt64() int64 {
	return time.Now().Unix()
}

// TimeNowFormat :
func TimeNowFormat(format string) string {
	timeLocal, _ := time.LoadLocation("Asia/Chongqing")
	time.Local = timeLocal
	curNow := time.Now().Local()
	return curNow.Format(format)
}
