package main

import "fmt"

func main() {

	//创建切片
	//create()

	//扩容
	//add()

	//copySlice()

	//遍历
	rangeSlice()
}

func rangeSlice() {

	a := [...]int{2, 3, 5, 7, 11}
	fmt.Printf("  a = %v, Pointer = %p, len = %d, cap = %d\n", a, &a, len(a), cap(a))
	for _, prime := range a {
		prime -= 1
	}
	fmt.Printf("  a = %v, Pointer = %p, len = %d, cap = %d\n", a, &a, len(a), cap(a))
	for i, v := range a {
		fmt.Printf(" i = %v, Pointer = %p,  v = %v, Pointer = %p\n", i, &i, v, &v)
		a[i] -= 1
	}
	fmt.Printf("  a = %v, Pointer = %p, len = %d, cap = %d\n", a, &a, len(a), cap(a))

}

func copySlice() {
	dst := []int{1, 2, 3, 77, 7, 4, 2}
	src := []int{4, 5, 6, 7, 8}
	fmt.Printf("Before dst = %v, Pointer = %p, len = %d, cap = %d\n", dst, &dst, len(dst), cap(dst))
	fmt.Printf("Before src = %v, Pointer = %p, len = %d, cap = %d\n", src, &src, len(src), cap(src))
	//copy 函数将数据从源 Slice复制到目标 Slice。它返回复制的元素数。
	n := copy(dst, src)
	fmt.Printf("n = %d After dst = %v, Pointer = %p, len = %d, cap = %d\n", n, dst, &dst, len(dst), cap(dst))
	fmt.Printf("n = %d After src = %v, Pointer = %p, len = %d, cap = %d\n", n, src, &src, len(src), cap(src))
}

func add() {
	slice := []int{10, 20, 30, 40}
	newSlice := append(slice, 50)
	fmt.Printf("Before slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	fmt.Printf("Before newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))
	newSlice[1] += 10
	slice[3] += 10
	fmt.Printf("After slice = %v, Pointer = %p, len = %d, cap = %d\n", slice, &slice, len(slice), cap(slice))
	fmt.Printf("After newSlice = %v, Pointer = %p, len = %d, cap = %d\n", newSlice, &newSlice, len(newSlice), cap(newSlice))

}

func create() {
	//声明切片
	var s1 []int
	if s1 == nil {
		fmt.Println("是空") //是空
	} else {
		fmt.Println("不是空")
	}
	fmt.Println(s1, len(s1), cap(s1)) //[] 0 0

	//make切片
	var s3 []int = make([]int, 10)
	fmt.Println(s3, len(s3), cap(s3)) //[0 0 0 0 0 0 0 0 0 0] 10 10

	var s4 []int = make([]int, 0, 10)
	fmt.Println(s4, len(s4), cap(s4)) //[] 0 10

	//从数组中切片
	arr := [5]int{1, 2, 3, 4, 5}
	var s6 []int
	// 前包后不包
	s6 = arr[1:4]
	fmt.Println(s6, len(s6), cap(s6)) //[2 3 4] 3 4
	s6 = arr[1:4:4]
	fmt.Println(s6, len(s6), cap(s6)) //[2 3 4] 3 3

}
