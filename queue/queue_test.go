package queue

import (
	"testing"
)

func less(x, y int64) bool {
	return x < y
}

func eq(a, b []int64) bool {
	for idx, av := range a {
		if av != b[idx] { return false }
	}
	return true
}


func TestSimple(t *testing.T) {
	q := make([]int64, 0)
	q = Qadd(q, 5, less)
	q = Qadd(q, 3, less)
	q = Qadd(q, 1, less)
	q = Qadd(q, 2, less)
	q = Qadd(q, 4, less)

	a := make([]int64, 0, len(q))

	for {
		var item *int64
		var ok bool
		item, q, ok = Qpop(q, less)
		if !ok { break }
		a = append(a, *item)
	}

	if !eq(a, []int64{1,2,3,4,5}) {
		t.Fatalf("Got %v", a)
	}
}
