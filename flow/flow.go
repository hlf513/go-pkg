package flow

import (
	"context"
	"errors"
)

// Flower: NodeExecutor(Tasker) -> NodeExecutor(Tasker) ...
type Flower interface {
	Run(ctx context.Context) error
}

type Flow struct {
	opts             *FlowOptions
	executedExecutor map[NodeExecutor]struct{}
}

func NewFlow(ctx context.Context, opts ...FlowOption) Flower {
	opt := NewFlowOptions(opts...)
	return &Flow{opts: opt}
}

func (f *Flow) Run(ctx context.Context) error {
	var (
		executor   NodeExecutor
		err        error
		flowStatus = f.opts.Task.GetStatus()
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

		nextStatus := f.opts.Task.GetStatus()

		// equal means the status is not changed by the executor, so we need to set the next status
		if flowStatus == nextStatus {
			nextStatus, err = f.getNextNodeStatus(flowStatus)
			if err != nil {
				return err
			}
			if err := f.opts.Task.SetStatus(nextStatus); err != nil {
				return err
			}
		}
		// for next loop
		flowStatus = nextStatus

		if executor, err = f.setNextNode(nextStatus); err != nil {
			return err
		}
	}
	return nil
}

func (f *Flow) getNextNodeStatus(status FlowStatus) (FlowStatus, error) {
	if nextStatus, ok := f.opts.StatusTrans[status]; ok {
		return nextStatus, nil
	}
	return FlowStatus(0), errors.New("failed to find next executor status in transition configuration")
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

func (f *Flow) setNextNode(flowStatus FlowStatus) (NodeExecutor, error) {
	nextExecutor, err := f.getNodeExecutor(flowStatus)
	if err != nil {
		return nil, err
	}
	return nextExecutor, nil
}

func (f *Flow) execexecutor(ctx context.Context, executor NodeExecutor) error {
	var err error
	if err = f.isExecutedExecutor(executor); err != nil {
		return err
	}
	if err = executor.Execute(ctx, f.opts.Task); err != nil {
		return err
	}
	return f.addExecutedExecutor(executor)
}
