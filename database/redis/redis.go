package redis

import (
	"fmt"
	//redis2 "github.com/go-liam/cache/redis"
)

var server = new(SvRedis) //cache.InCache
var isInit = false

func GetServer() *SvRedis {
	if !isInit {
		InitConfig()
		InitServer()
	}
	return server
}

func InitServer() {

	server.URL = fmt.Sprintf("%s@%s:%s", config.RedisPwd, config.RedisHost, config.RedisPort) // "Liam123@localhost:6379"
	server.RedisPrefix = config.RedisPrefix
	server.NewCache(server.URL, server.RedisPrefix)
}
