package main

import "testing"

func TestSort(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{},
		{[]int{}, []int{}},
		{[]int{0}, []int{0}},
		{[]int{1}, []int{1}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{3, 9, 1}, []int{1, 3, 9}},
		{[]int{2, 93, 16, 476, 9}, []int{2, 9, 16, 93, 476}},
		{[]int{1, 5, 1, 0, 1, 4}, []int{0, 1, 1, 1, 4, 5}},
		{[]int{-2, -3, -10, 3, 5}, []int{-10, -3, -2, 3, 5}},
		{[]int{-1, 34, 92, 0, 119, -10, -3}, []int{-10, -3, -1, 0, 34, 92, 119}},
		{[]int{-1, -2}, []int{-2, -1}},
		{[]int{2, 9, -1}, []int{-1, 2, 9}},
	} {
		got := make([]int, len(tc.s))
		copy(got, tc.s)
		sort(got)
		if !equal(got, tc.want) {
			t.Errorf("Sort(%v) = %v, want = %v", tc.s, got, tc.want)
		}
	}
}
func equal(a, b []int) bool {
	if len(a) != len(b) {
		return false
	}
	for i := range a {
		if a[i] != b[i] {
			return false
		}
	}
	return true
}

var input = []int{-1, 0, -28, 0, 22, 22, -28, -3, 43, 18, 92, 12, 54}

func BenchmarkSort(b *testing.B) {
	s := make([]int, len(input))
	for n := 0; n < b.N; n++ {
		copy(s, input)
		sort(s)
	}
}
