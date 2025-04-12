package main

import (
	"fmt"
	"sort"
	"time"
)

type Task struct {
	ID       int
	Name     string
	Priority int
}

func (t Task) Run() {
	fmt.Printf("Task %s (priority: %d) is running\n", t.Name, t.Priority)
	time.Sleep(1 * time.Second)
	fmt.Printf("FINISHED %s \n", t.Name)
}

type Scheduler struct {
	task []Task
}

func (s *Scheduler) AddTask(t Task) {
	s.task = append(s.task, t)
}

func (s *Scheduler) RunAll() {
	sort.Slice(s.task, func(i, j int) bool {
		return s.task[i].Priority > s.task[j].Priority
	})

	for _, task := range s.task {
		task.Run()
	}
}

func main() {
	s := Scheduler{}

	s.AddTask(Task{ID: 1, Name: "Clear Logs", Priority: 1})
	s.AddTask(Task{ID: 2, Name: "Create Report", Priority: 3})
	s.AddTask(Task{ID: 3, Name: "Backup Database", Priority: 5})
	s.AddTask(Task{ID: 4, Name: "API Test", Priority: 2})

	fmt.Println("Start tasks by priority...")
	s.RunAll()
	fmt.Println("Finished all tasks.")
}
