package translator

type SeedRange struct {
	start int
	len   int
}

func (r *SeedRange) Start() int {
	return r.start
}

func (r *SeedRange) Length() int {
	return r.len
}
