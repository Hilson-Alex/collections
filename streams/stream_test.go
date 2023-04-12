package streams

import (
	"reflect"
	"strconv"
	"testing"
)

func TestStream_Foreach(t *testing.T) {
	var arr = []float64{0.75, 2.45, 6.67, 7.90, 2.31}
	var index = 0
	var arrCopy []float64
	Stream[float64](arr).Foreach(func(f float64, i int) {
		if i != index {
			t.Fatalf("Wrong index. Expected %v, but received %v", index, i)
		}
		index++
		arrCopy = append(arrCopy, f)
	})
	if !reflect.DeepEqual(arrCopy, arr) {
		t.Fatalf("expected: %v, received: %v", arr, arrCopy)
	}
}

func TestStream_Filter(t *testing.T) {
	var arr = []int{1, 10, 4, 5, 7, 20, 3}
	var expected = []int{10, 7, 20}
	var filtered = Stream[int](arr).Filter(func(i int, _ int) bool {
		return i > 5
	})
	if !reflect.DeepEqual([]int(filtered), expected) {
		t.Fatalf("expected: %v, received: %v", expected, filtered)
	}
}

func TestStream_Every(t *testing.T) {
	var arr = []string{"aaaa", "bbbb", "cc", "dd", "eeee", "ffff"}
	var condition = Stream[string](arr).Every(func(i string, _ int) bool {
		return len(i) > 0
	})
	if !condition {
		t.FailNow()
	}
	condition = Stream[string](arr).Every(func(i string, _ int) bool {
		return len(i) == 4
	})
	if condition {
		t.FailNow()
	}
	condition = Stream[string](arr).Every(func(i string, _ int) bool {
		return i == "baba"
	})
	if condition {
		t.FailNow()
	}
}

func TestStream_Some(t *testing.T) {
	var arr = []string{"aaaa", "bbbb", "cc", "dd", "eeee", "ffff"}
	var condition = Stream[string](arr).Some(func(i string, _ int) bool {
		return len(i) > 0
	})
	if !condition {
		t.FailNow()
	}
	condition = Stream[string](arr).Some(func(i string, _ int) bool {
		return len(i) == 4
	})
	if !condition {
		t.FailNow()
	}
	condition = Stream[string](arr).Some(func(i string, _ int) bool {
		return i == "baba"
	})
	if condition {
		t.FailNow()
	}
}

func TestTransform_Map(t *testing.T) {
	var arr = []int{10, 20, 50, 30, 70, 120}
	var expected = []string{"1", "2", "5", "3", "7", "12"}
	var result = Transform[int, int](arr).
		Map(func(i int, _ int) int {
			return i / 10
		}).
		Wrap(AsTransform[int, string]).(Transform[int, string]).
		Map(func(i int, _ int) string {
			return strconv.Itoa(i)
		})
	if !reflect.DeepEqual(expected, []string(result)) {
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
	var sum = Transform[int, int](arr).Reduce(func(acc int, value int, _ int) int {
		return acc + value
	}, 0)
	if sum != 50 {
		t.Fatalf("expected: %v, received: %v", 50, sum)
	}
}
