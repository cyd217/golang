package main

import (
	"fmt"
	"time"
)

// 停止信号：关闭channel或向channel发送一个元素，使接收方通过channel获得信息后做相应的操作
func signClose() {
	c := make(chan int)
	out := make(chan bool)
	go func() {
		go func() {
			time.Sleep(time.Second)
			c <- 2
		}()
		select {
		case <-time.After(time.Second * 1):
			fmt.Println("超时了")
		case v := <-c:
			fmt.Println("c中读取出的数据：", v)
		}
		out <- true
	}()
	<-out
}

func TimeTask() {
	timer := time.NewTicker(time.Second)
	for {
		select {
		case <-timer.C:
			fmt.Println("执行了") // 每隔1秒执行一次
		}
	}
}
func main() {
	//signClose()
	TimeTask()
}
