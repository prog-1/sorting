package main

import (
	"reflect"
	"testing"
)

func TestInsertionSort(t *testing.T) {
	for _, tc := range []struct {
		name string
		got  []int
		want []int
	}{
		{"nil", nil, nil},
		{"empty", []int{}, []int{}},
		{"one number", []int{1}, []int{1}},
		{"twosame numbers", []int{1, 1}, []int{1, 1}},
		{"with negative numbers", []int{1, -1}, []int{-1, 1}},
		{"different numbers", []int{1, 8, 4, 0, -7, -23, 1}, []int{-23, -7, 0, 1, 1, 4, 8}},
	} {
		t.Run("", func(t *testing.T) {
			if insertionsort(tc.got); !reflect.DeepEqual(tc.got, tc.want) {
				t.Errorf("got = %v, want = %v", tc.got, tc.want)
			}
		})
	}
}

func TestSelectionSort(t *testing.T) {
	for _, tc := range []struct {
		name string
		got  []int
		want []int
	}{
		{"nil", nil, nil},
		{"empty", []int{}, []int{}},
		{"one number", []int{1}, []int{1}},
		{"twosame numbers", []int{1, 1}, []int{1, 1}},
		{"with negative numbers", []int{1, -1}, []int{-1, 1}},
		{"different numbers", []int{1, 8, 4, 0, -7, -23, 1}, []int{-23, -7, 0, 1, 1, 4, 8}},
	} {
		t.Run("", func(t *testing.T) {
			if selectionsort(tc.got); !reflect.DeepEqual(tc.got, tc.want) {
				t.Errorf("got = %v, want = %v", tc.got, tc.want)
			}
		})
	}
}

func BenchmarkInsertionSort(b *testing.B) {
	for _, tc := range []struct {
		name string
		got  []int
		want []int
	}{
		{"nil", nil, nil},
		{"empty", []int{}, []int{}},
		{"one number", []int{1}, []int{1}},
		{"twosame numbers", []int{1, 1}, []int{1, 1}},
		{"with negative numbers", []int{1, -1}, []int{-1, 1}},
		{"different numbers", []int{1, 8, 4, 0, -7, -23, 1}, []int{-23, -7, 0, 1, 1, 4, 8}},
	} {
		for i := 0; i < b.N; i++ {
			insertionsort(tc.got)
		}
	}
}

func BenchmarkSelectionSort(b *testing.B) {
	for _, tc := range []struct {
		name string
		got  []int
		want []int
	}{
		{"nil", nil, nil},
		{"empty", []int{}, []int{}},
		{"one number", []int{1}, []int{1}},
		{"twosame numbers", []int{1, 1}, []int{1, 1}},
		{"with negative numbers", []int{1, -1}, []int{-1, 1}},
		{"different numbers", []int{1, 8, 4, 0, -7, -23, 1}, []int{-23, -7, 0, 1, 1, 4, 8}},
	} {
		for i := 0; i < b.N; i++ {
			selectionsort(tc.got)
		}
	}
}
