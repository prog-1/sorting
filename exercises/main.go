package main

import "fmt"

func insertSort(x []int) {
	for i := 1; i < len(x); i++ {
		for j := i; j > 0; j-- {
			if x[j-1] > x[j] { // <-- on this sign depends if slice will be sorted in ascending or descending order
				x[j-1], x[j] = x[j], x[j-1]
			}
		}
	}
}

func findInsertionIndex(x []int, a int) int {
	i, j := 0, len(x)-1

	for i = 0; i <= j; j-- {
		mid := i + j>>1
		fmt.Println(mid)
		if a < x[mid] {
			j = mid - 1
		} else {
			i = mid + 1
		}
	}
	return i
}
func insertBisSort(x []int) {
	for i := 1; i < len(x); i++ {
		indx := findInsertionIndex(x[0:i], x[i])
		temp := x[i]
		for j := i; j > indx; j-- {
			x[j] = x[j-1]
		}
		x[indx] = temp
	}
} //bisect insertion sort scetch

//Found this quite interesting piece on https://medium.com/star-gazers/insertion-sorting-algorithm-in-go-26f1ec49936c
//It was written on base of pseudocode
// But helped to understand
// No if statements? so must a lot loop iterations
func insertSortBySteveHook(x []int) []int {
	for j := 1; j < len(x); j++ {
		key := x[j]
		i := j - 1
		for i > -1 && x[i] < key {
			x[i+1] = x[i]
			i -= 1
		}
		x[i+1] = key
	}
	return x
} //not sure if it is effective tho

func selectionSort(x []int) {
	for i := 0; i < len(x)-1; i++ {
		var min = i
		for j := i + 1; j < len(x); j++ {
			if x[j] < x[min] {
				min = j
			}
		}
		x[i], x[min] = x[min], x[i]
	}
}

func main() {
	x := []int{5, 9, 7, 0, 3}
	insertBisSort(x)
	fmt.Println(x)
}
