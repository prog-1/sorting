package main

func insertSort(x []int) {
	for i := 1; i < len(x); i++ {
		for j := i; j > 0; j-- {
			if x[j-1] > x[j] { // <-- on this sign depends if slice will be sorted in ascending or descending order
				x[j-1], x[j] = x[j], x[j-1]
			}
		}
	}
}

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
	// what is the most important in sorting algorythms?
	// How to know which is more effective?°
	//BenchmarkInsertSort-2      	 2295277	       530.6 ns/op	       0 B/op	       0 allocs/op
	//BenchmarkInsertions-2      	 8224591	       147.5 ns/op	       0 B/op	       0 allocs/op
	//BenchmarkSelectionSort-2   	 1417420	       839.0 ns/op	       0 B/op	       0 allocs/op
}
