package main

import (
	"reflect"
	"testing"
)

func TestBinsertion(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{},
		{nil, nil},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 1}, []int{1, 1}},
		{[]int{-1, 1}, []int{-1, 1}},
		{[]int{1, -1}, []int{-1, 1}},
		{[]int{3, 2, 1, 0, -1, -2, -3}, []int{-3, -2, -1, 0, 1, 2, 3}},
	} {
		t.Run("", func(t *testing.T) {
			if Binsertion(tc.s); !reflect.DeepEqual(tc.s, tc.want) {
				t.Errorf("got = %v, want = %v", tc.s, tc.want)
			}
		})
	}
}
func TestSselect(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{},
		{nil, nil},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 1}, []int{1, 1}},
		{[]int{-1, 1}, []int{1, -1}},
		{[]int{1, -1}, []int{1, -1}},
		{[]int{-3, -2, -1, 0, 1, 2, 3}, []int{3, 2, 1, 0, -1, -2, -3}},
	} {
		t.Run("", func(t *testing.T) {
			if Sselect(tc.s); !reflect.DeepEqual(tc.s, tc.want) {
				t.Errorf("got = %v, want = %v", tc.s, tc.want)
			}
		})
	}
}
func BenchmarkBinsertion(b *testing.B) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{},
		{nil, nil},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 1}, []int{1, 1}},
		{[]int{-1, 1}, []int{-1, 1}},
		{[]int{1, -1}, []int{-1, 1}},
		{[]int{3, 2, 1, 0, -1, -2, -3}, []int{-3, -2, -1, 0, 1, 2, 3}},
	} {
		for i := 0; i < b.N; i++ {
			Binsertion(tc.s)
		}
	}

}
func BenchmarkSselect(b *testing.B) {
	for _, tc := range []struct {
		s    []int
		want []int
	}{
		{},
		{nil, nil},
		{[]int{}, []int{}},
		{[]int{1}, []int{1}},
		{[]int{1, 1}, []int{1, 1}},
		{[]int{-1, 1}, []int{1, -1}},
		{[]int{1, -1}, []int{1, -1}},
		{[]int{-3, -2, -1, 0, 1, 2, 3}, []int{3, 2, 1, 0, -1, -2, -3}},
	} {
		for i := 0; i < b.N; i++ {
			Sselect(tc.s)
		}
	}

}
