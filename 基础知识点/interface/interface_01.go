package main

import "fmt"

type eater interface {
	eat()
}
type dog struct {
	Name string
}

type cat struct {
	Name string
}

// dog实现了Sayer接口
func (d dog) eat() {
	fmt.Println(d.Name, " 汪汪汪")
}

// cat实现了Sayer接口
func (c cat) eat() {
	fmt.Println(c.Name, " 喵喵喵")
	c.Name = "CAT"
}

func func01() {
	var (
		x eater        // 声明一个Sayer类型的变量x
		a = cat{"cat"} // 实例化一个cat
		b = dog{"dog"} // 实例化一个dog
	)
	fmt.Printf("a =%p\n", &a)
	x = a // 可以把cat实例直接赋值给x  发生了值拷贝
	fmt.Printf("x =%p\n,x=%v\n", &x, x)
	x.eat() // cat  喵喵喵
	x = b   // 可以把dog实例直接赋值给x
	x.eat() // dog  汪汪汪
}
func main() {
	func01()
}
