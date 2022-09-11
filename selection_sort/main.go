package main

import "fmt"

func SelectionSort(s []int) {
	for i := 0; i < len(s)-1; i++ {
		min := i                          // assume that min is the first element in the slice
		for j := i + 1; j < len(s); j++ { // test against elements after i to find the smallest
			if s[j] < s[min] {
				min = j // found the new minimum, remember its index
			}
		}
		if min != i {
			s[i], s[min] = s[min], s[i]
		}
	}
}

func main() {
	s := []int{3, 2, 1, 6, 7, 5}
	SelectionSort(s)
	fmt.Println(s)
}
