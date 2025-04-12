package main

import (
	"fmt"
	"time"
)

// cooperative context switch
func main() {

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("goroutine1 is running")
			time.Sleep(300 * time.Millisecond)
		}
	}()

	go func() {
		for i := 0; i < 3; i++ {
			fmt.Println("goroutine2 is running")
			time.Sleep(300 * time.Millisecond)
		}
	}()
	time.Sleep(2 * time.Second)
}
