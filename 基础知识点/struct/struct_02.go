package main

import "fmt"

type Person struct {
	Name string
	Age  int
	sex  bool
}

func main() {
	//func1()
	//func2()
	//func3()
	//func4()
	func5()
}

func func5() {
	type student struct {
		name string
		age  int
	}
	m := make(map[string]*student)
	stus := []student{
		{name: "pprof.cn", age: 18},
		{name: "测试", age: 23},
		{name: "博客", age: 28},
	}

	for _, stu := range stus {
		m[stu.name] = &stu
	}
	for k, v := range m {
		fmt.Println(k, "=>", v.name)
	}
}

func func4() {
	n := Person{
		"caicai",
		23,
		false,
	}
	fmt.Printf("n %p\n", &n)
	fmt.Printf("n.name %p\n", &n.Name)
	fmt.Printf("n.age %p\n", &n.Age)
	fmt.Printf("n.sex %p\n", &n.sex)
}
func func3() {
	var user struct {
		Name string
		Age  int
	}
	user.Name = "pprof.cn"
	user.Age = 18
	fmt.Printf("%#v\n", user)
}

func func1() {
	var p1 Person
	p1.Name = "Tom"
	p1.Age = 30
	fmt.Println("p1 =", p1)

	var p2 = Person{Name: "Burke", Age: 31}
	fmt.Println("p2 =", p2)

	p3 := Person{Name: "Aaron", Age: 32}
	fmt.Println("p2 =", p3)
}

func func2() {
	//var p1 *Person
	//p1.Name = "Tom"
	//p1.Age = 30
	//fmt.Println("p1 =", p1) //panic: runtime error: invalid memory address or nil pointer dereference

	var p2 = &Person{Name: "Burke", Age: 31}
	fmt.Printf("类型: %T,p2 =%v,p2地址 = %p", p2, p2, p2)
	var p3 = new(Person)
	p3.Name = "测试"
	p3.Age = 18
	fmt.Printf("p3=%#v\n", p3) //p3=&main.person{name:"测试", city:"北京", age:18}
}
