package concurrency

import (
	"context"
	"sync/atomic"
	"testing"
)

var cnt int32

type Task struct {
	id int
}

func (t *Task) Run() error {
	atomic.AddInt32(&cnt, 1)
	return nil
}

func NewTask(id int) *Task {
	return &Task{id: id}
}

func TestRun(t *testing.T) {
	// init concurrent tasks
	var tasks []Tasker
	for i := 0; i < 1000; i++ {
		tasks = append(tasks, NewTask(i))
	}

	// run concurrent tasks
	Run(
		context.Background(),
		ConcurrentNum(100), // set concurrent num
		Tasks(tasks),       // register concurrent tasks
	)

	// check concurrent tasks
	if cnt != 1000 {
		t.Error("the concurrent task failed")
	}
}
