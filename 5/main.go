package main

import (
	"fmt"
	"io"
	"os"

	"day5/router"
	"day5/translator"
)

const inFile = "input1.txt"

func main() {
	f, err := os.Open(inFile)
	if err != nil {
		panic(err)
	}

	input, err := io.ReadAll(f)
	if err != nil {
		panic(err)
	}

	trans, seeds := translator.LoadAll(string(input))
	routr := router.New(trans)

	fmt.Printf("Nearest seed: %v\n", routr.Nearest(seeds))
}
