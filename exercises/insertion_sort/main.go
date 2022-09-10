package main

import (
	"fmt"
	"math/rand"
	"time"
)

//insertion sort
func Sort(s []int) {
	for i := 1; i < len(s); i++ {
		x := s[i]
		j := 0
		for s[i] > s[j] {
			j++
		}
		for k := i; k > j; k-- {
			s[k] = s[k-1]
		}
		s[j] = x
	}
}

//insertion sort
func Sort2(s []int) {
	for i := 1; i < len(s); i++ {
		x := i
		for j := i; j >= 0; j-- {
			if s[j] > s[x] {
				s[j], s[j+1] = s[j+1], s[j]
			}
			x = j
		}
	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	s := make([]int, rand.Intn(10)+11)
	for i := range s {
		s[i] = rand.Intn(201) - 100
	}
	s2 := make([]int, rand.Intn(10)+11)
	for i := range s2 {
		s2[i] = rand.Intn(201) - 100
	}
	fmt.Println(s)
	Sort(s)
	fmt.Println(s)
	fmt.Println(s2)
	Sort2(s2)
	fmt.Println(s2)
}
