package gotask

type Result[T any] struct {
	ok  T
	err error
}

func NewResult[T any](ok T, err error) Result[T] {
	return Result[T]{
		ok:  ok,
		err: err,
	}
}

func ResultOk[T any](ok T) Result[T] {
	return Result[T]{
		ok: ok,
	}
}

func ResultErr[T any](err error) Result[T] {
	return Result[T]{
		err: err,
	}
}

func UnwrapMany[T any](results ...Result[T]) ([]T, []error) {
	values := make([]T, 0, len(results))
	errs := make([]error, 0, len(results))
	for _, result := range results {
		if result.IsOk() {
			values = append(values, result.Unwrap())
		} else {
			errs = append(errs, result.UnwrapErr())
		}
	}

	return values, errs
}

func (r Result[T]) IsOk() bool {
	return isZero(r.err)
}

func (r Result[T]) IsErr() bool {
	return !isZero(r.err)
}

func (r Result[T]) Unwrap() T {
	if r.IsErr() {
		panic("an error result was unwrapped as ok")
	}

	return r.ok
}

func (r Result[T]) UnwrapErr() error {
	if r.IsOk() {
		panic("an ok result was unwrapped as error")
	}

	return r.err
}

func (r Result[T]) AsTuple() (T, error) {
	return r.ok, r.err
}
