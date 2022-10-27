package main

import (
	"fmt"
)

// 一个泛型类型的结构体。可用 int 或 sring 类型实例化
type MyStruct[T int | string] struct {
	Name string
	Data T
}
type MyStruct2[T int | string, A int | bool] struct {
	Name string
	Data T
	Sex  A
}

func StructFunc() {
	var my MyStruct[int] = MyStruct[int]{
		Name: "caicai",
		Data: 23,
	}
	fmt.Println(my)

	var my2 MyStruct2[int, bool] = MyStruct2[int, bool]{
		Name: "caicai",
		Data: 23,
		Sex:  true,
	}
	fmt.Println(my2)
}
