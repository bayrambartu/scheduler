package main

import (
	"fmt"
	"runtime"
	"time"
)

func main() {
	runtime.GOMAXPROCS(1)

	for i := 0; i < 3; i++ {
		go func(id int) {
			fmt.Printf("Goroutine %d is running\n", id)
			time.Sleep(500 * time.Millisecond)
		}(i)

	}
	time.Sleep(time.Second * 2)
}
