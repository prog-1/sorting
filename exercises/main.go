package main

func insertion(s []int) {
	for i := 1; i < len(s); i++ {
		for j := i; j > 0; j-- {
			if s[j-1] > s[j] {
				s[j-1], s[j] = s[j], s[j-1]
			}
		}
	}
}
func insertionBisect(s []int) {
	for i := 1; i < len(s); i++ {
		x, j, k := s[i], 0, i
		for j != k {
			m := (j + k) / 2
			if s[m] <= x {
				j = m + 1
			} else {
				k = m
			}
		}
		copy(s[j+1:], s[j:i])
		s[j] = x
	}
}
func selection(s []int) {
	for i := 0; i < len(s); i++ {
		min := i
		for j := i + 1; j < len(s); j++ {
			if s[j] < s[min] {
				min = j
			}
		}
		s[i], s[min] = s[min], s[i]
	}
}
func reverseIns(s []int) {
	for i := len(s) - 2; i >= 0; i-- {
		for j := i; j < len(s)-1; j++ {
			if s[j+1] < s[j] {
				s[j+1], s[j] = s[j], s[j+1]
			}
		}
	}
}
func main() {
}
