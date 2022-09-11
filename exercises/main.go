package main

import (
	"fmt"
	"math/rand"
	"time"
)

func Bsearch(x []int, b int) int {
	var mid int
	low, high := 0, len(x)-1
	for low <= high {
		mid = (low + high) / 2
		if x[mid] == b {
			return mid
		}
		if b > x[mid] {
			low = mid + 1
		}
		if b < x[mid] {
			high = mid - 1
		}
	}
	return mid
}

func Binsertion(x []int) {
	for i := 1; i < len(x); i++ {

		j := Bsearch(x[0:i], x[i])

		if j < i {
			temp := x[i]
			for k := i; k > j; k-- {
				x[k] = x[k-1]
			}
			x[j] = temp
		}

	}
}
func Sselect(x []int) {
	for i := 0; i < len(x); i++ {
		small := i
		for j := i; j < len(x); j++ {
			if x[j] < x[small] {
				small = j

			}
			x[j], x[small] = x[small], x[j]
		}

	}
}

func main() {
	rand.Seed(time.Now().UnixNano())
	x := make([]int, rand.Intn(10)+11)
	for i := range x {
		x[i] = rand.Intn(201) - 100
	}

	x2 := make([]int, rand.Intn(10)+11)
	for i := range x2 {
		x2[i] = rand.Intn(201) - 100
	}
	fmt.Println(x)
	Binsertion(x)
	fmt.Println(x)
	fmt.Println(x2)
	Sselect(x2)
	fmt.Println(x2)
}
