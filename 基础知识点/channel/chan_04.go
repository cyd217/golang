package main

import (
	"fmt"
	"runtime"
	"time"
)

// 不关闭通道会怎么样
func close1() {
	fmt.Println("NumGoroutine:", runtime.NumGoroutine())
	ich := make(chan int)

	// sender
	go func() {
		for i := 0; i < 3; i++ {
			ich <- i
		}
	}()

	// receiver
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(<-ich)
		}
	}()

	time.Sleep(time.Second)
	fmt.Println("NumGoroutine:", runtime.NumGoroutine())
	/**
	  NumGoroutine: 1
	  0
	  1
	  2
	  NumGoroutine: 1
	*/
}

func close2() {
	fmt.Println("NumGoroutine:", runtime.NumGoroutine())
	ich := make(chan int)

	// sender
	go func() {
		for i := 0; i < 2; i++ {
			ich <- i
		}
	}()

	// receiver
	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println(<-ich)
		}
	}()

	time.Sleep(time.Second)
	fmt.Println("NumGoroutine:", runtime.NumGoroutine())

	// Output:
	// NumGoroutine: 1
	// 0
	// 1
	// NumGoroutine: 2
}

// 使用 channel 的多重返回值
func isClose1() {

	var ch = make(chan int)
	go func() {
		defer close(ch)
		ch <- 1
	}()

	go func() {
		for i := 0; i < 3; i++ {
			if v, ok := <-ch; ok {
				fmt.Println(i, v)
			}
		}
	}()

	time.Sleep(time.Second)
	// Output:
	// 0 1

}

// 使用 for range 简化语法
func isClose2() {
	var errCh = make(chan int)
	go func() {
		defer close(errCh)
		errCh <- 1
	}()

	go func() {
		i := 0
		for err := range errCh {
			fmt.Println(i, err)
			i++
		}
	}()

	time.Sleep(time.Second)

	// Output:
	// 0 1
}
func main() {
	//close1()
	//close2()

	isClose1()
	isClose2()
}
