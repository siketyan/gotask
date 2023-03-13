package gotask

import (
	"context"
)

// Parallel creates a task that runs the children in parallel, and returns all values resolved by each child task.
// On any task returned an error, the remaining tasks will be canceled immediately and returns the error.
// This is only compatible with tasks that returns Result.
func Parallel[T any, E comparable](tasks ...Task[Result[T, E]]) Task[Result[[]T, E]] {
	return NewTask(func(ctx context.Context) Result[[]T, E] {
		resultChan := make(chan Result[T, E])

		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		for _, task := range tasks {
			task.DoAsync(ctx, resultChan)
		}

		values := make([]T, 0, len(tasks))
		for {
			result := <-resultChan
			if result.IsErr() {
				return ResultErr[[]T, E](result.UnwrapErr())
			}

			values = append(values, result.Unwrap())
			if len(tasks) <= len(values) {
				break
			}
		}

		return ResultOk[[]T, E](values)
	})
}

// ParallelSettled creates a task that runs the children in parallel, and returns all values resolved by each child task.
// The run will be continued even if any child task is returned an error, and returns all the resolved values.
// This is compatible with tasks that does not return Result.
func ParallelSettled[T any](tasks ...Task[T]) Task[[]T] {
	return NewTask(func(ctx context.Context) []T {
		valueChan := make(chan T)

		ctx, cancel := context.WithCancel(ctx)
		defer cancel()

		for _, task := range tasks {
			task.DoAsync(ctx, valueChan)
		}

		results := make([]T, 0, len(tasks))
		for {
			if len(tasks) <= len(results) {
				break
			}

			results = append(results, <-valueChan)
		}

		return results
	})
}
