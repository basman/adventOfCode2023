package entity

import (
	"fmt"
	"strconv"

	"github.com/dlclark/regexp2"
)

type Card struct {
	Id        int
	WinNums   []int
	GivenNums []int
}

func LoadCards(lines chan string) []Card {
	var cards []Card

	for l := range lines {
		cards = append(cards, ParseCard(l))
	}

	return cards
}

var cardRe1 = regexp2.MustCompile(`Card\s*(\d+):(?:\s*(\d+))+\s+\|(?:\s*(\d+))+`, 0)

func ParseCard(l string) Card {
	m, err := cardRe1.FindStringMatch(l)
	if err != nil {
		panic(err)
	}

	if m == nil {
		panic(fmt.Sprintf("card line mismatch '%v'", l))
	}

	idStr := m.Groups()[1].String()
	id, err := strconv.Atoi(idStr)
	if err != nil {
		panic(err)
	}

	if len(m.Groups()) != 4 {
		panic("wrong number of capture groups")
	}

	wins := captures2sliceInt(m.Groups()[2].Captures)
	givens := captures2sliceInt(m.Groups()[3].Captures)

	return Card{
		Id:        id,
		WinNums:   wins,
		GivenNums: givens,
	}
}

func captures2sliceInt(captures []regexp2.Capture) []int {
	var nums []int
	for _, str := range captures {
		num, err := strconv.Atoi(str.String())
		if err != nil {
			panic(err)
		}
		nums = append(nums, num)
	}
	return nums
}
