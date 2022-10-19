package main

type Destructable interface {
	Destruct()
}

type DesWrapper[T any] struct {
	ptr *T
}

func (desWrapper DesWrapper[T]) Destruct() {
	Delete(desWrapper.ptr)
}

type DDesWrapper[T Destructable] struct {
	ptr *T
}

func (desWrapper DDesWrapper[T]) Destruct() {
	(*desWrapper.ptr).Destruct()
	Delete(desWrapper.ptr)
}
