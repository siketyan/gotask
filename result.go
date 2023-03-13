package gotask

type Result[T any, E comparable] struct {
	ok  T
	err E
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
