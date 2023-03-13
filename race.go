package gotask

import (
	"context"
)

// Race creates a task that runs the children in parallel and returns the value resolved by any task at first.
func Race[T any](tasks ...Task[T]) Task[T] {
	return NewTask(func(ctx context.Context) T {
		valueChan := make(chan T)

		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		for _, task := range tasks {
			task.DoAsync(ctx, valueChan)
		}

		return <-valueChan
	})
}
