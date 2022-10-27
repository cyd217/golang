package main

import "fmt"

type IntUintFloat interface {
	int | int8 | int16 | int32 | int64 | uint | uint8 | uint16 | uint32 | uint64 | float32 | float64
}

type Slicee[T IntUintFloat] []T

func main() {

	s := make(Slicee[int], 0)
	s = append(s, 1)
	fmt.Printf("%v\n", s)
}
