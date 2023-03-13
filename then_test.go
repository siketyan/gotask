package gotask

import (
	"context"
	"errors"
	"fmt"
	"strconv"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestThen(t *testing.T) {
	task := Then(
		NewTask(func(ctx context.Context) Result[string, error] {
			return ResultOk[string, error]("123")
		}),
		func(ctx context.Context, value string) Result[int, error] {
			return NewResult(strconv.Atoi(value))
		},
	)

	ctx := context.Background()
	value := task.Do(ctx).Unwrap()
	assert.Equal(t, 123, value)
}

func TestCatch(t *testing.T) {
	task := Catch(
		NewTask(func(ctx context.Context) Result[string, error] {
			return ResultErr[string, error](errors.New("error occurred"))
		}),
		func(ctx context.Context, err error) string {
			return fmt.Sprintf("ERROR: %s", err.Error())
		},
	)

	ctx := context.Background()
	value := task.Do(ctx)
	assert.Equal(t, "ERROR: error occurred", value)
}
