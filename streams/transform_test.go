package streams

import (
	"reflect"
	"strconv"
	"testing"
)

func TestTransform_Map(t *testing.T) {
	var arr = []int{10, 20, 50, 30, 70, 120}
	var expected = Stream[string]{"1", "2", "5", "3", "7", "12"}
	var result = Transform[int, int](arr).
		Map(func(i, _ int) int {
			return i / 10
		}).
		Wrap(AsTransform[int, string]).(Transform[int, string]).
		Map(func(i int, _ int) string {
			return strconv.Itoa(i)
		})
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("expected: %q, received: %q", expected, result)
	}
}

func TestTransform_FlatMap(t *testing.T) {
	var arr = [][]int{
		{10, 20},
		{50, 30, 70},
		{120},
	}
	var expected = []int{10, 20, 50, 30, 70, 120}
	var result = Transform[[]int, int](arr).FlatMap(func(ints []int, _ int) []int {
		return ints
	})
	if !reflect.DeepEqual(expected, []int(result)) {
		t.Fatalf("expected: %v, received: %v", expected, result)
	}
}

func TestTransform_Reduce(t *testing.T) {
	var arr = []int{1, 10, 4, 5, 7, 20, 3}
	var sum = Transform[int, string](arr).Reduce(func(acc string, value int, _ int) string {
		return acc + strconv.Itoa(value)
	}, "")
	if sum != "110457203" {
		t.Fatalf("expected: %v, received: %v", "110457203", sum)
	}
}
