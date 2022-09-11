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

func PosBisect(s []int, x int) int {
	start, end := 0, len(s)-1
	var middle int
	for start <= end {
		middle = (start + end) / 2
		if s[middle] == x {
			return middle
		} else if s[middle] > x {
			end = middle - 1
		} else {
			start = middle + 1
		}
	}
	if s[middle] > x {
		return start
	}
	return end + 1
}

func insertionSort(s []int) {
	// +++ ---
	//     i
	// jjj
	//         3
	// 1 2 3 4 5
	//     j   i

	for i := 1; i < len(s); i++ {
		pos := PosBisect(s[:i], s[i])
		/*
			tmp := s[i]
			s = append(s[:i], s[i+1:]...)
			s = append(s[:pos], append([]int{tmp}, s[pos:]...)...)
		*/
		for j := i; j > pos; j-- {
			s[j], s[j-1] = s[j-1], s[j]
		}
	}
}

func main() {
	a := []int{5, 8, 1, 3, 9, 3, 2, 1}
	insertionSort(a)
	fmt.Println(a, PosBisect(a, 4))
}
