package main

import (
	"fmt"
	"math/rand"
	"reflect"
	"testing"
	"time"
)

func TestInsertionSort(t *testing.T) {
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
			InsertionSort(got)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("InsertionSort(%v) = %v, want = %v", tc.input, got, tc.want)
			}
		})
	}
}

func TestPosBisect(t *testing.T) {
	for _, tc := range []struct {
		s    []int
		x    int
		want int
	}{
		{[]int{0}, 0, 0},
		{[]int{0}, 1, 1},
		{[]int{0}, 2, 1},
		{[]int{1}, 1, 0},
		{[]int{1}, 2, 1},
		{[]int{1}, -1, 0},
		{[]int{2}, 2, 0},
		{[]int{2}, 1, 0},
		{[]int{2}, 3, 1},
		{[]int{1, 2}, 1, 0},
		{[]int{1, 2}, 2, 1},
		{[]int{1, 2}, 3, 2},
		{[]int{1, 2}, 0, 0},
		{[]int{1, 3}, 3, 1},
		{[]int{2, 3}, 2, 0},
		{[]int{2, 3}, 3, 1},
		{[]int{2, 3}, 1, 0},
		{[]int{2, 3}, 4, 2},
		{[]int{2, 4}, 4, 1},
		{[]int{2, 4}, 3, 1},
		{[]int{0, 1, 2, 3, 4}, 3, 3},
		{[]int{0, 1, 2, 3, 4}, 4, 4},
		{[]int{1, 2, 3, 4, 5}, 4, 3},
		{[]int{1, 2, 3, 8, 9}, 3, 2},
		{[]int{-1}, -1, 0},
		{[]int{-1}, 0, 1},
		{[]int{-1}, -2, 0},
		{[]int{-1}, 1, 1},
		{[]int{-2}, -2, 0},
		{[]int{-2}, -1, 1},
		{[]int{-2}, -3, 0},
		{[]int{-2}, 2, 1},
		{[]int{-2, -1, 0, 2}, -1, 1},
		{[]int{-2, -1, 0, 2}, 2, 3},
		{[]int{-2, -1, 0, 2}, 1, 3},
		{[]int{-3, -2, 0, 3, 4, 8, 11}, -3, 0},
		{[]int{-3, -2, 0, 3, 4, 8, 11}, -2, 1},
		{[]int{-3, -2, 0, 3, 4, 8, 11}, 3, 3},
		{[]int{-3, -2, 0, 3, 4, 8, 11}, 8, 5},
	} {
		t.Run(fmt.Sprint(tc.s)+","+fmt.Sprint(tc.x), func(t *testing.T) {
			if got := PosBisect(tc.s, tc.x); !reflect.DeepEqual(got, tc.want) {
				t.Errorf("PosBisect(%v, %v) = %v, want = %v", tc.s, tc.x, got, tc.want)
			}
		})
	}
}

func TestInsertionSortWithPosBisect(t *testing.T) {
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
			InsertionSortWithPosBisect(got)
			if !reflect.DeepEqual(got, tc.want) {
				t.Errorf("InsertionSortWithBisect(%v) = %v, want = %v", tc.input, got, tc.want)
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

func BenchmarkInsertionSort(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		s := make([]int, len(input))
		copy(s, input)
		b.StartTimer()
		InsertionSort(s)
	}
}

func BenchmarkPosBisect(b *testing.B) {
	rand.Seed(time.Now().UnixNano())
	input = make([]int, 1_000)
	for i := range input {
		input[i] = rand.Intn(100)
	}
	x := rand.Intn(100)
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		s := make([]int, len(input))
		copy(s, input)
		b.StartTimer()
		PosBisect(s, x)
	}
}

func BenchmarkInsertionSortWithPosBisect(b *testing.B) {
	for n := 0; n < b.N; n++ {
		b.StopTimer()
		s := make([]int, len(input))
		copy(s, input)
		b.StartTimer()
		InsertionSortWithPosBisect(s)
	}
}
