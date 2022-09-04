package main

import (
	"encoding/json"
	"fmt"
	"time"
)

func func1() {
	n := 1000000
	sl1 := make([]int, 0)
	start1 := time.Now().UnixMicro()
	for i := 0; i < n; i++ {
		sl1 = append(sl1, i)
	}
	end1 := time.Now().UnixMicro()
	fmt.Println(end1 - start1) //6742

	sl2 := make([]int, 0, n)
	start2 := time.Now().UnixMicro()
	for i := 0; i < n; i++ {
		sl2 = append(sl2, i)
	}
	end2 := time.Now().UnixMicro()
	fmt.Println(end2 - start2) //1544
}

func func2(a []int) {
	a[1] = 10
	fmt.Printf("func2  a =%v,pointer =%p,len = %d, cap = %d\n", a, &a, len(a), cap(a))
}

func func3(a []int) {
	fmt.Printf(" f2 append before  a =%v,pointer =%p,len = %d, cap = %d\n", a, &a, len(a), cap(a))
	a = append(a, 5)
	fmt.Printf(" f2 append after  a =%v,pointer =%p,len = %d, cap = %d\n", a, &a, len(a), cap(a))
}

func func4(a []int) {
	fmt.Printf(" f2 append before  a =%v,pointer =%p,len = %d, cap = %d\n", a, &a, len(a), cap(a))
	a = append(a, 5)
	fmt.Printf(" f2 append after  a =%v,pointer =%p,len = %d, cap = %d\n", a, &a, len(a), cap(a))
}
func main() {

	//func1()
	//-------------------------------------
	//a := []int{1, 2, 3, 4}
	//fmt.Printf(" main  a =%v,pointer =%p,len = %d, cap = %d\n", a, &a, len(a), cap(a))
	//func2(a)
	//fmt.Printf(" main  a =%v,pointer =%p,len = %d, cap = %d\n", a, &a, len(a), cap(a))
	//-------------------------------------
	//b := make([]int, 4, 4)
	//fmt.Printf(" main  b =%v,pointer =%p,len = %d, cap = %d\n", b, &b, len(b), cap(b))
	//func3(b)
	//fmt.Printf(" main  b =%v,pointer =%p,len = %d, cap = %d\n", b[0:4], &b, len(b), cap(b))
	//-------------------------------------
	//c := make([]int, 0, 4)
	//fmt.Printf(" main  c =%v,pointer =%p,len = %d, cap = %d\n", c, &c, len(c), cap(c))
	//func4(c)
	//fmt.Printf(" main  b =%v,pointer =%p,len = %d, cap = %d\n", c, &c, len(c), cap(c))
	//fmt.Printf(" main  b =%v,pointer =%p,len = %d, cap = %d\n", c[0:4], &c, len(c), cap(c))
	//--------------------------

	var d1 []int
	d2 := make([]int, 0)
	json1, err := json.Marshal(d1)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(json1))
	}
	json2, err := json.Marshal(d2)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(string(json2))
	}
}
