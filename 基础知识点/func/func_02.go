package main

import "fmt"

// 命名返回参数允许 defer 延迟调用通过闭包读取和修改。
func add(x, y int) (z int) {
	defer func() {
		z += 100
	}()

	z = x + y
	return z // 执行顺序: ( z ) -> (call defer) -> (return)
}

// 闭包
// 函数a()外的变量c引用了函数a()内的函数b() 形成闭包
// 由于闭包的存在使得函数a()返回后，a中的i始终存在，这样每次执行c()，i都是自加1后的值。 从上面可以看出闭包的作用就是在a()执行完并返回后，闭包使得垃圾回收机制GC不会收回a()所占用的资源，因为a()的内部函数b()的执行需要依赖a()中的变量i。
func a() func() int {
	i := 0
	b := func() int {
		i++
		fmt.Println(i)
		return i
	}
	return b
}

func main() {
	c := a()
	c()
	c()
	c()
	a()
	//println(add(1, 2))
}
