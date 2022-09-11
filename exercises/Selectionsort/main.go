package main

import (
	"fmt"
	"math/rand"
	"time"
)

func sort(s []int) {
	for i := range s {
		tmp := i
		for j := i; j < len(s); j++ {
			if s[tmp] > s[j] {
				tmp = j
			}
		}
		s[tmp], s[i] = s[i], s[tmp]
	}
}
func main() {
	rand.Seed(time.Now().UnixNano())
	s := make([]int, rand.Intn(2)+6)
	for i := range s {
		s[i] = rand.Intn(11)
	}
	fmt.Println(s)
	sort(s)
	fmt.Println(s)
}
