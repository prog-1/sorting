package main

import (
	"fmt"
	"math/rand"
	"time"
)

//insertion sort (bisect sort)
func Sort(s []int) {
	for i := 1; i < len(s); i++ {
		x := s[i]
		pos := 0
		last := i - 1
		for pos <= last {
			mid := pos + (last-pos)/2
			if x < s[mid] {
				last = mid - 1
			} else {
				pos = mid + 1
			}
		}

		for j := i - 1; j >= pos; j-- {
			s[j+1] = s[j]
		}

		s[pos] = x
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	s := make([]int, rand.Intn(10)+11)
	for i := range s {
		s[i] = rand.Intn(201) - 100
	}
	fmt.Println(s)
	Sort(s)
	fmt.Println(s)
}
