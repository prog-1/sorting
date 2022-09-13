package main

import (
	"fmt"
	"reflect"
	"testing"
)

//SORTINGS:

type sortFunc struct {
	name string
	f    func([]int)
}

var sortFuncs = []sortFunc{
	{"bubbleSort", bubbleSort},
	{"insertionSort", insertionSort},
	{"insertionSortwithBS", insertionSortwithBS},
	{"selectionSort", selectionSort},
}

//TESTS:

func testSort(t *testing.T, sortFunc func([]int)) {

	var valArray = []values{ // array of value structures
		{[]int{4, 8, 2, 1}, []int{1, 2, 4, 8}},                       //usual values
		{[]int{67, 88, 202, 604, 110}, []int{67, 88, 110, 202, 604}}, //a bit bigger values
		{[]int{-44, -202, -510, -70}, []int{-510, -202, -70, -44}},   // negative values
		{[]int{5, 5}, []int{5, 5}},                                   // identical values
		{[]int{1}, []int{1}},                                         // one value
		{nil, nil},                                                   // nil
	}

	for i := 0; i < len(valArray); i++ { //comparing results for each value array element
		sortFunc(valArray[i].inpValues)
		if !reflect.DeepEqual(valArray[i].inpValues, valArray[i].expResult) { // Alternative: !compareSlices(result, valArray[i].expResult)
			t.Fatal("Dude, your sorting function", sortFuncs, "doesn't work. Got:", valArray[i].inpValues, ", expected:", valArray[i].expResult)
		} else {
			fmt.Println("Test with values", valArray[i].inpValues, "passed!")
		}

	}

}

func TestSort(t *testing.T) {
	for _, test := range sortFuncs {
		fmt.Println("Running", test.name)
		t.Run(test.name, func(t *testing.T) { testSort(t, test.f) })
		fmt.Println()
		fmt.Println()
	}
}

//BENCHMARKS:

func benchmarkSort(b *testing.B, sortFunc func([]int)) {
	arr := []int{43, 18, 104, 27, 12}
	for n := 0; n < b.N; n++ {
		sortFunc(arr)
	}
}

func BenchmarkSort(b *testing.B) {
	for _, bm := range sortFuncs {
		b.Run(bm.name, func(b *testing.B) { benchmarkSort(b, bm.f) })
	}
}
