package main

import (
	"math/rand"
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

func BenchmarkSelectionSortRandom(b *testing.B) {
	s := make([]int, len(input))
	for n := 0; n < b.N; n++ {
		copy(s, input)
		selectionSort(s)
	}
}
