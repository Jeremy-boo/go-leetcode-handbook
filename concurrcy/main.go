package concurrcy

import (
	"context"
)

type Task struct {
}

func (t *Task) Do(ctx context.Context) error {

	return nil
}

func ConcurrentDoTask(ctx context.Context, tasks []*Task, concurrency int64) error {
	errChan := make(chan error)
	cChan := make(chan int, concurrency)
	// todo task中拿出任务
	for _, task := range tasks {
		go func(ctx context.Context) {
			cChan <- 1
			err := task.Do(ctx)
			if err != nil {
				errChan <- err
			}
			// 从channel 消费一个数据
			ctx.Done()
			<-cChan
		}(ctx)
	}
	select {
	case err := <-errChan:
		return err
	default:
		return nil
	}
}
