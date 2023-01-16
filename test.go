package main

import "fmt"

func CalSquare() {
	src := make(chan int)
	dest := make(chan int, 3)
	go func() {
		defer close(src)
		for i := 0; i < 10; i++ {
			src <- i
		}
	}()
	go func() {
		defer close(dest)
		for i := range src { //src 实现通信
			dest <- i * i
		}
	}()
	for i := range dest {
		//复杂操作
		fmt.Println(i)
	}
}

func main() {
	CalSquare()
}
