package streams

import (
	"sort"
)

// Stream is the base type for functional methods. Used for methods that either return
// a specified type (like boolean) or don't change the type at all.
// The stream methods never change the original stream, but return a new.
type Stream[T any] []T

// wrapFunc defines a callback to change the Stream type
// without losing the fluent interface.
type wrapFunc[T any] func([]T) interface{}

// streamFunc is a callback that receives an item from
// the stream and its index and then returns something.
type streamFunc[T, U any] func(T, int) U

// Foreach applies the callback to each item of the stream, but doesn't return anything.
// The callback receives an item from the stream and its index.
func (stream Stream[T]) Foreach(callback func(T, int)) {
	for index, item := range stream {
		callback(item, index)
	}
}

// Some returns true if any of the items are approved by the callback.
// The callback receives an item and its index, then returns a boolean.
func (stream Stream[T]) Some(callback streamFunc[T, bool]) bool {
	for index, item := range stream {
		if callback(item, index) {
			return true
		}
	}
	return false
}

// None returns true if none of the items are approved by the callback.
// The callback receives an item and its index, then returns a boolean.
func (stream Stream[T]) None(callback streamFunc[T, bool]) bool {
	for index, item := range stream {
		if callback(item, index) {
			return false
		}
	}
	return true
}

// Every returns true if every item is approved by the callback.
// The callback receives an item and its index, then returns a boolean.
func (stream Stream[T]) Every(callback streamFunc[T, bool]) bool {
	for index, item := range stream {
		if !callback(item, index) {
			return false
		}
	}
	return true
}

// Reduce reduces from the first element using an associative callback. If you want
// the return to have a different type from the elements on the Stream, then check
// Transform.Reduce.
// the callback receives the accumulator, one element from the array, and its index.
// Then, returns the result of the reducer function.
func (stream Stream[T]) Reduce(reducer func(T, T, int) T) T {
	var acc = stream[0]
	for index, item := range stream[1:] {
		acc = reducer(acc, item, index+1)
	}
	return acc
}

// Concat adds the elements to the end of the stream.
func (stream Stream[T]) Concat(elements ...T) Stream[T] {
	return append(cloneStream(stream), elements...)
}

// Prepend adds the elements to the beginning of the stream.
func (stream Stream[T]) Prepend(elements ...T) Stream[T] {
	return append(cloneStream(elements), stream...)
}

// AddAt adds the elements on the given index.
func (stream Stream[T]) AddAt(index int, elements ...T) Stream[T] {
	return append(stream[:index].Concat(elements...), stream[index:]...)
}

// With replaces the element on the given index.
func (stream Stream[T]) With(index int, element T) Stream[T] {
	var newArr Stream[T] = cloneStream(stream)
	newArr[index] = element
	return newArr
}

// Map changes a Stream[T] in another Stream[T] applying a mapper function to
// each element of the Stream. If you want to change the type of the items on the
// Stream then check Transform.Map.
// The mapper function receives an element and its index, then return the new value.
func (stream Stream[T]) Map(mapper streamFunc[T, T]) Stream[T] {
	var newArr []T
	for index, item := range stream {
		newArr = append(newArr, mapper(item, index))
	}
	return newArr
}

// Filter returns a stream consisting of the items approved  by the callback.
// The callback receives an item and its index, then returns a boolean.
func (stream Stream[T]) Filter(callback streamFunc[T, bool]) Stream[T] {
	var newArr []T
	for index, item := range stream {
		if callback(item, index) {
			newArr = append(newArr, item)
		}
	}
	return newArr
}

// Sort the stream. If true, i comes before j.
// The callback receives two items, then returns a boolean.
func (stream Stream[T]) Sort(callback func(i, j T) bool) Stream[T] {
	var newArr Stream[T] = cloneStream(stream)
	sort.Slice(newArr, func(i, j int) bool {
		return callback(newArr[i], newArr[j])
	})
	return newArr
}

// Wrap changes the stream into anything without losing the fluent interface.
func (stream Stream[T]) Wrap(wrapper wrapFunc[T]) interface{} {
	return wrapper(stream)
}

// cloneStream creates a new reference with a copy to a slice.
func cloneStream[T any](slice []T) []T {
	var newArr = make([]T, len(slice))
	copy(newArr, slice)
	return newArr
}
