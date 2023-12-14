package main

import (
	"bufio"
	"fmt"
	"os"

	"day13/pattern"
	"day13/reflection"
)

// 1719: That's not the right answer; your answer is too low.
// 2696: That's not the right answer; your answer is too low.
const filename = "input2.txt"

func readFileLines(fname string) chan string {
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}

	ch := make(chan string)
	scan := bufio.NewScanner(f)

	go func() {
		for scan.Scan() {
			if scan.Err() != nil {
				panic(scan.Err())
			}

			ch <- scan.Text()
		}

		f.Close()
		close(ch)
	}()

	return ch
}

func search(patterns []pattern.Pattern) {
	colSum := 0
	rowSum := 0
	for i, p := range patterns {
		h := reflection.FindHorizontalAxis(p)
		if h >= 0 {
			fmt.Printf("Pattern %v horizontal axis between rows %v,%v", i+1, h+1, h+2)
			rowSum += h + 1
		}

		v := reflection.FindVerticalAxis(p)
		if v >= 0 {
			fmt.Printf("Pattern %v vertical axis between columns %v,%v", i+1, v+1, v+2)
			colSum += v + 1
		}
	}

	fmt.Printf("Sum of reflections (cols=%v, rows=%v): %v", colSum, rowSum, colSum+100*rowSum)
}

func main() {
	lineCh := readFileLines(filename)
	patterns := pattern.LoadPatterns(lineCh)
	search(patterns)
}
