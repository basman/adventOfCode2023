package reflection

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"

	"day13/pattern"
)

func TestFindVerticalAxis(t *testing.T) {
	tbl := []struct {
		pattern pattern.Pattern
		leftIdx int
	}{
		{pattern.New([]string{
			"#..#",
			"#..#",
			"#..#",
		}),
			1,
		},
		{pattern.New([]string{
			".#..#",
			".#..#",
			".#..#",
		}),
			2,
		},
		{pattern.New([]string{
			"#..#",
			"##.#",
			"#..#",
		}),
			-1,
		},
		{pattern.New([]string{
			".#..#",
			".##.#",
			".#..#",
		}),
			-1,
		},
		{pattern.New([]string{
			"#.#.#.##.#.",
			"......##...",
			".#..#....#.",
			"#..#..##..#",
		}),
			6,
		},
	}

	for i, tt := range tbl {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			leftIdx := FindVerticalAxis(tt.pattern)

			assert.Equal(t, tt.leftIdx, leftIdx)
		})
	}

}

func TestFindHorizontalAxis(t *testing.T) {
	tbl := []struct {
		pattern  pattern.Pattern
		lowerIdx int
	}{
		{pattern.New([]string{
			"#.##",
			"#.##",
			"....",
		}),
			0,
		},
		{pattern.New([]string{
			"....",
			"#.##",
			"#.##",
		}),
			1,
		},
		{pattern.New([]string{
			"....",
			"#.##",
			"#.##",
			"....",
		}),
			1,
		},
		{pattern.New([]string{
			".#..",
			"....",
			"#.##",
			"#.##",
			"....",
		}),
			2,
		},
		{pattern.New([]string{
			"....",
			"....",
			"#.##",
			"#.##",
		}),
			0,
		},
		{pattern.New([]string{
			"....",
			"....",
			"#..#",
			"#.##",
		}),
			-1,
		},
	}

	for i, tt := range tbl {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			lowerIdx := FindHorizontalAxis(tt.pattern)

			assert.Equal(t, tt.lowerIdx, lowerIdx)
		})
	}

}
