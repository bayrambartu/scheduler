package scheduler

import (
	"fmt"
	"sync"
	"time"
)

type Task struct {
	ID   int
	Name string
}

func (t *Task) Run(mu *sync.Mutex, wg *sync.WaitGroup) {
	defer wg.Done()

	fmt.Printf("Task %d (%s) waiting...\n", t.ID, t.Name)
	mu.Lock()
	fmt.Printf("Task %d (%s) running...\n", t.ID, t.Name)
	time.Sleep(1 * time.Second)
	fmt.Printf("Task %d (%s) done.\n", t.ID, t.Name)
	mu.Unlock()
}
