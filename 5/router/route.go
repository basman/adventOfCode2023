package router

import (
	"fmt"

	"day5/translator"
)

type Router struct {
	translations map[string]map[string]*translator.Translator
}

func New(tr []*translator.Translator) *Router {
	r := &Router{
		translations: make(map[string]map[string]*translator.Translator),
	}

	for _, t := range tr {
		from, to := t.From(), t.To()

		m, ok := r.translations[from]
		if !ok {
			m = make(map[string]*translator.Translator)
			r.translations[from] = m
		}

		m[to] = t
	}

	return r
}

const (
	source = "seed"
	target = "location"
)

func (r *Router) Nearest(seeds []translator.SeedRange) int {
	minDist := -1
	minSeed := -1
	for _, seedRange := range seeds {
		for seedId := seedRange.Start(); seedId < seedRange.Start()+seedRange.Length(); seedId++ {
			dist := r.route(seedId, source, "")
			if dist < minDist || minDist < 0 {
				minDist = dist
				minSeed = seedId
			}
		}
	}

	if minSeed < 0 {
		panic("could not find nearest seed id")
	}

	return minDist
}

func (r *Router) route(id int, from string, to string) int {
	if to == "" {
		to = r.findTranslationTarget(from)
	}

	var destId = -1
	if tr, ok := r.translations[from][to]; !ok {
		panic(fmt.Sprintf("no translation found for %v->%v", from, to))
	} else {
		destId = tr.Map(id)

		if to == target {
			return destId
		}
	}

	if destId < 0 {
		panic("no destination id found")
	}

	nextSource := to
	nextTarget := r.findTranslationTarget(to)

	return r.route(destId, nextSource, nextTarget)
}

func (r *Router) findTranslationTarget(src string) string {
	m, ok := r.translations[src]

	if !ok {
		panic(fmt.Sprintf("no translation from %v to anywhere", src))
	}

	nextTarget := ""
	for cand := range m {
		if nextTarget == "" {
			nextTarget = cand
			continue
		}

		panic(fmt.Sprintf("ambiguous translation found: %v->%v vs %v->%v", src, nextTarget, src, cand))
	}
	return nextTarget
}
