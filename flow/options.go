package flow

import (
	"context"
)

type FlowStatus int

type NodeExecutor interface {
	Execute(ctx context.Context) error
}

type FlowOptions struct {
	StatusTrans   map[FlowStatus]FlowStatus
	NodeExecutors map[FlowStatus]NodeExecutor
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
