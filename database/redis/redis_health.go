package redis

import (
	"log"
	"time"
)

func Health() string {
	log.Println("[INFO] redis url=", server.URL)
	log.Println("[INFO] redis RedisPrefix=", server.RedisPrefix)
	n := server.RedisPrefix + "|health|0"
	server.Set(n, time.Now().Unix(), 2)
	f, err := server.Get(n)
	if err != nil {
		log.Printf("[ERROR] redis health:%+v\n", err)
		server.NewCache(server.URL, server.RedisPrefix) //重新连接
	}
	return f
}
