package main

import (
	"reflect"
	"testing"
)

func TestInsertion(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{[]int{3, 4, 5, 2, 1}, []int{1, 2, 3, 4, 5}},
		{nil, nil},
		{[]int{1}, []int{1}},
		{[]int{-1, -2, -10, -3}, []int{-10, -3, -2, -1}},
		{[]int{3, 4, 5, 2, 1, 7, 8, -1, -3}, []int{-3, -1, 1, 2, 3, 4, 5, 7, 8}},
	} {
		t.Run("", func(t *testing.T) {
			if insertion(tc.s); !reflect.DeepEqual(tc.s, tc.want) {
				t.Errorf("got = %v, want = %v", tc.s, tc.want)
			}
		})
	}
}

func TestSelection(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{[]int{3, 4, 5, 2, 1}, []int{1, 2, 3, 4, 5}},
		{nil, nil},
		{[]int{1}, []int{1}},
		{[]int{-1, -2, -10, -3}, []int{-10, -3, -2, -1}},
		{[]int{3, 4, 5, 2, 1, 7, 8, -1, -3}, []int{-3, -1, 1, 2, 3, 4, 5, 7, 8}},
	} {
		t.Run("", func(t *testing.T) {
			if selection(tc.s); !reflect.DeepEqual(tc.s, tc.want) {
				t.Errorf("got = %v, want = %v", tc.s, tc.want)
			}
		})
	}
}

func TestReverseIns(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{[]int{3, 4, 5, 2, 1}, []int{1, 2, 3, 4, 5}},
		{nil, nil},
		{[]int{1}, []int{1}},
		{[]int{-1, -2, -10, -3}, []int{-10, -3, -2, -1}},
		{[]int{3, 4, 5, 2, 1, 7, 8, -1, -3}, []int{-3, -1, 1, 2, 3, 4, 5, 7, 8}},
	} {
		t.Run("", func(t *testing.T) {
			if reverseIns(tc.s); !reflect.DeepEqual(tc.s, tc.want) {
				t.Errorf("got = %v, want = %v", tc.s, tc.want)
			}
		})
	}
}

var input = []int{3, 4, 5, 2, 1, 7, 8, -1, -3, 7, 7, -10, 27, 12, 123, 123, 23, 123, 21, 3, -2341}

func BenchmarkInsertion(b *testing.B) {
	s := make([]int, len(input))
	for i := 0; i < b.N; i++ {
		copy(s, input)
		insertion(s)
	}
}

func BenchmarkSelection(b *testing.B) {
	s := make([]int, len(input))
	for i := 0; i < b.N; i++ {
		copy(s, input)
		selection(s)
	}

}

func BenchmarkReverseIns(b *testing.B) {
	s := make([]int, len(input))
	for i := 0; i < b.N; i++ {
		copy(s, input)
		reverseIns(s)
	}
}
