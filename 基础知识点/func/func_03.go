package main

import "fmt"

func func01() {
	var whatever [5]struct{}

	for i := range whatever {
		defer fmt.Println(i)
	}
}

// 闭包
// 也就是说函数正常执行,由于闭包用到的变量 i 在执行的时候已经变成4,所以输出全都是4.
func func02() {
	var whatever [5]struct{}
	for i := range whatever {
		defer func() { fmt.Println(i) }()
	}
}

type Test struct {
	name string
}

func (t *Test) Close() {
	fmt.Println(t.name, " closed")
}
func func03() {
	ts := []Test{{"a"}, {"b"}, {"c"}}
	for _, t := range ts {
		defer t.Close()
	}
}

func Close(t Test) {
	t.Close()
}

func func03More() {
	ts := []Test{{"a"}, {"b"}, {"c"}}
	// 01
	for _, t := range ts {
		defer Close(t)
	}
	//02
	for _, t := range ts {
		t2 := t
		defer t2.Close()
	}
}

// FILO
func func04(x int) {
	defer println("a")
	defer println("b")

	defer func() {
		println(100 / x) // div0 异常未被捕获，逐步往外传递，最终终止进程。
	}()

	defer println("c")
}

// 延迟调用参数在注册时求值或复制，可用指针或闭包 "延迟" 读取。
func func05() {
	x, y := 10, 20

	defer func(i int) {
		println("defer:", i, y) // y 闭包引用
	}(x) // x 被复制

	x += 10
	y += 100
	println("x =", x, "y =", y)
}
func main() {
	// func01()

	//func02()

	//func03()
	//func03More()

	//func04(0)
	func05()
}
