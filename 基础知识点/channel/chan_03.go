package main

import (
	"sync"
	"time"
)

// 未初始化时关闭
func panic1() {
	var ch chan int
	close(ch) //panic: close of nil channel
}

// 重复关闭
func panic2() {
	var ch = make(chan int)
	close(ch)
	close(ch) //panic: close of nil channel
}

// 关闭后发送
func panic3() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		close(ch)
		ch <- 1 //panic: send on closed channel
	}()
	wg.Wait()
}

// 4.发送时关闭
func panic4() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(1)

	go func() {
		defer wg.Done()
		defer close(ch)

		go func() {
			ch <- 1 //阻塞
		}()

		time.Sleep(time.Second) // 等待向 errCh 发送数据
	}()

	wg.Wait()

	// Output:
	//
}
func main() {
	//panic1()
	//panic2()

	panic3()
}
