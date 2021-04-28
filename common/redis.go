package common

import (
	"github.com/gomodule/redigo/redis"
	"go-playground/config"
)

var pool *redis.Pool

func init() {
	pool = &redis.Pool{
		MaxIdle:     16,  // 最大连接数
		MaxActive:   0,   // 连接池最大连接数量,不确定可以用0（0表示自动定义），按需分配
		IdleTimeout: 300, // 连接关闭时间 300秒 （300秒不使用自动关闭）
		Dial: func() (redis.Conn, error) {
			return redis.Dial("tcp", config.REDIS)
		},
	}
}

func GetRedisPool() *redis.Pool {
	return pool
}

func GetRedisConnect() redis.Conn {
	return GetRedisPool().Get()
}
