package main

type Animal interface {
	Name() (string, error)
	Age() (int, error)
}

type Cat struct {
	name string
	age  int
}

func (c *Cat) Name() (string, error) {
	return c.name, nil
}
func (c *Cat) Age() (int, error) {
	return c.age, nil
}

// func main() {
// 	var c = Cat{name: "cat"}
// 	var action Animal = &c
// 	fmt.Println(action.Name())

// }
