package flow

type FlowStatus int

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

func WithTask(task Tasker) FlowOption {
	return func(f *FlowOptions) {
		f.Task = task
	}
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
