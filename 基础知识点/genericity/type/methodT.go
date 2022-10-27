package main

// type add[T int | string | float32] T
func Add[T int | float64](a T, b T) T {
	return a + b
}
