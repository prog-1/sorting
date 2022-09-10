package main

import (
	"fmt"
	"math/rand"
	"time"
)

//selection sort
func Sort(s []int) {
	for i := 0; i < len(s); i++ {
		min := i
		for j := i; j < len(s); j++ {
			if s[j] < s[min] {
				min = j
			}
		}
		s[i], s[min] = s[min], s[i]
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
