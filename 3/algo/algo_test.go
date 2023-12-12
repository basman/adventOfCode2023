package algo

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"

	"day3/parse"
)

func processString(sketch string) int {
	r := strings.NewReader(sketch)
	ch := parse.ReadLines(r)
	numbers, symbols := scanNumsNSyms(ch)
	return computeSum(numbers, symbols)
}

func TestProcess(t *testing.T) {
	tbl := []struct {
		sketch string
		sum    int
	}{
		{"", 0},
		{`...*45...`, 45},
		{`...45%...`, 45},
		{`...*11%...`, 11},
		{`*11...`, 11},
		{`17-...`, 17},
		{`..*11`, 11},
		{`...11+`, 11},
		{`*11`, 11},
		{`11=`, 11},
		{`11=11`, 22},
		{`..11=11..`, 22},
		{"..11.12..\n.-.......", 11},
		{".-.......\n..11.12..", 11},
		{"..11.12..\n-........", 0},
		{"-........\n..11.12..", 0},
		{"..11.12..\n.....-...", 12},
		{".....-...\n..11.12..", 12},
		{"..11.12..\n....-....", 23},
		{"....-....\n..11.12..\n", 23},
		{"467..114..\n...*......\n..35..633.\n......#...\n617*......\n.....+.58.\n..592.....\n......755.\n...$.*....\n.664.598..", 4361},
	}

	for i, tt := range tbl {
		t.Run(fmt.Sprintf("%v %v", i, tt.sum), func(t *testing.T) {
			sum := processString(tt.sketch)
			assert.Equal(t, tt.sum, sum)
		})
	}
}
