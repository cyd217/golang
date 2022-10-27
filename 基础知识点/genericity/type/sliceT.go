package main

import "fmt"

type SliceT[T int | float32 | int32 | string] []T

// 类型形参: T, KEY ,VALUE
// 类型约束: int | float32 | int32, int | string  float32 | float64
// 泛型类型:type SliceT[T int | float32 | int32] []T  type MyMap[KEY int | string, VALUE float32 | float64] map[KEY]VALUE
func SliceFunc() {
	// 这里传入了类型实参int，泛型类型Slice[T]被实例化为具体的类型 Slice[int]
	var a SliceT[int] = []int{1, 2, 3}
	fmt.Printf("Type Name: %T\n", a) //输出：Type Name: Slice[int]

	// 传入类型实参float32, 将泛型类型Slice[T]实例化为具体的类型 Slice[string]
	var b SliceT[float32] = []float32{1.0, 2.0, 3.0}
	fmt.Printf("Type Name: %T\n", b) //输出：Type Name: Slice[float32]

	// ✗ 错误。string不在类型约束 int|float32|float64 中，不能用来实例化泛型类型
	//var c SliceT[string] = []string{"Hello", "World"}
}
