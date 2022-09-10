package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bubblesort(s []int) {
	for i := len(s) - 1; i > 0; i-- {
		for j := 0; j < i; j++ {
			if s[j+1] < s[j] {
				s[j], s[j+1] = s[j+1], s[j]
			}
		}
	}
}

func insertionsort(s []int) {
	for i := 1; i < len(s); i++ {
		for j := i; j > 0; j-- {
			if s[j-1] > s[j] {
				s[j-1], s[j] = s[j], s[j-1]
			}
		}
	}
}

func selectionsort(s []int) {
	for i := 0; i < len(s); i++ {
		var minIdx = i
		for j := i; j < len(s); j++ {
			if s[j] < s[minIdx] {
				minIdx = j
			}
		}
		s[i], s[minIdx] = s[minIdx], s[i]
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	s := make([]int, rand.Intn(11)+10)
	for i := range s {
		s[i] = rand.Intn(201) - 100
	}
	fmt.Println("unsorted slice:  ", s)
	bubblesort(s)
	fmt.Println("bubble sort:  ", s)
	selectionsort(s)
	fmt.Println("selection sorting:  ", s)
	insertionsort(s)
	fmt.Println("insertion sort:  ", s)
}
