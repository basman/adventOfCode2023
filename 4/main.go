package main

import (
	"fmt"

	"day4/entity"
	"day4/parse"
)

func main() {
	lines := parse.ReadLinesFile("input1.txt")
	cards := entity.LoadCards(lines)

	fmt.Println(cards)
}
