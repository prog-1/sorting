package main

func selectionSort(arr []int) {
	for i := 0; i < len(arr); i++ {
		min := i
		for j := min + 1; j < len(arr); j++ {
			if arr[min] > arr[j] {
				min = j
			}
		}
		arr[i], arr[min] = arr[min], arr[i]
	}
}
