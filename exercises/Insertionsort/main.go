package main

import (
	"fmt"
	"math/rand"
	"time"
)

func bsearch(s []int, num int) (out int) {
	max := len(s) - 1
	min := 0
	var mid int
	for min <= max {
		mid = (min + max) / 2
		if num == s[mid] {
			return mid
		}
		if s[mid] < num {
			min = mid + 1
		} else {
			max = mid - 1
		}
	}
	if s[mid] > num {
		return min
	} else {
		return max + 1
	}
}

func sort(s []int) {
	for i := 1; i < len(s); i++ {
		pos := bsearch(s[:i], s[i])
		for j := i; j > pos; j-- {
			s[j], s[j-1] = s[j-1], s[j]
		}
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
