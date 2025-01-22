package flow

import (
	"context"
)

type FlowStatus int

type Tasker interface {
	GetStatus() FlowStatus
}

type NodeExecutor interface {
	Execute(ctx context.Context, task Tasker) error
}

type FlowOptions struct {
	StatusTrans   map[FlowStatus]FlowStatus
	NodeExecutors map[FlowStatus]NodeExecutor
	Task          Tasker
}

type FlowOption func(*FlowOptions)

func NewFlowOptions(opts ...FlowOption) *FlowOptions {
	f := &FlowOptions{}
	for _, opt := range opts {
		opt(f)
	}
	return f
}

func WithNodeExecutors(NodeExecutors map[FlowStatus]NodeExecutor) FlowOption {
	return func(f *FlowOptions) {
		f.NodeExecutors = NodeExecutors
	}
}

func WithStatusTrans(statusTrans map[FlowStatus]FlowStatus) FlowOption {
	return func(f *FlowOptions) {
		f.StatusTrans = statusTrans
	}
}

func WithTask(task Tasker) FlowOption {
	return func(f *FlowOptions) {
		f.Task = task
	}
}
