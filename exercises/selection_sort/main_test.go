package main

import (
	"testing"
)

func TestSort(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{},
		{[]int{}, []int{}},
		{[]int{7}, []int{7}},
		{[]int{10, 10}, []int{10, 10}},
		{[]int{-9, 1}, []int{-9, 1}},
		{[]int{20, -1}, []int{-1, 20}},
		{[]int{-44, -35}, []int{-44, -35}},
		{[]int{20, -1, 19}, []int{-1, 19, 20}},
		{[]int{1, 5, 8, 9, 10}, []int{1, 5, 8, 9, 10}},
		{[]int{3, 0, 2}, []int{0, 2, 3}},
		{[]int{4, 7, 11, -3, -4, 0}, []int{-4, -3, 0, 4, 7, 11}},
		{[]int{-40, 57, -5, 82, 55, -74, -93, -19, 32, -100, -37}, []int{-100, -93, -74, -40, -37, -19, -5, 32, 55, 57, 82}},
	} {
		got := make([]int, len(tc.s))
		copy(got, tc.s)
		Sort(got)
		if !equal(got, tc.want) {
			t.Errorf("Sort1(%v) = %v, want = %v", tc.s, got, tc.want)
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

var (
	s = []int{11, 10, -20, 100000000, -1, 0, 0, 7, 35, 88, 10, 200, 99, 2, 3, 2, 8, 29, -68, 92, 93, 77, -41, -45, -62, 82, 40, 72, 71, -82, 56, -92, -22, 41, 60, 93, 47, -14, -88, 37, 76, -72, -75, -65, -83, 52, 14, -35}
)

func BenchmarkSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		Sort(s)
	}
}
