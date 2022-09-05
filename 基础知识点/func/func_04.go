package main

import (
	"fmt"
	"time"
)

func func401() {
	defer func() {
		if err := recover(); err != nil {
			println(err.(string)) // 将 interface{} 转型为具体类型。
		}
		fmt.Println("22222")
	}()

	panic("panic error!")
	fmt.Println("111111")
}

// 如果需要保护代码 段，可将代码块重构成匿名函数，如此可确保后续代码被执 。
func func402(x, y int) {
	var z int

	func() {
		defer func() {
			if recover() != nil {
				z = 0
			}
		}()
		panic("test panic")
		z = x / y
		return
	}()

	fmt.Printf("x / y = %d\n", z)
}

// 携程内panic，主线程捕获不到
func func403() {

	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
	}()
	go func() {
		fmt.Println("11111111")
		panic("test panic")
	}()
	time.Sleep(time.Second)
}

func main() {
	//func401()
	//	func402(1, 2)
	func403()
}
