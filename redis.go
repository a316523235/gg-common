package gg_common

import (
	"fmt"
	"github.com/a316523235/gg-common/conf"
	"github.com/gomodule/redigo/redis"
)

var redisPool = map[string]*redis.Pool{}

func GetRedis() redis.Conn {
	connKey := "default"
	return getRedisPool(connKey).Get()
}

func getRedisPool(connKey string) *redis.Pool {
	if redisPool[connKey] == nil {
		redisPool[connKey] = &redis.Pool{
			Dial: func() (conn redis.Conn, e error) {
				db := conf.Conn["mysql"][connKey];
				host, port, _ := db["host"], db["port"], db["pwd"]
				dataSource := fmt.Sprintf("%s:%s", host, port)
				conn, e = redis.Dial("tcp", dataSource)
				if e != nil {
					fmt.Println("redis error : "+e.Error())
					panic(e.Error())
				}
				return
			},
		}
	}
	return redisPool[connKey]
}