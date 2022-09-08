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

func insertionsort(s []int) {

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
