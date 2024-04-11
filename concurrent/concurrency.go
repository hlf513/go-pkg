package concurrent

import (
	"context"

	"github.com/schollz/progressbar/v3"
)

func Executor(ctx context.Context, opts ...Option) []error {
	opt := NewOptions(opts...)
	var ch = make(chan struct{}, opt.ConcurrentNum)
	var errors []error
	defer close(ch)

	var bar *progressbar.ProgressBar
	if opt.ShowProgressbar {
		bar = progressbar.Default(int64(len(opt.Tasks)))
	}

	for _, task := range opt.Tasks {
		ch <- struct{}{}
		opt.WG.Add(1)
		go func(task Tasker) {
			defer func() {
				opt.WG.Done()
				<-ch
			}()
			if opt.ShowProgressbar {
				bar.Add(1)
			}
			if err := task.Execute(); err != nil {
				errors = append(errors, err)
			}
		}(task)
	}
	opt.WG.Wait()

	return errors
}
