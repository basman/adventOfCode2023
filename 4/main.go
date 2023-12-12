package main

import (
	"fmt"
	"math"
	"slices"

	"day4/entity"
	"day4/parse"
)

func sumUpPart2(cards []entity.Card) int {
	origCards := append([]entity.Card{}, cards...)
	total := 0

	for i := 0; i < len(cards); i++ {
		winCount := 0
		for _, w := range cards[i].WinNums {
			if slices.Contains(cards[i].GivenNums, w) {
				winCount++
			}
		}

		total++ // this card matched winning numbers

		if winCount == 0 {
			continue
		}

		cardId := cards[i].Id

		// insert copies
		tail := append([]entity.Card{}, cards[i+1:]...)                   // preserve tail
		cards = append(cards[:i+1], origCards[cardId:cardId+winCount]...) // insert copies
		cards = append(cards, tail...)                                    // append tail

		fmt.Printf("Card %v: %v copies [%v-%v]: %v\n", cardId, winCount, cardId+1, cardId+winCount, entity.Stack(cards))
	}

	return total
}

func sumUpPart1(cards []entity.Card) int {
	stackTotal := 0
	for _, card := range cards {
		winCount := 0
		for _, w := range card.WinNums {
			if slices.Contains(card.GivenNums, w) {
				winCount++
			}
		}

		points := powInt(2, winCount-1)
		fmt.Printf("Card %v: %v points\n", card.Id, points)
		stackTotal += points
	}

	return stackTotal
}

func powInt(x, y int) int {
	return int(math.Pow(float64(x), float64(y)))
}

func main() {
	lines := parse.ReadLinesFile("input2.txt")
	cards := entity.LoadCards(lines)

	//fmt.Println(sumUpPart1(append([]entity.Card{}, cards...)))
	fmt.Println(sumUpPart2(append([]entity.Card{}, cards...)))
}
