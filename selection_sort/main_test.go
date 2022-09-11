package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestSelectionSort(t *testing.T) {
	for _, tc := range []struct {
		input []int
		want  []int
	}{
		{[]int{0}, []int{0}},
		{[]int{1}, []int{1}},
		{[]int{2}, []int{2}},
		{[]int{1, 2}, []int{1, 2}},
		{[]int{2, 1}, []int{1, 2}},
		{[]int{1, 1}, []int{1, 1}},
		{[]int{2, 2}, []int{2, 2}},
		{[]int{1, 2, 3}, []int{1, 2, 3}},
		{[]int{1, 2, 4}, []int{1, 2, 4}},
		{[]int{3, 2, 1}, []int{1, 2, 3}},
		{[]int{3, 1, 2}, []int{1, 2, 3}},
		{[]int{2, 2, 3, 2}, []int{2, 2, 2, 3}},
		{[]int{6, 2, 5, 9, 1, 6, 3, 6, 6, 5}, []int{1, 2, 3, 5, 5, 6, 6, 6, 6, 9}},
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{10, 9, 8, 7, 6, 5, 4, 3, 2, 1, 0}, []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10}},
		{[]int{-1}, []int{-1}},
		{[]int{100}, []int{100}},
		{[]int{236, 326, 236}, []int{236, 236, 326}},
		{[]int{239874266}, []int{239874266}},
		{[]int{1019282772338393884}, []int{1019282772338393884}},
	} {
		t.Run(fmt.Sprint(tc.input), func(t *testing.T) {
			got := make([]int, len(tc.input))
			copy(got, tc.input)
			SelectionSort(got)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("SelectionSort(%v) = %v, want = %v", tc.input, got, tc.want)
			}
		})
	}
}

var input []int

func init() {
	rand.Seed(time.Now().UnixNano())
	input = make([]int, 10_000)
	for i := range input {
		input[i] = rand.Intn(1_000)
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		s := make([]int, len(input))
		copy(s, input)
		b.StartTimer()
		SelectionSort(s)
	}
}
