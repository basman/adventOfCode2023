package reflection

import "day13/pattern"

// FindHorizontalAxis returns the row index above the symmetrical axis or -1 if not found
func FindHorizontalAxis(p pattern.Pattern) int {
	height := len(p.Fields)
	for startRow := 0; startRow < height/2; startRow++ {
		// make sure search range spans across even number of rows
		lastRow := height - 1
		if (lastRow-startRow+1)%2 != 0 {
			lastRow--
		}

		symmetry := true

	outer:
		for y := startRow; y <= lastRow; y++ {
			for x, me := range p.Fields[y] {
				you := p.Fields[startRow+lastRow-y][x]

				if me != you {
					symmetry = false
					break outer
				}
			}
		}

		if symmetry {
			return startRow + (lastRow-startRow)/2
		}
	}

	return -1
}

// FindVerticalAxis returns the column index left of the symmetrical axis or -1 if not found
func FindVerticalAxis(p pattern.Pattern) int {
	width := len(p.Fields[0])
	for startCol := 0; startCol < width/2; startCol++ {
		// make sure search range spans across even number of columns
		lastCol := width - 1
		if (lastCol-startCol+1)%2 != 0 {
			lastCol--
		}

		symmetry := true

	outer:
		for y := 0; y < len(p.Fields); y++ {
			for x := startCol; x <= startCol+(lastCol-startCol)/2; x++ {
				me := p.Fields[y][x]
				you := p.Fields[y][startCol+lastCol-x]

				if me != you {
					symmetry = false
					break outer
				}
			}
		}

		if symmetry {
			return startCol + (lastCol-startCol)/2
		}
	}

	return -1
}
