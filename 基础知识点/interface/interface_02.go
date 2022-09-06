package main

import "fmt"

type Inter interface {
	foo()
}

type S struct{}

// 值类型
//func (s S) foo() {
//	fmt.Println("s.foo")
//}

// 或者指针类型
func (s *S) foo() {
	fmt.Println("s.foo")
}

// 接收者是值类型时，调用方可以时值类型也可以指针类型
func fun01() {
	//赋值
	//var s1 Inter = S{} // 值类型
	//s1.foo()
	var s2 Inter = &S{} // 指针类型
	s2.foo()
}

// 接收者是指针类型时，调用方只能是指针类型
func fun02() {
	//赋值
	//var s1 Inter = S{} // 值类型
	//s1.foo()
	var s2 Inter = &S{} // 指针类型
	s2.foo()
}
func main() {
	fun01()
}
