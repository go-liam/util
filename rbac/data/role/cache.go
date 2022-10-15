package role

import (
	"encoding/json"
	"fmt"

	"github.com/go-liam/util/database/redis"
)

const redisSecondTime = 32

func (e *SrvRole) CacheMulti() ([]*Model, error) {
	key := fmt.Sprintf("%s_rb_role_ls", redis.GetConfig().RedisPrefix)
	v, err := redis.GetServer().GetBytes(key)
	var got []*Model
	var err2 error
	if err == nil {
		err2 = json.Unmarshal(v, &got)
		if err2 == nil {
			//println("redis data")
			return got, nil
		}
	}
	got, _ = e.FindMultiByNil()
	if len(got) > 0 {
		byteValue, _ := json.Marshal(&got)
		redis.GetServer().Set(key, byteValue, redisSecondTime)
	}
	return got, nil
}
