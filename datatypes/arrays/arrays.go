package main

import (
	"fmt"
)

func main() {
	ArrayAndSlicePointers()
}

// RunSamples runs and print the samples.
func RunSamples() {
	arr := Array1()
	fmt.Println(arr)
	slc1 := Slice1(arr, 0, 5)
	fmt.Println(slc1)
	slc2 := Slice1(arr, 5, 5)
	fmt.Println(slc2)
	slc2[0] = 888
	slc3 := slc2
	slc3[1] = 999

	fmt.Println("Operations Results")
	fmt.Println(slc2)
	fmt.Println(slc3)
	fmt.Println(&slc3)
	fmt.Println(&slc2 == &slc3)
	fmt.Println(arr)
}

// Array1 creates a 10 elements int array.
func Array1() [10]int {
	counter := 100
	var arr [10]int
	for i := 0; i < len(arr); i++ {
		arr[i] = counter
		counter++
	}
	return arr
}

// Slice1 creates a slice from a given 10 elements array
func Slice1(arr [10]int, idx, total int) []int {
	return arr[idx : idx+total]
}
