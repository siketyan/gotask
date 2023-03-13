package gotask

import (
	"context"
)

type Task[T any] struct {
	fn func(context.Context) T
}

// NewTask creates a new task from the closure.
func NewTask[T any](fn func(context.Context) T) Task[T] {
	return Task[T]{
		fn: fn,
	}
}

// TasksFrom creates tasks from the closures.
func TasksFrom[T any](fns ...func(context.Context) T) []Task[T] {
	tasks := make([]Task[T], 0, len(fns))
	for _, fn := range fns {
		tasks = append(tasks, NewTask(fn))
	}

	return tasks
}

// Do runs the task synchronously, blocking the current thread.
func (t Task[T]) Do(ctx context.Context) T {
	return t.fn(ctx)
}

// DoAsync runs the task asynchronously, writing the resolved value onto the channels.
func (t Task[T]) DoAsync(ctx context.Context, valueChan chan<- T) {
	go func() {
		valueChan <- t.Do(ctx)
	}()
}
