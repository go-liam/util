package user

import (
	"encoding/json"
	"fmt"

	"github.com/go-liam/util/database/redis"
)

const redisSecondTime = 32

func (e *SrvUser) CacheOne(id int64) (*Model, error) {
	key := fmt.Sprintf("%s_rb_user_%d", redis.GetConfig().RedisPrefix, id)
	v, err := redis.GetServer().GetBytes(key)
	var info *Model
	var err2 error
	if err == nil {
		err2 = json.Unmarshal(v, &info)
		if err2 == nil {
			//println("redis data")
			return info, nil
		}
	}
	info, _ = e.FindOne(id)
	if info != nil && (info.ID) > 0 {
		byteValue, _ := json.Marshal(&info)
		redis.GetServer().Set(key, byteValue, redisSecondTime)
	}
	return info, nil
}
