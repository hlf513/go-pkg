package flow

import (
	"context"
	"errors"
)

type Flow struct {
	opts             *FlowOptions
	executedExecutor map[NodeExecutor]struct{}
}

func NewFlow(ctx context.Context, opts ...FlowOption) *Flow {
	opt := NewFlowOptions(opts...)
	return &Flow{opts: opt}
}

func (f *Flow) Run(ctx context.Context, flowStatus FlowStatus) error {
	var (
		executor NodeExecutor
		err      error
	)
	if err = f.init(); err != nil {
		return err
	}
	if executor, err = f.getNodeExecutor(flowStatus); err != nil {
		return err
	}
	for executor != nil {
		if err = f.execexecutor(ctx, executor); err != nil {
			return err
		}
		if flowStatus, executor, err = f.setNextNode(flowStatus); err != nil {
			return err
		}
	}
	return nil
}

func (f *Flow) init() error {
	f.clearExecutedExecutor()
	return nil
}

func (f *Flow) isExecutedExecutor(executor NodeExecutor) error {
	if _, ok := f.executedExecutor[executor]; ok {
		return errors.New("executor has already been executed")
	}
	return nil
}

func (f *Flow) addExecutedExecutor(executor NodeExecutor) error {
	f.executedExecutor[executor] = struct{}{}
	return nil
}

func (f *Flow) clearExecutedExecutor() {
	f.executedExecutor = make(map[NodeExecutor]struct{})
}

func (f *Flow) getNodeExecutor(status FlowStatus) (NodeExecutor, error) {
	if executor, ok := f.opts.NodeExecutors[status]; ok {
		return executor, nil
	}
	// Return nil when executor not found, which is expected for the last node
	return nil, nil
}

func (f *Flow) getNextNodeStatus(status FlowStatus) (FlowStatus, error) {
	if nextStatus, ok := f.opts.StatusTrans[status]; ok {
		return nextStatus, nil
	}
	return FlowStatus(0), errors.New("failed to find next executor status in transition configuration")
}

func (f *Flow) setNextNode(flowStatus FlowStatus) (FlowStatus, NodeExecutor, error) {
	nextStatus, err := f.getNextNodeStatus(flowStatus)
	if err != nil {
		return FlowStatus(0), nil, err
	}
	nextExecutor, err := f.getNodeExecutor(nextStatus)
	if err != nil {
		return FlowStatus(0), nil, err
	}
	return nextStatus, nextExecutor, nil
}

func (f *Flow) execexecutor(ctx context.Context, executor NodeExecutor) error {
	var err error
	if err = f.isExecutedExecutor(executor); err != nil {
		return err
	}
	if err = executor.Execute(ctx); err != nil {
		return err
	}
	return f.addExecutedExecutor(executor)
}
