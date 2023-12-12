package parse

import (
	"bufio"
	"fmt"
	"os"
)

func ReadLines(filename string) (chan string, error) {
	var err error
	r := os.Stdin
	if filename != "" {
		r, err = os.Open(filename)
		if err != nil {
			panic(fmt.Sprintf("failed to open file %v: %v", filename, err))
		}
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

	return out, nil
}
