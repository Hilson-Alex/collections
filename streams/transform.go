package streams

// Transform is used for transformation in collections.
// Changes a Stream of type T into the type U or into
// a Stream of U.
type Transform[T, U any] Stream[T]

// AsTransform can be combined with Stream.Wrap to turn a Stream
// with type T into a Transform from type T to type U.
func AsTransform[T, U any](stream []T) interface{} {
	return Transform[T, U](stream)
}

// Map changes a Transform based on a Stream[T] into a Stream[U] using the
// mapper function.
// The mapper receives an element and its index, then return the new value.
func (transform Transform[T, U]) Map(mapper streamFunc[T, U]) Stream[U] {
	var newArr = Stream[U]{}
	for index, item := range transform {
		newArr = append(newArr, mapper(item, index))
	}
	return newArr
}

// FlatMap changes a Transform based on a Stream[T] into a Stream[U] using the
// mapper function and then flat the result in one level.
// The mapper receives an element and its index, then return an array of items.
func (transform Transform[T, U]) FlatMap(mapper streamFunc[T, []U]) Stream[U] {
	var newArr = Stream[U]{}
	for index, item := range transform {
		newArr = append(newArr, mapper(item, index)...)
	}
	return newArr
}

// Reduce reduces the transform from a given initial value and aggregate the stream
// values using an associative callback.
// The reducer function receives the previous result (starting with the initial value),
// the current element and its index, then return the new value for the next iteration.
func (transform Transform[T, U]) Reduce(reducer func(U, T, int) U, initialValue U) U {
	var acc = initialValue
	for index, item := range transform {
		acc = reducer(acc, item, index)
	}
	return acc
}
