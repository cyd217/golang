package main

import (
	"fmt"
)

// nil通道  写阻塞
func blocking1() {
	ch := make(chan int)
	fmt.Println("blocking1 ch := make(chan int)")
	ch <- 1
	fmt.Println("finish", <-ch) //fatal error: all goroutines are asleep - deadlock!
}

// nil通道  读阻塞
func blocking2() {
	ch := make(chan int)
	fmt.Println("blocking2 ch := make(chan int)")
	fmt.Println("finish", <-ch) //fatal error: all goroutines are asleep - deadlock!
}

// 无缓冲的通道  写阻塞
func blocking3() {
	ch := make(chan int)
	fmt.Println("blocking3 ch := make(chan int)")
	ch <- 1
	fmt.Println("finish", <-ch) //fatal error: all goroutines are asleep - deadlock!
}

// 缓冲区满的通道  写阻塞
func blocking5() {
	fmt.Println("blocking5 ch := make(chan int,2)")
	ch := make(chan int, 2)
	for i := 0; i < 3; i++ {
		ch <- i
	}

	fmt.Println("finish", <-ch) //fatal error: all goroutines are asleep - deadlock!
}

// 缓冲区空的通道  读阻塞
func blocking6() {
	ch := make(chan int, 3)
	fmt.Println("blocking6 ch := make(chan int,3)")
	fmt.Println("finish", <-ch) //fatal error: all goroutines are asleep - deadlock!
}

func main() {
	//blocking1()

	//blocking2()

	//blocking3()

	//blocking4()

	//blocking5()

	blocking6()
}
