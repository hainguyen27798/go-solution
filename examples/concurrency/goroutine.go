package main

import (
	"fmt"
	"time"
)

func someFunc(ch <-chan bool) {
	for {
		select {
		case msg := <-ch:
			fmt.Println(msg)
			return
		default:
			fmt.Println("Working")
		}
	}
}

func main() {
	ch := make(chan bool)

	go someFunc(ch)

	time.Sleep(3 * time.Second)
	close(ch)
}
