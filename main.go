package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bubblesort(s []int) {
	for l := len(s); l > 1; l-- {
		for i := 1; i < l; i++ {
			if s[i] < s[i-1] {
				s[i], s[i-1] = s[i-1], s[i]
			}
		}
	}

}

func insertionSort(s []int) {
	// +++ ---
	//     i
	// jjj
	//         3
	// 1 2 3 4 5
	//     j   i

	for i := 1; i < len(s); i++ {
		j := 0
		for ; s[j] < s[i]; j++ {

		}
		if j < i {
			temp := s[i]
			for k := i; k > j; k-- {
				s[k] = s[k-1]
			}
			s[j] = temp
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	s := make([]int, rand.Intn(11)+10)
	for i := range s {
		s[i] = rand.Intn(201) - 100
	}
	fmt.Println(s)
	bubblesort(s)
	fmt.Println(s)
}
