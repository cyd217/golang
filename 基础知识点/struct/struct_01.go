package main

import "fmt"

func main() {
	//NewStruct()
	//SameSruct()
}

func SameSruct() {
	type MyString = string
	str := "hello"
	a := MyString(str)
	b := MyString("A" + str)
	fmt.Printf("str type is %T\n", str)      //str type is string
	fmt.Printf("a type is %T\n", a)          //a type is string
	fmt.Printf("a == str is %t\n", a == str) //a == str is true
	fmt.Printf("b > a is %t\n", b > a)       //b > a is false
}

// 类型定义
func NewStruct() {
	type MyString string
	str := "hello"
	a := MyString(str)
	b := MyString("A" + str)
	c := MyString(str)

	fmt.Printf("str type is %T\n", str) //str type is string
	fmt.Printf("a type is %T\n", a)     //a type is main.MyString
	fmt.Printf("a value is %#v\n", a)   //a value is "hello"
	fmt.Printf("b value is %#v\n", b)   //b value is "Ahello"
	fmt.Printf("c value is %#v\n", c)   //b value is "Ahello"
	// fmt.Printf("a == str is %t\n", a == str) //报错
	//s1 := make([]string, 0)
	//	s3 := []MyString(s1)

}
