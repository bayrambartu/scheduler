package main

import (
	"fmt"
	"net/http"
	"time"
)

func fetch(id int, url string) {
	fmt.Printf("Goroutine is  %d starts.. \n", id)
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("Goroutine %d error: %v\n", id, err)
		return
	}
	defer resp.Body.Close()
	fmt.Printf("Goroutine is  %d ends. \n", id)

}
func main() {
	for i := 1; i <= 3; i++ {
		go fetch(i, "http://httpbin.org/delay/1")

	}
	time.Sleep(3 * time.Second)

}
