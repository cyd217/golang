package main

import (
	"bufio"
	"bytes"
	"fmt"
)

func main() {
	b := bytes.NewBufferString("aaaa")
	bu := bufio.NewReader(b)
	b.WriteString("cccc")
	fmt.Println(bu)
}
