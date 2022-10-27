package pack

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/go-redis/redis/v8"
)

var rdb *redis.Client

func init() {
	addr := "192.168.101.141"
	port := "6379"
	passwd := ""
	rdb = redis.NewClient(&redis.Options{
		Addr:     addr + ":" + port,
		DB:       0,
		Password: passwd,
		PoolSize: 100,
	})

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	_, err := rdb.Ping(ctx).Result()
	if err != nil {
		log.Fatalln("redis ping err:", err)
		panic("redis连接失败：" + err.Error())
	} else {
		fmt.Println("PING, success")
	}

}

func GetRedisDb() *redis.Client {
	return rdb
}
