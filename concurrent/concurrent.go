package concurrent

import (
	"context"

	"github.com/schollz/progressbar/v3"
	"golang.org/x/sync/errgroup"
)

func Executor(ctx context.Context, opts ...Option) error {
	opt := NewOptions(opts...)
	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(opt.MaxG)

	var bar *progressbar.ProgressBar
	if opt.ShowProgressbar {
		bar = progressbar.Default(int64(len(opt.Tasks)))
	}

	for _, task := range opt.Tasks {
		task := task
		g.Go(func() error {
			if opt.ShowProgressbar {
				bar.Add(1)
			}
			if err := task.Execute(ctx); err != nil {
				return err
			}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}

func ExecutorInterrupt(ctx context.Context, opts ...Option) error {
	opt := NewOptions(opts...)
	g, ctx := errgroup.WithContext(ctx)
	g.SetLimit(opt.MaxG)

	var bar *progressbar.ProgressBar
	if opt.ShowProgressbar {
		bar = progressbar.Default(int64(len(opt.Tasks)))
	}

	for _, task := range opt.Tasks {
		task := task
		g.Go(func() error {
			if opt.ShowProgressbar {
				bar.Add(1)
			}
			select {
			case <-ctx.Done():
				return ctx.Err()
			default:
			}

			if err := task.Execute(ctx); err != nil {
				return err
			}
			return nil
		})
	}
	if err := g.Wait(); err != nil {
		return err
	}

	return nil
}
