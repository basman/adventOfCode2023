package parse

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func ReadLinesFile(filename string) chan string {
	var err error
	r := os.Stdin

	if filename != "" {
		r, err = os.Open(filename)
		if err != nil {
			panic(fmt.Sprintf("failed to open file %v: %v", filename, err))
		}
	}

	return ReadLines(r)
}

func ReadLines(r io.Reader) chan string {
	if r == nil {
		r = os.Stdin
	}

	s := bufio.NewScanner(r)

	out := make(chan string)

	go func() {
		for s.Scan() {
			out <- s.Text()
		}

		if s.Err() != nil {
			panic(s.Err())
		}

		close(out)
	}()

	return out
}
