package scheduler

import "sync"

type Scheduler struct {
	Tasks []Task
	Lock  sync.Mutex
}

func (s *Scheduler) AddTask(task Task) {
	s.Tasks = append(s.Tasks, task)
}

func (s *Scheduler) RunAll() {
	var wg sync.WaitGroup
	for _, task := range s.Tasks {
		wg.Add(1)
		go task.Run(&s.Lock, &wg)
	}
	wg.Wait()
}
