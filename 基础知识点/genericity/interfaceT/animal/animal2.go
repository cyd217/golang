package main

import "fmt"

type Animal2 interface {
	~int
	Age2() (int, error)
}

type AnimalImpl[T Animal2] struct {
	age T
}

type myInt int

func (c myInt) Age2() (int, error) {
	return 1, nil
}

// type myString string

// func (c myString) Age2() (int, error) {
// 	return 1, nil
// }

func main() {

	var my myInt = 1
	var s1 AnimalImpl[myInt] = AnimalImpl[myInt]{
		age: my,
	}
	fmt.Println(s1)
	// //myString does not implement Animal2
	// var mystring myString = "1"
	// var s2 AnimalImpl[myString] = AnimalImpl[myString]{
	// 	age: mystring,
	// }
	// fmt.Println(s2)

}
