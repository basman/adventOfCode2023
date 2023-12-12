package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"

	"day2/parse"
)

type Game struct {
	Id      int
	Subsets []DiceSet
}

type DiceSet struct {
	Blue  int
	Red   int
	Green int
}

func (d DiceSet) String() string {
	return fmt.Sprintf("%v red, %v green, %v blue", d.Red, d.Green, d.Blue)
}

func main() {
	var games []Game

	for l := range parse.ReadLines("input_a2.txt") {
		games = append(games, parseGame(l))
	}

	sum := sumup(games)

	fmt.Printf("%v+\n", games)
	fmt.Printf("sum: %v\n", sum)

	m := power(games)

	fmt.Printf("total power: %v\n", m)
}

func power(games []Game) int {
	total := 0
	for _, g := range games {
		set := DiceSet{}
		for _, s := range g.Subsets {
			set.Red = max(set.Red, s.Red)
			set.Green = max(set.Green, s.Green)
			set.Blue = max(set.Blue, s.Blue)
		}

		p := set.Red * set.Green * set.Blue
		total += p
	}
	return total
}

func sumup(games []Game) int {
	sum := 0
	for _, g := range games {
		good := true
		for _, s := range g.Subsets {
			if s.Red > 12 || s.Green > 13 || s.Blue > 14 {
				good = false
				break
			}
		}

		if good {
			sum += g.Id
		}
	}

	return sum
}

func parseGame(l string) Game {
	re1 := regexp.MustCompile(`^Game (\d+): (.*)$`)

	g1 := re1.FindStringSubmatch(l)
	if len(g1) != 3 {
		panic(fmt.Sprintf("g1 mismatched '%v'", l))
	}

	id, err := strconv.Atoi(g1[1])
	if err != nil {
		panic("failed to convert id")
	}

	subsets := parseSubsets(g1[2])

	return Game{id, subsets}
}

func parseSubsets(in string) []DiceSet {
	sets := strings.Split(in, "; ")

	if len(sets) < 1 {
		panic(fmt.Sprintf("empty subset from '%v'", in))
	}

	var subsets []DiceSet

	re2 := regexp.MustCompile(`(\d+) (blue|red|green)(?:, )?`)
	for _, s := range sets {
		m := re2.FindAllStringSubmatch(s, -1)
		if len(m) < 1 {
			panic(fmt.Sprintf("empty dice from '%v'", s))
		}

		dset := DiceSet{}
		for _, dice := range m {
			if len(dice) != 3 {
				panic("dice mismatch")
			}

			diceVal, err := strconv.Atoi(dice[1])
			if err != nil {
				panic(fmt.Sprintf("parse dice value '%v' failed", dice[1]))
			}

			color := dice[2]

			switch color {
			case "red":
				if dset.Red > 0 {
					panic("red already set")
				}
				dset.Red = diceVal
			case "blue":
				if dset.Blue > 0 {
					panic("blue already set")
				}
				dset.Blue = diceVal
			case "green":
				if dset.Green > 0 {
					panic("green already set")
				}
				dset.Green = diceVal
			default:
				panic(fmt.Sprintf("unknown color '%v'", color))
			}
		}

		subsets = append(subsets, dset)
	}

	return subsets
}
