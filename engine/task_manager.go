package engine

import (
	"context"
	"errors"
	"sync"
)

type Task struct {
	ID       string
	url      string
	progress int
	done     bool
}

type TaskManager struct {
	wg          sync.WaitGroup
	taskQueue   chan *Task
	workerCount int
	ctx         context.Context
	cancelFn    context.CancelFunc
}

func NewTaskManager(workerCount int, queueSize int) *TaskManager {
	tm := &TaskManager{
		taskQueue:   make(chan *Task, queueSize),
		workerCount: workerCount,
	}
	tm.startWorkers()
	return tm
}

// Add a task to the task queue. Returns an error if the
// queue is full, otherwise nil.
func (tm *TaskManager) AddTask(task *Task) error {
	// Non blocking send to task queue channel.
	select {
	case tm.taskQueue <- task:
		return nil
	default:
		return errors.New("Full")
	}
}

func (tm *TaskManager) Shutdown() {
	// Send signal to workers to stop
	close(tm.taskQueue)
	tm.wg.Wait()
}

func (tm *TaskManager) startWorkers() {
	for i := 0; i < tm.workerCount; i++ {
		tm.wg.Add(1)
		go tm.worker()
	}
}

func (tm *TaskManager) worker() {
	defer tm.wg.Done()
	for task := range tm.taskQueue {
		tm.processTask(task)
	}
}

func (tm *TaskManager) processTask(task *Task) error {
	// do something with task
	return nil
}
