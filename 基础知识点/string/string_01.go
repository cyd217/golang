package main

import (
	"bytes"
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"

// 随机生成长度是n的字符串
func randomString(n int) string {
	b := make([]byte, n)
	for i := range b {
		b[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(b)
}

func concatPlus(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s += str
	}
	return s
}

func concatSprintf(n int, str string) string {
	s := ""
	for i := 0; i < n; i++ {
		s = fmt.Sprintf("%s%s", s, str)
	}
	return s
}

func concatBuilder(n int, str string) string {
	var builder strings.Builder
	for i := 0; i < n; i++ {
		builder.WriteString(str)
	}
	return builder.String()
}

func concatBuffer(n int, s string) string {
	buf := new(bytes.Buffer)
	for i := 0; i < n; i++ {
		buf.WriteString(s)
	}
	return buf.String()
}

func concatByte(n int, str string) string {
	buf := make([]byte, 0)
	for i := 0; i < n; i++ {
		buf = append(buf, str...)
	}
	return string(buf)
}

func concatPreByte(n int, str string) string {
	buf := make([]byte, 0, n*len(str))
	for i := 0; i < n; i++ {
		buf = append(buf, str...)
	}
	return string(buf)
}
func main() {
	str := randomString(10)
	start := time.Now().UnixMicro()
	concatPlus(100000, str)
	end := time.Now().UnixMicro()
	fmt.Printf("concatPlus() time:%d\n", end-start)

	start = time.Now().UnixMicro()
	concatSprintf(100000, str)
	end = time.Now().UnixMicro()
	fmt.Printf("concatSprintf() time:%d\n", end-start)

	start = time.Now().UnixMicro()
	concatBuilder(100000, str)
	end = time.Now().UnixMicro()
	fmt.Printf("concatBuilder() time:%d\n", end-start)

	start = time.Now().UnixMicro()
	concatBuffer(100000, str)
	end = time.Now().UnixMicro()
	fmt.Printf("concatBuffer() time:%d\n", end-start)

	start = time.Now().UnixMicro()
	concatByte(100000, str)
	end = time.Now().UnixMicro()
	fmt.Printf("concatByte() time:%d\n", end-start)

	start = time.Now().UnixMicro()
	concatPreByte(100000, str)
	end = time.Now().UnixMicro()
	fmt.Printf("concatPreByte() time:%d\n", end-start)

}
