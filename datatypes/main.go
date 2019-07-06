package main

import (
	"fmt"
	"strings"
)

func main() {
	PrintZeroedValues()
	PrintZeroedPointers()
	PrintCustomTypes()
	fmt.Println()
}

func PrintZeroedValues() {
	var numInt int
	var numUInt uint
	var str string
	var bufArr [10]byte
	var bufSlice []byte

	fmt.Println()
	fmt.Println(" Zeroed values")
	fmt.Println(" " + strings.Repeat("-", 30))

	fmt.Printf(" Type %-12T Value \"%v\"\n", numInt, numInt)
	fmt.Printf(" Type %-12T Value \"%v\"\n", numUInt, numUInt)
	fmt.Printf(" Type %-12T Value \"%v\"\n", str, str)
	fmt.Printf(" Type %-12T Value \"%v\"\n", bufArr, bufArr)
	fmt.Printf(" Type %-12T Value \"%v\"\n", bufSlice, bufSlice)
}

func PrintZeroedPointers() {
	var pointerInt *int
	var pointerStr *string
	var pointerArr *[10]byte
	var pointerSlice *[]byte

	fmt.Println()
	fmt.Println(" Zeroed pointers")
	fmt.Println(" " + strings.Repeat("-", 30))

	fmt.Printf(" Type %-12T Value \"%v\"\n", pointerInt, pointerInt)
	fmt.Printf(" Type %-12T Value \"%v\"\n", pointerStr, pointerStr)
	fmt.Printf(" Type %-12T Value \"%v\"\n", pointerArr, pointerArr)
	fmt.Printf(" Type %-12T Value \"%v\"\n", pointerSlice, pointerSlice)
}

func PrintCustomTypes() {
	type tInt int
	type tStr string
	type tArr [10]byte
	type tSlice []byte
	type tStruct struct {
		Field1 tInt
		Field2 tStr
		Field3 tArr
	}

	var num tInt
	var str tStr
	var arr tArr
	var slc tSlice
	var rec1 tStruct
	var rec2 = tStruct{}
	pRec1 := &rec1
	pRec2 := &rec2

	fmt.Println()
	fmt.Println(" Custom Types - Zeroed")
	fmt.Println(" " + strings.Repeat("-", 30))

	fmt.Printf(" Type %-12T Value \"%v\"\n", num, num)
	fmt.Printf(" Type %-12T Value \"%v\"\n", str, str)
	fmt.Printf(" Type %-12T Value \"%v\"\n", arr, arr)
	fmt.Printf(" Type %-12T Value \"%v\"\n", slc, slc)
	fmt.Printf(" Type %-12T Value \"%v\"\n", rec1, rec1)
	fmt.Printf(" Type %-12T Value \"%v\"\n", rec2, rec2)
	fmt.Printf(" Type %-12T Value \"%v\"\n", pRec1, pRec1)
	fmt.Printf(" Type %-12T Value \"%v\"\n", pRec2, pRec2)

	fmt.Println()
	fmt.Println(" Custom Types - BE CAREFUL")
	fmt.Println(" " + strings.Repeat("-", 30))

	type tError error
	err1 := new(error)
	err2 := new(tError)

	fmt.Printf(" Type %-12T EqualNIL(%v) Value \"%v\"\n", err1, err1 == nil, err1)
	fmt.Printf(" Type %-12T EqualNIL(%v) Value \"%v\"\n", err2, err2 == nil, err2)

	// num = new(tInt)
	// str = new(tStr)
	// arr = new(tArr)
	// slc = new(tSlice)

	// fmt.Println()
	// fmt.Println(" Custom Types - Initialized")
	// fmt.Println(" " + strings.Repeat("-", 30))

	// fmt.Printf(" Type %-12T Value \"%v\"\n", num, num)
	// fmt.Printf(" Type %-12T Value \"%v\"\n", str, str)
	// fmt.Printf(" Type %-12T Value \"%v\"\n", arr, arr)
	// fmt.Printf(" Type %-12T Value \"%v\"\n", slc, slc)
}
