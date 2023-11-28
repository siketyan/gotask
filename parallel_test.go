package gotask

import (
	"context"
	"errors"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestParallel(t *testing.T) {
	task := Parallel(
		NewTask(func(ctx context.Context) Result[string] {
			time.Sleep(900 * time.Millisecond)

			return ResultOk("The first task is resolved!")
		}),
		NewTask(func(ctx context.Context) Result[string] {
			time.Sleep(1000 * time.Millisecond)

			return ResultOk("The second task is resolved!")
		}),
	)

	ctx := context.Background()
	startedAt := time.Now()
	actual := task.Do(ctx).Unwrap()
	finishedAt := time.Now()

	assert.True(t, finishedAt.Sub(startedAt) < 1100*time.Millisecond)
	assert.Equal(t, []string{
		"The first task is resolved!",
		"The second task is resolved!",
	}, actual)
}

func TestParallel_Error(t *testing.T) {
	task := Parallel(
		NewTask(func(ctx context.Context) Result[string] {
			time.Sleep(900 * time.Millisecond)

			return ResultOk("The first task is resolved!")
		}),
		NewTask(func(ctx context.Context) Result[string] {
			time.Sleep(1000 * time.Millisecond)

			return ResultErr[string](errors.New("the second task occurred an error"))
		}),
	)

	ctx := context.Background()
	actual := task.Do(ctx)

	assert.True(t, actual.IsErr())
	assert.Equal(t, "the second task occurred an error", actual.UnwrapErr().Error())
}

func TestParallelSettled(t *testing.T) {
	task := ParallelSettled(
		NewTask(func(ctx context.Context) Result[string] {
			time.Sleep(800 * time.Millisecond)

			return ResultOk("The first task is resolved!")
		}),
		NewTask(func(ctx context.Context) Result[string] {
			time.Sleep(900 * time.Millisecond)

			return ResultErr[string](errors.New("the second task occurred an error"))
		}),
		NewTask(func(ctx context.Context) Result[string] {
			time.Sleep(1000 * time.Millisecond)

			return ResultOk("The third task is resolved!")
		}),
	)

	ctx := context.Background()
	startedAt := time.Now()
	actual := task.Do(ctx)
	finishedAt := time.Now()

	assert.True(t, finishedAt.Sub(startedAt) < 1100*time.Millisecond)
	assert.Equal(t, []Result[string]{
		ResultOk("The first task is resolved!"),
		ResultErr[string](errors.New("the second task occurred an error")),
		ResultOk("The third task is resolved!"),
	}, actual)
}
