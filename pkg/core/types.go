package core

type Type interface {
}

type Array[T any] struct {
	t T
}
