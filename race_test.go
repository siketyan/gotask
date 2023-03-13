package gotask

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestRace(t *testing.T) {
	task := Race(
		NewTask(func(ctx context.Context) Result[string, error] {
			time.Sleep(1000 * time.Millisecond)

			return ResultOk[string, error]("The first task is resolved!")
		}),
		NewTask(func(ctx context.Context) Result[string, error] {
			time.Sleep(500 * time.Millisecond)

			return ResultOk[string, error]("The second task is resolved!")
		}),
	)

	ctx := context.Background()
	startedAt := time.Now()
	actual := task.Do(ctx).Unwrap()
	finishedAt := time.Now()

	assert.True(t, finishedAt.Sub(startedAt) < 600*time.Millisecond)
	assert.Equal(t, "The second task is resolved!", actual)
}
