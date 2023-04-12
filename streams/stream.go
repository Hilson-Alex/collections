package streams

type Stream[T any] []T

type Transform[T, U any] Stream[T]

type wrapFunc[T any] func([]T) interface{}

type streamFunc[T, U any] func(T, int) U

func AsTransform[T, U any](stream []T) interface{} {
	return Transform[T, U](stream)
}

func (stream Stream[T]) Foreach(callback func(T, int)) {
	for index, item := range stream {
		callback(item, index)
	}
}

func (stream Stream[T]) Filter(callback streamFunc[T, bool]) Stream[T] {
	var newArr = Stream[T]{}
	for index, item := range stream {
		if callback(item, index) {
			newArr = append(newArr, item)
		}
	}
	return newArr
}

func (stream Stream[T]) Some(callback streamFunc[T, bool]) bool {
	for index, item := range stream {
		if callback(item, index) {
			return true
		}
	}
	return false
}

func (stream Stream[T]) Every(callback streamFunc[T, bool]) bool {
	for index, item := range stream {
		if !callback(item, index) {
			return false
		}
	}
	return true
}

func (stream Stream[T]) Wrap(wrapper wrapFunc[T]) interface{} {
	return wrapper(stream)
}

func (transform Transform[T, U]) Map(mapper streamFunc[T, U]) Stream[U] {
	var newArr = Stream[U]{}
	for index, item := range transform {
		newArr = append(newArr, mapper(item, index))
	}
	return newArr
}

func (transform Transform[T, U]) FlatMap(mapper streamFunc[T, []U]) Stream[U] {
	var newArr = Stream[U]{}
	for index, item := range transform {
		newArr = append(newArr, mapper(item, index)...)
	}
	return newArr
}

func (transform Transform[T, U]) Reduce(reducer func(U, T, int) U, initialValue U) U {
	var acc = initialValue
	for index, item := range transform {
		acc = reducer(acc, item, index)
	}
	return acc
}
