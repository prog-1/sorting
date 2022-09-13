package main

func insertionSortwithBS(arr []int) {
	for i := 1; i < len(arr); i++ {
		tmp := arr[i]

		low := 0
		high := len(arr) - 1
		var mid int

		//binary search
		for low <= high {
			mid = (low + high) / 2

			if arr[mid] < arr[i] {
				low = mid + 1
			} else {
				high = mid - 1
			}
		}

		//moving elements to right
		for j := i; j > low; j-- {
			arr[j] = arr[j-1]
		}
		arr[low] = tmp
	}
}

// func main() {
// 	arr := []int{6, 1, 5, 8, 7}
// 	insertionSortwithBS(arr)
// 	fmt.Println(arr)
// }
