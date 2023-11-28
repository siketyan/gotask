package gotask

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestUnwrapMany(t *testing.T) {
	values, errs := UnwrapMany(
		ResultOk("abc"),
		ResultErr[string](errors.New("error occurred (1)")),
		ResultOk("def"),
		ResultErr[string](errors.New("error occurred (2)")),
		ResultOk("ghi"),
	)

	assert.Len(t, values, 3)
	assert.Len(t, errs, 2)
}
