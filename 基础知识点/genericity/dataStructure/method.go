package main

func Add[T int | float32 | float64](a T, b T) T {
	return a + b
}

type MyStruct[T int | string] struct {
}
