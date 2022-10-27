package main

import "fmt"

type MyChan[T int | string] chan T

func ChanFunc() {
	ch := make(MyChan[int], 3)
	for i := 1; i <= 3; i++ {
		ch <- i
	}
	fmt.Println(<-ch)
	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
