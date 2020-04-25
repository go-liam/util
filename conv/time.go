package conv

import "time"

// TimeNowInt64 :
func TimeNowInt64() int64 {
	return time.Now().Unix()
}