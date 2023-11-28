package gotask

import (
	"context"
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
)

func TestCancelableTask_Cancel(t *testing.T) {
	startedAt := time.Now()
	task := NewTask(func(ctx context.Context) Unit {
		time.Sleep(10 * time.Second)

		return U
	}).Cancelable()

	task.DoAsync(context.Background(), nil)
	task.Cancel()

	assert.Less(t, time.Now().Sub(startedAt), 5*time.Second)
}
