package utils

import (
	"sync"
)

type Task interface {
	Process(workerId int)
}

type WorkerPool struct {
	Tasks       []Task
	Concurrency int
	TasksChan   chan Task
	wg          sync.WaitGroup
}

func (wp *WorkerPool) Worker(workerId int) {
	for task := range wp.TasksChan {
		task.Process(workerId)
		wp.wg.Done()
	}
}

func (wp *WorkerPool) Run() {
	wp.TasksChan = make(chan Task, len(wp.Tasks))

	for i := 0; i < wp.Concurrency; i++ {
		go wp.Worker(i)
	}

	wp.wg.Add(len(wp.Tasks))
	for _, task := range wp.Tasks {
		wp.TasksChan <- task
	}
	close(wp.TasksChan)

	wp.wg.Wait()
}
