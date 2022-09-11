package main

import (
	"reflect"
	"testing"
)

func TestInsertionSortBisectSearch(t *testing.T) {
	for _, tc := range []struct {
		name string
		s    []int
		want []int
	}{
		{"case-1", nil, nil},
		{"case-2", []int{0}, []int{0}},
		{"case-3", []int{1, 2, 3}, []int{1, 2, 3}},
		{"case-4", []int{3, 2, 1}, []int{1, 2, 3}},
		{"case-5", []int{2, 1, 3}, []int{1, 2, 3}},
		{"case-6", []int{2, 1, 4, 3}, []int{1, 2, 3, 4}},
		{"case-7", []int{-1, 3, 2, 4, 2, -1}, []int{-1, -1, 2, 2, 3, 4}},
		{"case-8", []int{-4, 2, 1, -3, 2, 1, 3}, []int{-4, -3, 1, 1, 2, 2, 3}},
		{"case-9", []int{4, 2, 1, 3, 2, -1, 4}, []int{-1, 1, 2, 2, 3, 4, 4}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if insertionSortBisectSearch(tc.s); !reflect.DeepEqual(tc.s, tc.want) {
				t.Errorf("got = %v, want = %v", tc.s, tc.want)
			}
		})
	}
}

func BenchmarkInsertionSortBisectSearch(b *testing.B) {
	for _, bm := range []struct {
		name string
		s    []int
		want []int
	}{
		{"case-1", nil, nil},
		{"case-2", []int{0}, []int{0}},
		{"case-3", []int{1, 2, 3}, []int{1, 2, 3}},
		{"case-4", []int{3, 2, 1}, []int{1, 2, 3}},
		{"case-5", []int{2, 1, 3}, []int{1, 2, 3}},
		{"case-6", []int{2, 1, 4, 3}, []int{1, 2, 3, 4}},
		{"case-7", []int{-1, 3, 2, 4, 2, -1}, []int{-1, -1, 2, 2, 3, 4}},
		{"case-8", []int{-4, 2, 1, -3, 2, 1, 3}, []int{-4, -3, 1, 1, 2, 2, 3}},
		{"case-9", []int{4, 2, 1, 3, 2, -1, 4}, []int{-1, 1, 2, 2, 3, 4, 4}},
	} {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				insertionSortBisectSearch(bm.s)
			}
		})
	}
}

func TestSelectionSort(t *testing.T) {
	for _, tc := range []struct {
		name string
		s    []int
		want []int
	}{
		{"case-1", nil, nil},
		{"case-2", []int{0}, []int{0}},
		{"case-3", []int{1, 2, 3}, []int{1, 2, 3}},
		{"case-4", []int{3, 2, 1}, []int{1, 2, 3}},
		{"case-5", []int{2, 1, 3}, []int{1, 2, 3}},
		{"case-6", []int{2, 1, 4, 3}, []int{1, 2, 3, 4}},
		{"case-7", []int{-1, 3, 2, 4, 2, -1}, []int{-1, -1, 2, 2, 3, 4}},
		{"case-8", []int{-4, 2, 1, -3, 2, 1, 3}, []int{-4, -3, 1, 1, 2, 2, 3}},
		{"case-9", []int{4, 2, 1, 3, 2, -1, 4}, []int{-1, 1, 2, 2, 3, 4, 4}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if selectionSort(tc.s); !reflect.DeepEqual(tc.s, tc.want) {
				t.Errorf("got = %v, want = %v", tc.s, tc.want)
			}
		})
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	for _, bm := range []struct {
		name string
		s    []int
		want []int
	}{
		{"case-1", nil, nil},
		{"case-2", []int{0}, []int{0}},
		{"case-3", []int{1, 2, 3}, []int{1, 2, 3}},
		{"case-4", []int{3, 2, 1}, []int{1, 2, 3}},
		{"case-5", []int{2, 1, 3}, []int{1, 2, 3}},
		{"case-6", []int{2, 1, 4, 3}, []int{1, 2, 3, 4}},
		{"case-7", []int{-1, 3, 2, 4, 2, -1}, []int{-1, -1, 2, 2, 3, 4}},
		{"case-8", []int{-4, 2, 1, -3, 2, 1, 3}, []int{-4, -3, 1, 1, 2, 2, 3}},
		{"case-9", []int{4, 2, 1, 3, 2, -1, 4}, []int{-1, 1, 2, 2, 3, 4, 4}},
	} {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				selectionSort(bm.s)
			}
		})
	}
}

func TestInsertionSortLinearSearch(t *testing.T) {
	for _, tc := range []struct {
		name string
		s    []int
		want []int
	}{
		{"case-1", nil, nil},
		{"case-2", []int{0}, []int{0}},
		{"case-3", []int{1, 2, 3}, []int{1, 2, 3}},
		{"case-4", []int{3, 2, 1}, []int{1, 2, 3}},
		{"case-5", []int{2, 1, 3}, []int{1, 2, 3}},
		{"case-6", []int{2, 1, 4, 3}, []int{1, 2, 3, 4}},
		{"case-7", []int{-1, 3, 2, 4, 2, -1}, []int{-1, -1, 2, 2, 3, 4}},
		{"case-8", []int{-4, 2, 1, -3, 2, 1, 3}, []int{-4, -3, 1, 1, 2, 2, 3}},
		{"case-9", []int{4, 2, 1, 3, 2, -1, 4}, []int{-1, 1, 2, 2, 3, 4, 4}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			if insertionSortLinearSearch(tc.s); !reflect.DeepEqual(tc.s, tc.want) {
				t.Errorf("got = %v, want = %v", tc.s, tc.want)
			}
		})
	}
}

func BenchmarkInsertionSortLinearSearch(b *testing.B) {
	for _, bm := range []struct {
		name string
		s    []int
		want []int
	}{
		{"case-1", nil, nil},
		{"case-2", []int{0}, []int{0}},
		{"case-3", []int{1, 2, 3}, []int{1, 2, 3}},
		{"case-4", []int{3, 2, 1}, []int{1, 2, 3}},
		{"case-5", []int{2, 1, 3}, []int{1, 2, 3}},
		{"case-6", []int{2, 1, 4, 3}, []int{1, 2, 3, 4}},
		{"case-7", []int{-1, 3, 2, 4, 2, -1}, []int{-1, -1, 2, 2, 3, 4}},
		{"case-8", []int{-4, 2, 1, -3, 2, 1, 3}, []int{-4, -3, 1, 1, 2, 2, 3}},
		{"case-9", []int{4, 2, 1, 3, 2, -1, 4}, []int{-1, 1, 2, 2, 3, 4, 4}},
	} {
		b.Run(bm.name, func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				insertionSortLinearSearch(bm.s)
			}
		})
	}
}
