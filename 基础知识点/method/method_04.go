package main

import "fmt"

type teacher struct {
	id   int
	name string
}

func (t *teacher) TestP() {
	fmt.Printf("TestPointer: %p, %v\n", t, t)
}

func (t teacher) TestV() {
	fmt.Printf("TestValue: %p, %v\n", &t, t)
}

func main() {
	t := teacher{1, "Tom"}
	fmt.Printf("User: %p, %v\n", &t, t)

	mv := t.TestV
	mv()

	mp := (*teacher).TestP
	mp(&t)

	mp2 := (*teacher).TestV
	mp2(&t)
}
