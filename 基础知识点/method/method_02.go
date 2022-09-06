package main

import (
	"fmt"
)

// User 结构体
type user struct {
	Name  string
	Email string
}

// Notify 方法
func (u user) Notify() {
	fmt.Printf("%v : %v \n", u.Name, u.Email)
}

// Notify 方法
func (u user) SetName(name string) {
	fmt.Printf("u: %p\n", &u)
	u.Name = name
}

// Notify 方法
func (u *user) SetName2(name string) {
	fmt.Printf("*u: %p\n", &u)
	u.Name = name
}

// Go 调整和解引用指针使得调用可以被执行。
// 注意，当接受者不是一个指针时，该方法操作对应接受者的值的副本(意思就是即使你使用了指针调用函数，但是函数的接受者是值类型，所以函数内部操作还是对副本的操作，而不是指针操作。
func main() {
	// 值类型调用方法
	u1 := user{"golang", "golang@golang.com"}
	fmt.Printf("u1: %p\n", &u1)
	u1.SetName("java")
	u1.Notify() //golang : golang@golang.com

	u2 := user{"golang", "golang@golang.com"}
	fmt.Printf("u2: %p\n", &u2)
	u2.SetName2("java")
	u2.Notify() //java  : golang@golang.com

	// 指针类型调用方法
	u3 := &user{"go", "go@go.com"}
	fmt.Printf("u3: %p\n", &u3)
	u3.SetName("java")
	u3.Notify() //go : go@go.com

	u4 := &user{"go", "go@go.com"}
	fmt.Printf("u4: %p\n", &u4)
	u4.SetName2("java")
	u4.Notify() //java : go@go.com
}
