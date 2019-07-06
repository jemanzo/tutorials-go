package main

import (
	"fmt"
	"strconv"
)

func main() {
	var numberAsInterface interface{} = "123456.654321"

	numberAsString, ok := numberAsInterface.(string)
	if !ok {
		fmt.Printf("Error converting numberAsInterface to numberAsString")
	}

	numberAsFloat64, err := strconv.ParseFloat(numberAsString, 64)
	if err != nil {
		fmt.Printf("Error converting numberAsString to numberAsFloat64:\n  %e", err)
	}

	fmt.Printf("numberAsFloat64 - type %T value %f\n", numberAsFloat64, numberAsFloat64)
}
