package redis

import (
	"log"

	en "github.com/timest/env"
)

var (
	// MysqlConfig :
	config *redisConfig
)

type redisConfig struct {
	// redis
	RedisHost   string `env:"REDIS_HOST" default:"localhost"`
	RedisPort   string `env:"REDIS_PORT" default:"6379"`
	RedisPwd    string `env:"REDIS_PWD" default:"Makeblock123"`
	RedisPrefix string `env:"REDIS_PREFIX" default:"cms|"`
}

func GetConfig() *redisConfig {
	if !isInit {
		InitConfig()
		InitServer()
	}
	return config
}

func InitConfig() {
	if isInit {
		return
	}
	config = new(redisConfig)
	en.IgnorePrefix()
	err := en.Fill(config)
	log.Printf("[INFO] RedisConfig :%+v\n", config)
	if err != nil {
		log.Printf("[ERROR] RedisConfig :%+v\n", err)
	}
	isInit = true
}
