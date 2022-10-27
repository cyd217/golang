package main

type Int interface {
	int | int8 | int16 | int32 | int64
}

type Uint interface {
	uint | uint8 | uint16 | uint32
}

type Float interface {
	float32 | float64
}

type Slice[T Int | Uint | Float] []T // 使用 '|' 将多个接口类型组合

type SliceElement interface {
	Int | Uint | Float
}

type Slice33[T SliceElement] []T // 使用 '|' 将多个接口类型组合

type Int2 interface {
	~int | ~int8 | ~int16 | ~int32 | ~int64
}

type Uint2 interface {
	~uint | ~uint8 | ~uint16 | ~uint32
}
type Float2 interface {
	~float32 | ~float64
}

type Slice2[T Int2 | Uint2 | Float2] []T

var s Slice[int] // 正确

type MyInt2 int

var s2 Slice2[MyInt2] // MyInt底层类型是int，所以可以用于实例化

type MyMyInt MyInt2

var s3 Slice2[MyMyInt] // 正确。MyMyInt 虽然基于 MyInt ，但底层类型也是int，所以也能用于实例化
