package main

import "fmt"

// func InsertionSort(s []int) {
//	   for i := 1; i < len(s); i++ {
//		   x := s[i]
//		   j := i - 1
//		   for ; j >= 0 && s[j] > x; j-- {
//			   s[j+1] = s[j]
//		   }
//		   s[j+1] = x
//	   }
// }

func InsertionSort(s []int) {
	for i := 1; i < len(s); i++ {
		for j := i; j > 0 && s[j-1] > s[j]; j-- {
			s[j-1], s[j] = s[j], s[j-1]
		}
	}
}

func PosBisect(s []int, x int) (i int) {
	left, right := 0, len(s)
	for left < right {
		m := (left + right) / 2
		if s[m] < x {
			left = m + 1
		} else { // m >= x
			right = m
		}
	}
	if i == 0 { // if x is not in the slice
		return left
	}
	return i
}

func InsertionSortWithPosBisect(s []int) {
	for i := 1; i < len(s); i++ {
		x := PosBisect(s[:i], s[i]) // x is the position where the number will be inserted
		for j := i; j > x; j-- {
			s[j-1], s[j] = s[j], s[j-1]
		}
	}
}

func main() {
	s := []int{3, 2, 1, 6, 7, 5}
	InsertionSort(s)
	fmt.Println(s)

	ss := []int{3, 2, 1, 6, 7, 5}
	InsertionSortWithPosBisect(ss)
	fmt.Println(ss)
}
