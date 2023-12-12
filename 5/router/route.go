package router

import (
	"fmt"
	"sync"
	"time"

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

/*
81956385
That's not the right answer; your answer is too high.
*/
func (r *Router) Nearest(seeds []translator.SeedRange) int {
	ch := make(chan int)
	wg := sync.WaitGroup{}

	for srId, seedRange := range seeds {
		wg.Add(1)
		go func(sr translator.SeedRange, id int) {
			fmt.Printf("seedrange %v launched, len=%v\n", id, sr.Length())
			start := time.Now()
			for seedId := sr.Start(); seedId < sr.Start()+sr.Length(); seedId++ {
				ch <- r.route(seedId, source, "")
			}
			fmt.Printf("seedrange %v completed in %v seconds\n", id, time.Now().Sub(start).Seconds())
			wg.Done()
		}(seedRange, srId)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	minDist := -1
	for dist := range ch {
		if dist < minDist || minDist < 0 {
			minDist = dist
		}
	}

	if minDist < 0 {
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
