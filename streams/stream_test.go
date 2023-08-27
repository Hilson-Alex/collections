package streams

import (
	"fmt"
	"reflect"
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

func TestStream_Some(t *testing.T) {
	var arr = Stream[string]{"aaaa", "bbbb", "cc", "dd", "eeee", "ffff"}
	var condition = arr.Some(func(i string, _ int) bool {
		return len(i) > 0
	})
	if !condition {
		t.FailNow()
	}
	condition = arr.Some(func(i string, _ int) bool {
		return len(i) == 2
	})
	if !condition {
		t.FailNow()
	}
	condition = arr.Some(func(i string, _ int) bool {
		return i == "baba"
	})
	if condition {
		t.FailNow()
	}
}

func TestStream_None(t *testing.T) {
	var arr = Stream[string]{"aaaa", "bbbb", "cc", "dd", "eeee", "ffff"}
	var condition = arr.None(func(i string, _ int) bool {
		return len(i) > 0
	})
	if condition {
		t.FailNow()
	}
	condition = arr.None(func(i string, _ int) bool {
		return len(i) == 2
	})
	if condition {
		t.FailNow()
	}
	condition = arr.None(func(i string, _ int) bool {
		return i == "baba"
	})
	if !condition {
		t.FailNow()
	}
}

func TestStream_Every(t *testing.T) {
	var arr = Stream[string]{"aaaa", "bbbb", "cc", "dd", "eeee", "ffff"}
	var condition = arr.Every(func(i string, _ int) bool {
		return len(i) > 0
	})
	if !condition {
		t.FailNow()
	}
	condition = arr.Every(func(i string, _ int) bool {
		return len(i) == 4
	})
	if condition {
		t.FailNow()
	}
	condition = arr.Every(func(i string, _ int) bool {
		return i == "baba"
	})
	if condition {
		t.FailNow()
	}
}

func TestStream_Reduce(t *testing.T) {
	var ints = Stream[int]{1, 2, 3, 4, 5, 6}
	var expected = 1 + 2 + 3 + 4 + 5 + 6
	var result = ints.Reduce(func(sum, item int, _ int) int {
		return sum + item
	})
	if expected != result {
		t.Fatalf("expected: %v, received: %v", expected, result)
	}
}

func TestStream_Concat(t *testing.T) {
	var arr Stream[string] = []string{"aaaa", "bbbb", "cc", "dd", "eeee", "ffff"}
	arr = arr.Concat("a", "b", "c")
	if !reflect.DeepEqual(Stream[string]{"a", "b", "c"}, arr[6:]) {
		t.Fatalf("expecting: [a b c] at the end of the array, received: %v", arr)
	}
}

func TestStream_Prepend(t *testing.T) {
	var a = Stream[string]{"aaaa", "bbbb", "cc", "dd", "eeee", "ffff"}
	a.AddAt(0, "eita")
	fmt.Println(a)
	var arr Stream[string] = []string{"aaaa", "bbbb", "cc", "dd", "eeee", "ffff"}
	arr = arr.Prepend("a", "b", "c")
	if !reflect.DeepEqual(Stream[string]{"a", "b", "c"}, arr[:3]) {
		t.Fatalf("expecting: [a b c] at the beggining of the array, received: %v", arr)
	}
}

func TestStream_AddAt(t *testing.T) {
	var arr Stream[string] = []string{"aaaa", "bbbb", "cc", "dd", "eeee", "ffff"}
	arr = arr.AddAt(3, "a", "b", "c")
	if !reflect.DeepEqual(Stream[string]{"a", "b", "c"}, arr[3:6]) {
		t.Fatalf("expecting: [a b c] at the index 3, 4 and 5, received: %v", arr)
	}
}

func TestStream_Set(t *testing.T) {
	var arr Stream[string] = []string{"aaaa", "bbbb", "cc", "dd", "eeee", "ffff"}
	arr = arr.With(3, "a")
	if arr[3] != "a" {
		t.Fatalf("expecting: a at the index 3 of the array, received: %v", arr[3])
	}
}

func TestStream_Map(t *testing.T) {
	var arr = Stream[int]{1, 2, 3, 4, 5}
	var expected = Stream[int]{2, 4, 6, 8, 10}
	var result = arr.
		Map(func(i, _ int) int {
			return i * 2
		})
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("expected: %q, received: %q", expected, result)
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

func TestStream_Sort(t *testing.T) {
	var ints = Stream[int]{10, 2, 33, 14, 5, 67}
	var expected = Stream[int]{2, 5, 10, 14, 33, 67}
	var result = ints.Sort(func(i, j int) bool {
		return i < j
	})
	if !reflect.DeepEqual(expected, result) {
		t.Fatalf("expected: %v, received: %v", expected, result)
	}
}
