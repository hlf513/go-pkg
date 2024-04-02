package concurrency

import (
	"context"
)

func Run(ctx context.Context, opts ...Option) []error {
	opt := NewOptions(opts...)
	var cn = make(chan struct{}, opt.ConcurrentNum)
	var errors []error
	defer close(cn)

	cn <- struct{}{}

	for _, task := range opt.Tasks {
		opt.WG.Add(1)
		go func(task Tasker) {
			defer func() {
				opt.WG.Done()
				<-cn
			}()
			if err := task.Run(); err != nil {
				errors = append(errors, err)
			}
		}(task)
	}
	opt.WG.Wait()

	return errors
}
