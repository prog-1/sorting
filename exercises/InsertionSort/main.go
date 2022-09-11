package main

import "fmt"

func BinarySearch(s []int, value int) int {

	l, r := 0, len(s)-1

	for {
		m := r / 2

		if r-l <= 1 && s[m] != value {
			// If element is not in the slice, returning its possible position
			if value > s[r] {
				return r + 1
			} else if value <= s[l] {
				return l
			} else {
				return r
			}
		}

		if s[m] == value {
			return m
		}

		if value < s[m] {
			r = m
		}
		if value > s[m] {
			l = m
		}
	}
}

func InsertionSortLinear(s []int) {
	for i := 1; i < len(s); i++ {

		// Getting insertion index:
		j := 0
		for ; s[j] < s[i]; j++ {
		}

		// Moving elements right:
		if j < i {
			temp := s[i]
			for k := i; k > j; k-- {
				s[k] = s[k-1]
			}
			s[j] = temp
		}
	}
}

func InsertionSortBisect(s []int) {
	for i := 1; i < len(s); i++ {

		// Getting insertion index:
		j := BinarySearch(s[0:i], s[i])

		// Moving elements right:
		if j < i {
			temp := s[i]
			for k := i; k > j; k-- {
				s[k] = s[k-1]
			}
			s[j] = temp
		}

	}
}

func main() {
	var s []int = []int{-23, -32, -2, -10}
	InsertionSortLinear(s)
	fmt.Println(s)
}
