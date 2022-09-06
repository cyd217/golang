package main

import "fmt"

func method() {
	//发送（send）
	inch := make(chan int, 10)
	inch <- 10 // 把10发送到ch中
	fmt.Println(len(inch))

	//接受
	x, ok := <-inch
	fmt.Println(x, ok)

	//关闭
	close(inch)
}

// 创建channel
func create() {
	var ch chan int
	fmt.Println(ch) // <nil>

	inch := make(chan int, 10)
	fmt.Println(len(inch))
}
func main() {
	create()
}
