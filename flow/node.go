package flow

import "context"

type NodeExecutor interface {
	Execute(ctx context.Context, task Tasker) error
}
