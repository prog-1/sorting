package main

import (
	"math/rand"
	"testing"
	"time"
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

var s []int

func init() {
	rand.Seed(time.Now().UnixNano())
	s = make([]int, 50_000)
	for i := range s {
		s[i] = rand.Intn(1_000)
	}
}
func BenchmarkSort(b *testing.B) {
	s2 := make([]int, len(s))
	for n := 0; n < b.N; n++ {
		copy(s2, s)
		Sort(s)
	}
}
