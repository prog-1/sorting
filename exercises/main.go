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
func selection(s []int) {
	for i := 0; i < len(s); i++ {
		for j, min := i+1, i; j < len(s); j++ {
			if s[j] < s[min] {
				s[j], s[min] = s[min], s[j]
			}
		}
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
