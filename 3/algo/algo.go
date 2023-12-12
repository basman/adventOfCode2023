package algo

import (
	"fmt"
	"regexp"
	"strconv"

	"day3/entity"
	"day3/parse"
)

func ProcessPart1(filename string) int {
	ch := parse.ReadLinesFile(filename)
	numbers, symbols := scanNumsNSyms(ch)
	return computeSum(numbers, symbols)
}

func ProcessPart2(filename string) int {
	ch := parse.ReadLinesFile(filename)
	numbers, symbols := scanNumsNSyms(ch)
	return computeGearRatioSum(numbers, symbols)
}

func computeGearRatioSum(numbers []entity.Number, symbols []entity.Symbol) int {
	sum := 0
	for _, s := range symbols {
		if s.Symbol != '*' {
			continue
		}

		count := 0
		gearRatio := 0
		for _, n := range numbers {
			if n.Box.Contains(s.Pos) {
				count++
				if count == 1 {
					gearRatio = n.Value
				} else if count == 2 {
					gearRatio *= n.Value
				} else {
					break
				}
			}
		}

		if count == 2 {
			sum += gearRatio
		}
	}

	return sum
}

func computeSum(numbers []entity.Number, symbols []entity.Symbol) int {
	sum := 0
outer:
	for _, n := range numbers {
		for _, s := range symbols {
			if n.Box.Contains(s.Pos) {
				fmt.Printf("%v <- %v %v\n", n.Value, string(s.Symbol), s.Pos)
				sum += n.Value
				continue outer
			}
		}
	}
	return sum
}

func scanNumsNSyms(in chan string) (numbers []entity.Number, symbols []entity.Symbol) {
	numRe := regexp.MustCompile(`^\d+`)
	row := 0
	for l := range in {
		row++
		for col := 0; col < len(l); col++ {
			r := l[col]

			if r == '.' {
				continue // ignore empty spaces
			}

			p := entity.Pos{Y: row, X: col}

			if isDigit(rune(r)) {
				valStr := numRe.FindString(l[col:])
				n := len(valStr)
				if n < 1 {
					panic("failed extracting number")
				}

				num, err := strconv.Atoi(valStr)
				if err != nil {
					panic(fmt.Errorf("number conversion failed: %v", err))
				}

				if num < 1 || num > 9999 {
					panic("invalid number")
				}

				numBox := entity.Box{
					Start: entity.Pos{Y: row - 1, X: col - 1},
					End:   entity.Pos{Y: row + 1, X: col + n},
				}

				numbers = append(numbers, entity.Number{
					Pos:   p,
					Box:   numBox,
					Value: num,
				})

				col += n - 1 // skip remaining digits of current number
			} else {
				sym := entity.Symbol{
					Pos:    p,
					Symbol: rune(r),
				}

				symbols = append(symbols, sym)
			}
		}
	}

	fmt.Printf("found %v numbers, %v symbols\n", len(numbers), len(symbols))

	return
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
