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

func BenchmarkInsertion(b *testing.B) {
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
		for i := 0; i < b.N; i++ {
			insertion(tc.s)
		}
	}

}

func BenchmarkSelection(b *testing.B) {
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
		for i := 0; i < b.N; i++ {
			selection(tc.s)
		}
	}

}

func BenchmarkReverseIns(b *testing.B) {
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
		for i := 0; i < b.N; i++ {
			reverseIns(tc.s)
		}
	}

}
