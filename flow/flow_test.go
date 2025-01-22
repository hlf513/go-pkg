package flow

import (
	"context"
	"fmt"
	"testing"
)

type oneExecutorOptions struct {
	Id int
}

type oneExecutorOption func(*oneExecutorOptions)

func newOneExecutorOptions(opts ...oneExecutorOption) *oneExecutorOptions {
	opt := &oneExecutorOptions{}
	for _, o := range opts {
		o(opt)
	}
	return opt
}

type oneExecutor struct {
	opts *oneExecutorOptions
}

func (e *oneExecutor) Execute(ctx context.Context, task Tasker) error {
	fmt.Println("OneExecutor")
	return nil
}

func newOneExecutor(opts ...oneExecutorOption) *oneExecutor {
	opt := newOneExecutorOptions(opts...)
	return &oneExecutor{opts: opt}
}

type twoExecutor struct {
}

func (e *twoExecutor) Execute(ctx context.Context, task Tasker) error {
	fmt.Println("TwoExecutor")
	return nil
}

type task struct {
}

func (t *task) GetStatus() FlowStatus {
	return FlowStatus(0)
}

func TestFlow(t *testing.T) {
	flow := NewFlow(context.Background(),
		WithNodeExecutors(map[FlowStatus]NodeExecutor{
			FlowStatus(0): newOneExecutor(),
			FlowStatus(1): &twoExecutor{},
		}),
		WithStatusTrans(map[FlowStatus]FlowStatus{
			FlowStatus(0): FlowStatus(1),
			FlowStatus(1): FlowStatus(2),
		}),
		WithTask(new(task)),
	)

	err := flow.Run(context.Background())
	if err != nil {
		t.Error(err)
	}
}
