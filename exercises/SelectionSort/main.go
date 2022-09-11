package main

import "fmt"

func SelectionSort(s []int) {
	for i := 0; i < len(s); i++ {
		// Finding the smallest number in array:
		smallestIndex := i
		for j := i; j < len(s); j++ {
			if s[j] < s[smallestIndex] {
				smallestIndex = j
			}
		}

		// Swapping smallest number with the first unsorted element:
		s[i], s[smallestIndex] = s[smallestIndex], s[i]

	}
}

func main() {
	s := []int{18, 12, 5, 84, 1}

	SelectionSort(s)

	fmt.Println(s)
}
