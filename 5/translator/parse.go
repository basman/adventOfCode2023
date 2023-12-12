package translator

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

func LoadAll(input string) ([]*Translator, []SeedRange) {
	var translators []*Translator
	var seeds []SeedRange

	scan := bufio.NewScanner(strings.NewReader(input))

	var tr *Translator
	for scan.Scan() {
		l := scan.Text()
		l = strings.TrimSpace(l)

		if strings.HasPrefix(l, "seeds: ") {
			numsStr := strings.Split(l[7:], " ")
			if len(numsStr) < 1 {
				panic(fmt.Sprintf("invalid seeds list '%v'", l))
			}

			seeds = make([]SeedRange, len(numsStr)/2)

			var err error
			for i := 0; i < len(numsStr); i += 2 {
				seeds[i/2].start, err = strconv.Atoi(numsStr[i])
				if err != nil {
					panic(fmt.Sprintf("invalid seeds id start '%v'", numsStr[i]))
				}

				seeds[i/2].len, err = strconv.Atoi(numsStr[i+1])
				if err != nil {
					panic(fmt.Sprintf("invalid seeds id length '%v'", numsStr[i+1]))
				}
			}
		} else if strings.HasSuffix(l, "map:") {
			tr = &Translator{}
			translators = append(translators, tr)

			tr.from, tr.to = parseFromTo(l)
		} else if l == "" {
			tr = nil
			continue
		} else {
			var r idRange
			n, err := fmt.Sscanf(l, "%d %d %d", &r.dest, &r.source, &r.len)
			if err != nil {
				panic(err)
			}

			if n != 3 {
				panic(fmt.Sprintf("invalid map range '%v'", l))
			}

			tr.ranges = append(tr.ranges, r)
		}
	}

	return translators, seeds
}

var reFromTo = regexp.MustCompile(`^([^\-]+)-to-(\S+) map:`)

func parseFromTo(header string) (from string, to string) {
	m := reFromTo.FindStringSubmatch(header)
	if m == nil || len(m) != 3 || m[1] == "" || m[2] == "" {
		panic("header mismatch")
	}

	return m[1], m[2]
}
