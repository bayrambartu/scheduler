package main

import (
	"fmt"
	"time"
)

type Task struct {
	ID int
}

func (t *Task) Run() {
	fmt.Printf("Mission %d started \n", t.ID)
	time.Sleep(2 * time.Second)
	fmt.Printf("Mission %d finished \n", t.ID)
}

func main() {

	//array := []int
	tasks := []Task{
		{1}, {2}, {3}, {4}, {5},
		{6}, {7}, {8}, {9}, {10},
	}
	limiter := time.Tick(500 * time.Millisecond)

	for _, task := range tasks {
		<-limiter // rate limiting
		go task.Run()

	}

	time.Sleep(6 * time.Second)
	fmt.Println("All tasks completed.")
}

/*

package main

import (
	"fmt"
	"time"
)

func main() {
	requests := []int{1, 2, 3, 4, 5}

	limiter := time.Tick(1 * time.Second)

	for _, req := range requests {
		<-limiter
		fmt.Println("request:", req)
	}
}
*/

/*
package main

import (
	"fmt"
	"time"
)

func doWork(id int) {
	fmt.Printf("mission %d started \n", id)
	time.Sleep(time.Second * 2)
	fmt.Printf("mission %d finished \n", id)
}
func main() {
	limiter := time.Tick(time.Millisecond * 500)

	for i := 1; i <= 10; i++ {
		<-limiter
		go doWork(i)
	}

	time.Sleep(time.Second * 3)

}

*/
