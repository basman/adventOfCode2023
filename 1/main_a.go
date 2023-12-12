package main

import (
	"fmt"
	"strconv"

	"private/adventOfCode2023/1/parse"
)

func main() {
	lines, err := parse.ReadLines("input_a2.txt")
	if err != nil {
		panic(err)
	}

	sum := 0
	for l := range lines {
		firstDigit := -1
		lastDigit := -1

		for i, r := range l {
			if r >= '0' && r <= '9' {
				digit, err := strconv.Atoi(l[i : i+1])
				if err != nil {
					panic(err)
				}

				if firstDigit < 0 {
					firstDigit = digit
				} else {
					lastDigit = digit
				}
			}
		}

		if lastDigit < 0 {
			lastDigit = firstDigit
		}

		if firstDigit < 0 || lastDigit < 0 {
			panic("digit not found")
		}

		sum += firstDigit*10 + lastDigit
	}

	fmt.Println(sum)
}
