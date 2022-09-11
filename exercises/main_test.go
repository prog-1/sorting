package main

import (
	"math/rand"
	"reflect"
	"testing"
	"time"
)

var input []int

func init() {
	rand.Seed(time.Now().UnixNano())
	input = make([]int, 10_000)
	for i := range input {
		input[i] = rand.Intn(1_000)
	}
}

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
	s := make([]int, len(input))
	for n := 0; n < b.N; n++ {
		copy(s, input)
		insertionSortBisectSearch(s)
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
	s := make([]int, len(input))
	for n := 0; n < b.N; n++ {
		copy(s, input)
		selectionSort(s)
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
	s := make([]int, len(input))
	for n := 0; n < b.N; n++ {
		copy(s, input)
		insertionSortLinearSearch(s)
	}
}
