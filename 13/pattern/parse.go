package pattern

import "strings"

func LoadPatterns(lines chan string) []Pattern {
	var patterns []Pattern

	var pat []string
	for l := range lines {
		l = strings.TrimSpace(l)

		if l == "" {
			patterns = append(patterns, New(pat))
			pat = []string{}
		} else {
			pat = append(pat, l)
		}
	}

	if len(pat) > 0 {
		patterns = append(patterns, New(pat))
	}

	return patterns
}

func New(rows []string) Pattern {
	p := Pattern{
		Fields: make([][]uint8, len(rows)),
	}

	for r, row := range rows {
		p.Fields[r] = make([]uint8, len(row))
		for c := 0; c < len(row); c++ {
			p.Fields[r][c] = row[c]
		}
	}

	return p
}
