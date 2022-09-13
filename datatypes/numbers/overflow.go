package main

import (
	"fmt"
	"math"
)

func ShowLimitsAndOverflow() {
	ShowLimitsOfINT32()
	ShowLimitsOfUINT32()
	ShowLimitsOfINT64()
	ShowLimitsOfUINT64()
}

// LIMITS OF SIGNED INTEGER - INT32
func ShowLimitsOfINT32() {
	fmt.Println()
	var num32 int32
	fmt.Printf("Limits of signed integer                   %T\n", num32)
	fmt.Println("================================================")
	fmt.Printf("    minimum       %d\n", math.MinInt32)
	fmt.Printf("    maximum       %d\n", math.MaxInt32)
	fmt.Printf("    2^32 / 2      %d\n", uint32(math.Pow(2, 32)/2))
	fmt.Printf("    2^32 / 2 - 1  %d\n", uint32(math.Pow(2, 32)/2-1))
	fmt.Printf("    overflow:\n")
	num32 = math.MinInt32
	fmt.Printf("      minimum     %d - 1 = %d\n", num32, num32-1)
	num32 = math.MaxInt32
	fmt.Printf("      maximum     %d + 1 = %d\n", num32, num32+1)
	fmt.Println()
}

// LIMITS OF UNSIGNED INTEGER - UINT32
func ShowLimitsOfUINT32() {
	fmt.Println()
	var numU32 uint32
	fmt.Printf("Limits of unsigned integer                %T\n", numU32)
	fmt.Println("================================================")
	fmt.Printf("    minimum       %d\n", 0)
	fmt.Printf("    maximum       %d\n", math.MaxUint32)
	fmt.Printf("    2^32          %d\n", uint32(math.Pow(2, 32)))
	fmt.Printf("    2^32 - 1      %d\n", uint32(math.Pow(2, 32)-1))
	fmt.Printf("    overflow:\n")
	numU32 = 0
	fmt.Printf("      minimum     %d - 1 = %d\n", numU32, numU32-1)
	numU32 = math.MaxUint32
	fmt.Printf("      maximum     %d + 1 = %d\n", numU32, numU32+1)
	fmt.Println()
}

// LIMITS OF SIGNED INTEGER - INT64
func ShowLimitsOfINT64() {
	fmt.Println()
	// LIMITS OF INT64 - SIGNED
	var num64 int64
	fmt.Printf("Limits of signed integer                   %T\n", num64)
	fmt.Println("================================================")
	fmt.Printf("    minimum       %d\n", math.MinInt64)
	fmt.Printf("    maximum       %d\n", math.MaxInt64)
	fmt.Printf("    2^64 / 2      %d\n", uint64(math.Pow(2, 64)/2))
	fmt.Printf("    2^64 / 2 - 1  %d\n", uint64(math.Pow(2, 64)/2-1))

	fmt.Printf("    overflow:\n")
	num64 = math.MinInt64
	fmt.Printf("      minimum     %d - 1 = %d\n", num64, num64-1)
	num64 = math.MaxInt64
	fmt.Printf("      maximum     %d + 1 = %d\n", num64, num64+1)

	fmt.Println()
}

// LIMITS OF UNSIGNED INTEGER - UINT64
func ShowLimitsOfUINT64() {
	fmt.Println()
	var numU64 uint64
	fmt.Printf("Limits of unsigned integer                %T\n", numU64)
	fmt.Println("================================================")
	fmt.Printf("    minimum       %d\n", 0)
	fmt.Printf("    maximum       %d\n", uint64(math.MaxUint64))
	fmt.Printf("    2^64          %d\n", uint64(math.Pow(2, 64)))
	fmt.Printf("    2^64 - 1      %d\n", uint64(math.Pow(2, 64)-1))
	fmt.Printf("    overflow:\n")
	numU64 = 0
	fmt.Printf("      minimum     %d - 1 = %d\n", numU64, numU64-1)
	numU64 = math.MaxUint64
	fmt.Printf("      maximum     %d + 1 = %d\n", numU64, numU64+1)
	fmt.Println()
}
