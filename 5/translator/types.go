package translator

import (
	"bufio"
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

type Translator struct {
	from, to string
	ranges   []idRange
}

type idRange struct {
	dest   int
	source int
	len    int
}

func LoadAll(input string) ([]*Translator, []int) {
	var translators []*Translator
	var seeds []int

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

			seeds = make([]int, len(numsStr))

			var err error
			for i, nStr := range numsStr {
				seeds[i], err = strconv.Atoi(nStr)
				if err != nil {
					panic(fmt.Sprintf("invalid seeds number '%v'", nStr))
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

func (t *Translator) From() string {
	return t.from
}

func (t *Translator) To() string {
	return t.to
}

func (t *Translator) Map(source int) int {
	for _, r := range t.ranges {
		if source >= r.source && source <= r.source+r.len {
			return r.dest + source - r.source
		}
	}

	return source // fallback to direct translation
}
