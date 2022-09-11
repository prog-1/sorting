package main

func bisectSearch(s []int, x int) int {
	start, end := 0, len(s)-1
	for start <= end {
		mid := (start + end) / 2
		if s[mid] == x {
			return mid
		} else if s[mid] < x {
			start = mid + 1
		} else {
			end = mid - 1
		}
	}
	return start
}

func insertionSortBisectSearch(s []int) {
	for i := 1; i < len(s); i++ {
		x := bisectSearch(s[:i+1], s[i])
		for j := i; j > x; j-- {
			s[j], s[j-1] = s[j-1], s[j]
		}
	}
}

func selectionSort(s []int) {
	for i := 0; i < len(s); i++ {
		minIndex := i // minIndex - index of the minimum number
		for j := i + 1; j < len(s); j++ {
			if s[minIndex] > s[j] {
				minIndex = j
			}
		}
		if i != minIndex {
			s[i], s[minIndex] = s[minIndex], s[i]
		}
	}
}

func insertionSortLinearSearch(s []int) {
	for i := 1; i < len(s); i++ {
		for j := i; j > 0 && s[j-1] > s[j]; j-- {
			s[j], s[j-1] = s[j-1], s[j]
		}
	}
}
