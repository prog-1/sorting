package main

import (
	"reflect"
	"testing"
)

func TestSelectionSort(t *testing.T) {
	for _, tc := range []struct {
		name  string
		input []int
		want  []int
	}{
		{"various elements", []int{43, 17, 45, 75, 2}, []int{2, 17, 43, 45, 75}},
		{"nil", nil, nil},
		{"empty", []int{}, []int{}},
		{"one", []int{1}, []int{1}},
		{"same", []int{2, 2}, []int{2, 2}},
		{"negative", []int{-23, -32, -2, -10}, []int{-32, -23, -10, -2}},
	} {
		t.Run(tc.name, func(t *testing.T) {
			var got []int
			if tc.input != nil { // Make a copy to avoid side-effects.
				got = make([]int, len(tc.input))
				copy(got, tc.input)
			}
			SelectionSort(got)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("got = %v, want = %v", got, tc.want)
			}
		})
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	for i := 0; i < b.N; i++ {
		SelectionSort([]int{43, 18, 92, 12, 54})
	}
}
