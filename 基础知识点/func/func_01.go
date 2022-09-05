package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
)

func func1(a, b int) (c int) {
	fmt.Printf("func1 a = %p,b = %p,c = %p\n", &a, &b, &c)
	return a + b
}
func func2(a, b *int) (c int) {
	fmt.Printf("func2 a = %p,b = %p,&a = %p,&b = %p,c = %p\n", a, b, &a, &b, &c)
	c = *a + *b
	return
}

func MD5(str string) string {
	s := md5.New()
	s.Write([]byte(str))
	return hex.EncodeToString(s.Sum(nil))
}

func func3(s string, n ...int) string {
	var x int
	for _, i := range n {
		x += i
	}

	return fmt.Sprintf(s, x)
}

func main() {

	a := 1
	b := 2
	//c := func1(a, b)
	//fmt.Printf(" main a = %p,b = %p,c=%d ,c=%p\n", &a, &b, c, &c)

	c2 := func2(&a, &b)
	fmt.Printf(" main a = %p,b = %p,c=%d ,c=%p\n", &a, &b, c2, &c2)
}
