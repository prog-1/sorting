package main

import "fmt"

func selectionSort(s []int) {
	for i := 0; i < len(s)-1; i++ {
		minIndex := i
		for j := range s[i:] {
			if s[minIndex] > s[j+i] {
				minIndex = j + i
			}
		}
		s[i], s[minIndex] = s[minIndex], s[i]
	}
}

func main() {
	a := []int{5, 8, 1, 3, 9, 3, 2, 1}
	selectionSort(a)
	fmt.Println(a)
}
