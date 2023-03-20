package gotask

type Result[T any, E comparable] struct {
	ok  T
	err E
}

func NewResult[T any, E comparable](ok T, err E) Result[T, E] {
	return Result[T, E]{
		ok:  ok,
		err: err,
	}
}

func ResultOk[T any, E comparable](ok T) Result[T, E] {
	return Result[T, E]{
		ok: ok,
	}
}

func ResultErr[T any, E comparable](err E) Result[T, E] {
	return Result[T, E]{
		err: err,
	}
}

func UnwrapMany[T any, E comparable](results ...Result[T, E]) ([]T, []E) {
	values := make([]T, 0, len(results))
	errs := make([]E, 0, len(results))
	for _, result := range results {
		if result.IsOk() {
			values = append(values, result.Unwrap())
		} else {
			errs = append(errs, result.UnwrapErr())
		}
	}

	return values, errs
}

func (r Result[T, E]) IsOk() bool {
	return isZero(r.err)
}

func (r Result[T, E]) IsErr() bool {
	return !isZero(r.err)
}

func (r Result[T, E]) Unwrap() T {
	if r.IsErr() {
		panic("an error result was unwrapped as ok")
	}

	return r.ok
}

func (r Result[T, E]) UnwrapErr() E {
	if r.IsOk() {
		panic("an ok result was unwrapped as error")
	}

	return r.err
}

func (r Result[T, E]) AsTuple() (T, E) {
	return r.ok, r.err
}
