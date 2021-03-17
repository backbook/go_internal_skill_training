package main

import (
	"fmt"
	"time"
)

func hello() {
	fmt.Println("hello goroutine")
}

func main() {
	go hello()
	fmt.Println("main goroutine done")
	//挂载到协程。如果挂了就没了
	time.Sleep(time.Second)
}
