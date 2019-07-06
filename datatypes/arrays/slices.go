package main

import (
	"fmt"
	"strings"
)

func ArrayAndSlicePointers() {
	ArrayAndDerivedSlices()
	SlicesAndSubSlices()
	SlicesGrowing()
	SlicesShrinking()
}

func ArrayAndDerivedSlices() {
	// Array (a1)
	a1 := [5]int{1, 2, 3, 4, 5}

	// Slices (s1, s2) from (a1)
	s1 := a1[:3]
	s2 := a1[:4]

	fmt.Println()
	fmt.Println("SLICES OF A PREDEFINED ARRAY")
	fmt.Println(strings.Repeat("-", 60))
	fmt.Printf("Array (a1) pointer %p len %d cap %d value %v \n", &a1[0], len(a1), cap(a1), a1)
	fmt.Printf("Slice (s1) pointer %p len %d cap %d value %v \n", &s1[0], len(s1), cap(s1), s1)
	fmt.Printf("Slice (s2) pointer %p len %d cap %d value %v \n", &s2[0], len(s2), cap(s2), s2)
	fmt.Println()

	fmt.Println()
	fmt.Println("GROWING SLICES OF A PREDEFINED ARRAY")
	fmt.Println(strings.Repeat("-", 60))
	for _, v := range []int{6, 7, 8, 9} {
		s1 = append(s1, v)
		s1[0] = v
		fmt.Printf("Array (a1) pointer %p len %d cap %d value %v \n", &a1[0], len(a1), cap(a1), a1)
		fmt.Printf("Slice (s1) pointer %p len %d cap %d value %v \n", &s1[0], len(s1), cap(s1), s1)
	}
	fmt.Println()
}

func SlicesAndSubSlices() {
	// Slices
	s3 := []int{5, 4, 3, 2, 1}
	s4 := s3[2:]

	fmt.Println()
	fmt.Println("SLICES AND SUB-SLICES FROM IMPLICIT ARRAY")
	fmt.Println(strings.Repeat("-", 60))
	fmt.Printf("Slice (s3) pointer %p len %d cap %d value %v \n", &s3[2], len(s3), cap(s3), s3)
	fmt.Printf("Slice (s4) pointer %p len %d cap %d value %v \n", &s4[0], len(s4), cap(s4), s4)
	fmt.Println()
}

func SlicesGrowing() {
	// Slices - Growing
	s5 := []int{0}
	s6 := append(s5, 1)

	fmt.Println()
	fmt.Println("GROWING SLICES FROM IMPLICIT ARRAY (DOUBLING THE CAPACITY)")
	fmt.Println(strings.Repeat("-", 60))
	fmt.Printf("Slice (s5) pointer %p len %d cap %d value %v \n", &s5[0], len(s5), cap(s5), s5)
	s5[0] = 9
	fmt.Printf("Slice (s5) pointer %p len %d cap %d value %v \n", &s5[0], len(s5), cap(s5), s5)
	fmt.Println()

	var s7 []int

	for i := 2; i <= 16; i++ {
		if cap(s6) == 4 {
			s7 = s6[:]
			// Slices (s6) and (s7) points to
			// the same underline array so
			// this change will affect s6
			s7[0] = 7
		}
		if cap(s6) == 8 {
			// Slices (s7) still points to its original
			// array but (s6) is now pointing to a new
			// array as its capacity has changed so
			// this change will NOT affect s6 anymore
			s7[0] = 77
		}
		fmt.Printf("Slice (s6) pointer %p len %d cap %d value %v \n", &s6[0], len(s6), cap(s6), s6)
		s6 = append(s6, i)
	}
	fmt.Printf("Slice (s7) pointer %p len %d cap %d value %v \n", &s7[0], len(s7), cap(s7), s7)
	fmt.Println()
}

func SlicesShrinking() {
	// Slices - Shrinking
	s8 := []int{0, 1, 2, 3, 4, 5, 6, 7, 8, 9, 10, 11, 12, 13, 14, 15}

	fmt.Println()
	fmt.Println("SHRINKING SLICES FROM IMPLICIT ARRAY (FIXED CAPACITY)")
	fmt.Println(strings.Repeat("-", 60))
	for i := 0; i < 16; i++ {
		fmt.Printf("Slice (s8) pointer %p len %d cap %d value %v \n", &s8[0], len(s8), cap(s8), s8)
		s8 = s8[:len(s8)-1]
	}
	fmt.Println()
}
