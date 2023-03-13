package gotask

import (
	"context"
)

// Then creates a new task that runs two tasks serially.
func Then[T, U any, E comparable](
	task Task[Result[T, E]],
	fn func(context.Context, T) Result[U, E],
) Task[Result[U, E]] {
	return NewTask(func(ctx context.Context) Result[U, E] {
		result := task.Do(ctx)
		if result.IsErr() {
			return ResultErr[U, E](result.UnwrapErr())
		}

		return fn(ctx, result.Unwrap())
	})
}

// Catch creates a new task that fallbacks to onError behaviour on the task returned an error.
func Catch[T any, E comparable](
	task Task[Result[T, E]],
	onError func(context.Context, E) T,
) Task[T] {
	return NewTask(func(ctx context.Context) T {
		result := task.Do(ctx)
		if result.IsErr() {
			return onError(ctx, result.UnwrapErr())
		}

		return result.Unwrap()
	})
}
