package gotask

import (
	"context"
)

// Then creates a new task that runs two tasks serially.
func Then[T, U any](
	task Task[Result[T]],
	fn func(context.Context, T) Result[U],
) Task[Result[U]] {
	return NewTask(func(ctx context.Context) Result[U] {
		result := task.Do(ctx)
		if result.IsErr() {
			return ResultErr[U](result.UnwrapErr())
		}

		return fn(ctx, result.Unwrap())
	})
}

// Catch creates a new task that fallbacks to onError behaviour on the task returned an error.
func Catch[T any](
	task Task[Result[T]],
	onError func(context.Context, error) T,
) Task[T] {
	return NewTask(func(ctx context.Context) T {
		result := task.Do(ctx)
		if result.IsErr() {
			return onError(ctx, result.UnwrapErr())
		}

		return result.Unwrap()
	})
}
