package main

// 一个泛型接口
type IPrintData[T int | float32 | string] interface {
	Print(data T)
}
