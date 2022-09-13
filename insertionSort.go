package main

import "fmt"

func insertionSort(arr []int) {
	for i := 1; i < len(arr); i++ {
		tmp := arr[i]
		j := i - 1
		for ; j > -1 && arr[j] >= tmp; j-- {
			arr[j+1] = arr[j]
		}
		arr[j+1] = tmp
	}
}

func main() {
	arr := []int{4, 8, 2, 1, 7, 4, 3}
	insertionSort(arr)
	fmt.Println(arr)
}
