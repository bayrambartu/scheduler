package main

import (
	"fmt"
	"sync-blocking/scheduler"
)

func main() {
	s := scheduler.Scheduler{}

	s.AddTask(scheduler.Task{ID: 1, Name: "Read File"})
	s.AddTask(scheduler.Task{ID: 2, Name: "Process Data "})
	s.AddTask(scheduler.Task{ID: 3, Name: "Report"})
	s.AddTask(scheduler.Task{ID: 4, Name: "Send Result"})

	fmt.Println("All tasks will start in order...")
	s.RunAll()
	fmt.Println("Scheduler completed.")
}
