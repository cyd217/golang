package main

import (
	"fmt"
	"math/rand"
	"os"
	"time"
)

func main() {
	for i := 10; i < 15; i++ {
		name := fmt.Sprintf(`文件%d.txt`, i)
		back := make([]byte, 1024*1024*200)
		rand.Seed(time.Now().UnixNano())
		rand.Read(back)
		f, _ := os.Create(name)
		defer f.Close()
		f.Write(back)
		time.Sleep(time.Millisecond * 5)
	}

}
