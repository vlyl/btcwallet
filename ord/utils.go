package ord

import "os"

type Result[T any] struct {
	value T
	err   error
}

func (r Result[T]) Ok() bool {
	return r.err == nil
}

func (r Result[T]) Err() error {
	return r.err
}

func (r Result[T]) Unwrap() T {
	if r.err != nil {
		panic(r.err)
	}

	return r.value
}

func (r Result[T]) Expect(err error) T {
	if r.err != nil {
		panic(err)
	}

	return r.value
}

func New[T any](v T, err error) Result[T] {
	return Result[T]{
		value: v,
		err:   err,
	}
}

func Match[T any](r Result[T], okF func(T), errF func(error)) {
	if r.err != nil {
		errF(r.err)
	} else {
		okF(r.value)
	}
}

func FileExist(f string) bool {
	_, err := os.Stat(f)
	return !os.IsNotExist(err)
}

type FeeRate uint64
