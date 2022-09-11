package main

import (
	"reflect"
	"testing"
)

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
		for i := 0; i < b.N; i++ {
			insertSort(tc.got)
		}
	}
}

func BenchmarkInsertions(b *testing.B) {
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
		for i := 0; i < b.N; i++ {
			insertSortBySteveHook(tc.got)
		}
	}
}

func BenchmarkSelectionSort(b *testing.B) {
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
		for i := 0; i < b.N; i++ {
			selectionSort(tc.got)
		}
	}
}
