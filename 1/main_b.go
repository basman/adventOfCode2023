package main

import (
	"fmt"
	"strconv"

	"private/adventOfCode2023/1/parse"
)

var dm = map[string]int{
	"one":   1,
	"two":   2,
	"three": 3,
	"four":  4,
	"five":  5,
	"six":   6,
	"seven": 7,
	"eight": 8,
	"nine":  9,
}

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
			var digit int
			if r >= '0' && r <= '9' {
				digit, err = strconv.Atoi(l[i : i+1])
				if err != nil {
					panic(err)
				}

				if firstDigit < 0 {
					firstDigit = digit
				} else {
					lastDigit = digit
				}
			}

			for name, val := range dm {
				n := len(name)
				if i <= len(l)-n && l[i:i+n] == name {
					digit = val
					i += n - 1

					if firstDigit < 0 {
						firstDigit = digit
					} else {
						lastDigit = digit
					}

					break
				}
			}
		}

		if lastDigit < 0 {
			lastDigit = firstDigit
		}

		if firstDigit < 0 || lastDigit < 0 {
			panic("digit not found")
		}

		fmt.Printf(" %v -> %v\n", l, firstDigit*10+lastDigit)

		sum += firstDigit*10 + lastDigit
	}

	fmt.Println(sum)
}
