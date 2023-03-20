package gotask

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnwrapMany(t *testing.T) {
	values, errs := UnwrapMany(
		ResultOk[string, error]("abc"),
		ResultErr[string, error](errors.New("error occurred (1)")),
		ResultOk[string, error]("def"),
		ResultErr[string, error](errors.New("error occurred (2)")),
		ResultOk[string, error]("ghi"),
	)

	assert.Len(t, values, 3)
	assert.Len(t, errs, 2)
}
