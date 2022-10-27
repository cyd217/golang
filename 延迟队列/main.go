package main

import (
	"context"
	"delay_demo/pack"
	_ "delay_demo/pack"
	"delay_demo/pool"
	"fmt"
	"time"
)

func main() {

	var data pool.RedisJobData
	redisConn := pack.GetRedisDb()
	conn := context.Background()
	data.SetJobPool(conn, 10)
	pool2, err := pool.NewPool(5)
	if err != nil {
		panic(err)
	}

	c := time.Tick(1 * time.Second)

	for next := range c {
		fmt.Println("我在执行了")
		err := pool.TimerDelayBucket(redisConn, conn, pool2)
		if err != nil {
			fmt.Println("定时timer发生错误：", next, err)
		}
	}

	fmt.Println("end")
	pool2.Close()

}
