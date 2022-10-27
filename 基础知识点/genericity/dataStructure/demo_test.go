package main

import (
	"fmt"
	"testing"
)

func TestSlice(t *testing.T) {
	var mySlice = MySlice[int]{}
	mySlice = append(mySlice, []int{1, 2, 3, 4, 5, 6, 7, 8, 9}...)
	fmt.Println(mySlice.Sum())
}
func TestQueue(t *testing.T) {
	var q1 MyQueue[int] // 可存放int类型数据的队列
	q1.Push(1)
	q1.Push(2)
	q1.Push(3)
	fmt.Println(q1.Pop()) // 1
	fmt.Println(q1.Pop()) // 2
	fmt.Println(q1.Pop()) // 3

	var q2 MyQueue[string] // 可存放string类型数据的队列
	q2.Push("A")
	q2.Push("B")
	q2.Push("C")
	fmt.Println(q2.Pop()) // "A"
	fmt.Println(q2.Pop()) // "B"
	fmt.Println(q2.Pop()) // "C"

	// var q3 Queue[struct{ Name string }]
	// var q4 Queue[[]int]     // 可存放[]int切片的队列
	// var q5 Queue[chan int]  // 可存放int通道的队列
	// var q6 Queue[io.Reader] // 可存放接口的队列
}

func TestADD(t *testing.T) {
	fmt.Println(Add[int](100, 200))
	fmt.Println(Add[float64](10.5, 24.6))
}
