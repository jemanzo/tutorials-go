package main

import (
	"fmt"
)

type Point struct {
	x int
	y int
}

func main() {
	p1 := Point{1, 2}
	fmt.Println(p1)
	fmt.Printf("%p\n", &p1)
}
