package translator

type Translator struct {
	from, to string
	ranges   []idRange
}

type idRange struct {
	dest   int
	source int
	len    int
}

func (t *Translator) From() string {
	return t.from
}

func (t *Translator) To() string {
	return t.to
}

func (t *Translator) Map(source int) int {
	for _, r := range t.ranges {
		if source >= r.source && source < r.source+r.len {
			return r.dest + source - r.source
		}
	}

	return source // fallback to direct translation
}
