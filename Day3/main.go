// Day3 project main.go
package main

import (
	"fmt"
	"math"
)

func day3(input int) int {
	ring := 0
	fmt.Printf("Input %d\n", input)
	ring = int(math.Floor(math.Sqrt(float64(input)))+1) / 2
	idx := int(math.Pow(float64(ring*2-1), 2))
	fmt.Printf("Ring %d\n", ring)
	idx = input - idx
	idx = int(math.Abs(float64((idx%(ring*2) - ring))))
	fmt.Printf("offset %d\n", idx)
	fmt.Printf("dist %d\n", idx+ring)
	return idx + ring
}

func main() {
	for i := 25; i < 30; i++ {
		day3(i)
	}
	day3(277678)
}
