package main

import (
	"fmt"

	"day3/algo"
)

func main() {
	fname := "input_a3.txt"
	sum := algo.ProcessPart1(fname)
	fmt.Printf("sum: %v\n", sum)

	sum2 := algo.ProcessPart2(fname)
	fmt.Printf("gear ratio sum: %v\n", sum2)
}
