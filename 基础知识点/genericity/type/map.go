package main

import "fmt"

type MyMap[KEY int | string, VALUE float32 | float64] map[KEY]VALUE

func MapFunc() {
	// 用类型实参 string 和 flaot64 替换了类型形参 KEY 、 VALUE，泛型类型被实例化为具体的类型：MyMap[string, float64]
	var mp MyMap[string, float64] = map[string]float64{
		"jack_score": 9.6,
		"bob_score":  8.4,
	}
	fmt.Println(mp)
}
