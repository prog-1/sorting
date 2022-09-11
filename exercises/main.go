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

func PosBesect(s []int, x int) int {
	start, end := 0, len(s)-1
	var middle int
	for start-end != 1 {
		middle = (start + end) / 2
		if s[middle] == x {
			return middle
		} else if s[middle] > x {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}
	if middle > x {
		return start - 1
	}
	return end + 1
}

func main() {
	a := []int{5, 8, 1, 3, 9, 3, 2, 1}
	selectionSort(a)
	fmt.Println(a, PosBesect(a, 4))
}
