package main

import (
	"fmt"
)

type SliceT2[T int | float32 | int32] []T

type MyStruct3[T int | string] struct {
	Name  string
	Data  T
	habby SliceT[T]
}

func Struct2Func() {
	var my MyStruct3[string] = MyStruct3[string]{
		Name:  "caicai",
		Data:  "23",
		habby: []string{"swim"},
	}

	fmt.Println(my)
}
