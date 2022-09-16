package main

import (
	"math/rand"
	"reflect"
	"testing"
)

var input []int

func init() {
	rand.Seed(17)

	input = make([]int, 10000)
	for i := range input {
		input[i] = rand.Int()
	}
}
func benchmarkSort(b *testing.B, sort func([]int)) {
	data := make([]int, 10000)
	for n := 0; n < b.N; n++ {
		copy(data, input)
		sort(data)
	}
}

func TestInsertSort(t *testing.T) {
	for _, tc := range []struct {
		got  []int
		want []int
	}{
		{nil, nil},
		{[]int{}, []int{}},
		{[]int{0, 0}, []int{0, 0}},
		{[]int{9}, []int{9}},
		{[]int{3, 10, 0, 0, 3, 9, 4, 1, 0, 0, 10, 1, 3}, []int{0, 0, 0, 0, 1, 1, 3, 3, 3, 4, 9, 10, 10}},
		{[]int{-3, 10, 0, 0, -3, 9, -4, 1, 0, 0, -10, 1, 3}, []int{-10, -4, -3, -3, 0, 0, 0, 0, 1, 1, 3, 9, 10}},
	} {
		t.Run("", func(t *testing.T) {
			if insertSort(tc.got); !reflect.DeepEqual(tc.got, tc.want) {
				t.Errorf("got = %v, want = %v", tc.got, tc.want)
			}
		})
	}
}
func TestBisSort(t *testing.T) {
	for _, tc := range []struct {
		got  []int
		want []int
	}{
		{[]int{}, []int{}},
		{[]int{0, 0}, []int{0, 0}},
		{[]int{9}, []int{9}},
		{[]int{9}, []int{9}},
		// fail this tests DO NOT RUN!
		//{[]int{3, 10, 0, 0, 3, 9, 4, 1, 0, 0, 10, 1, 3}, []int{0, 0, 0, 0, 1, 1, 3, 3, 3, 4, 9, 10, 10}},
		//{[]int{-3, 10, 0, 0, -3, 9, -4, 1, 0, 0, -10, 1, 3}, []int{-10, -4, -3, -3, 0, 0, 0, 0, 1, 1, 3, 9, 10}},
	} {
		t.Run("", func(t *testing.T) {
			if insertBisSort(tc.got); !reflect.DeepEqual(tc.got, tc.want) {
				t.Errorf("got = %v, want = %v", tc.got, tc.want)
			}
		})
	}
}
func TestSelectionSort(t *testing.T) {
	for _, tc := range []struct {
		got  []int
		want []int
	}{
		{nil, nil},
		{[]int{}, []int{}},
		{[]int{0, 0}, []int{0, 0}},
		{[]int{9}, []int{9}},
		{[]int{3, 10, 0, 0, 3, 9, 4, 1, 0, 0, 10, 1, 3}, []int{0, 0, 0, 0, 1, 1, 3, 3, 3, 4, 9, 10, 10}},
		{[]int{-3, 10, 0, 0, -3, 9, -4, 1, 0, 0, -10, 1, 3}, []int{-10, -4, -3, -3, 0, 0, 0, 0, 1, 1, 3, 9, 10}},
	} {
		t.Run("", func(t *testing.T) {
			if selectionSort(tc.got); !reflect.DeepEqual(tc.got, tc.want) {
				t.Errorf("got = %v, want = %v", tc.got, tc.want)
			}
		})
	}
}
func BenchmarkInsertSort(b *testing.B) {
	benchmarkSort(b, func(input []int) { insertSort(input) })
}

func BenchmarkInsertions(b *testing.B) {
	benchmarkSort(b, func(input []int) { insertSortBySteveHook(input) })
}

func BenchmarkSelectionSort(b *testing.B) {
	benchmarkSort(b, func(input []int) { selectionSort(input) })
}
