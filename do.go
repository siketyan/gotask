package gotask

import (
	"context"
)

// DoAll is the shorthand of Parallel and Task.Do.
func DoAll[T any](ctx context.Context, tasks ...Task[Result[T]]) Result[[]T] {
	return Parallel(tasks...).Do(ctx)
}

// DoAllSettled is the shorthand of ParallelSettled and Task.Do.
func DoAllSettled[T any](ctx context.Context, tasks ...Task[T]) []T {
	return ParallelSettled(tasks...).Do(ctx)
}

// DoRace is the shorthand of Race and Task.Do.
func DoRace[T any](ctx context.Context, tasks ...Task[T]) T {
	return Race(tasks...).Do(ctx)
}

// DoAllFns is the shorthand of TasksFrom + Parallel + Task.Do.
func DoAllFns[T any, E comparable](ctx context.Context, fns ...func(context.Context) Result[T]) Result[[]T] {
	return DoAll(ctx, TasksFrom(fns...)...)
}

// DoAllFnsSettled is the shorthand of TasksFrom + ParallelSettled + Task.Do.
func DoAllFnsSettled[T any](ctx context.Context, fns ...func(context.Context) T) []T {
	return DoAllSettled(ctx, TasksFrom(fns...)...)
}

// DoRaceFns is the shorthand of TasksFrom + Race + Task.Do.
func DoRaceFns[T any](ctx context.Context, fns ...func(context.Context) T) T {
	return DoRace(ctx, TasksFrom(fns...)...)
}
