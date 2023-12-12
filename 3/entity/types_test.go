package entity

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestBox_Contains(t *testing.T) {
	tbl := []struct {
		b Box
		p Pos
		r bool
	}{
		// one square within
		{b: Box{Start: Pos{X: 0, Y: 0}, End: Pos{X: 0, Y: 0}}, p: Pos{X: 0, Y: 0}, r: true},
		// one square around
		{b: Box{Start: Pos{X: 0, Y: 0}, End: Pos{X: 0, Y: 0}}, p: Pos{X: -1, Y: 0}, r: false},
		{b: Box{Start: Pos{X: 0, Y: 1}, End: Pos{X: 0, Y: 0}}, p: Pos{X: 1, Y: 0}, r: false},
		{b: Box{Start: Pos{X: 0, Y: 0}, End: Pos{X: 0, Y: 0}}, p: Pos{X: 0, Y: -1}, r: false},
		{b: Box{Start: Pos{X: 0, Y: 0}, End: Pos{X: 0, Y: 0}}, p: Pos{X: 0, Y: 1}, r: false},

		// 4 squares within
		{b: Box{Start: Pos{X: 0, Y: 0}, End: Pos{X: 1, Y: 1}}, p: Pos{X: 0, Y: 0}, r: true},
		{b: Box{Start: Pos{X: 0, Y: 0}, End: Pos{X: 1, Y: 1}}, p: Pos{X: 1, Y: 0}, r: true},
		{b: Box{Start: Pos{X: 0, Y: 0}, End: Pos{X: 1, Y: 1}}, p: Pos{X: 1, Y: 1}, r: true},
		{b: Box{Start: Pos{X: 0, Y: 0}, End: Pos{X: 1, Y: 1}}, p: Pos{X: 0, Y: 1}, r: true},

		// 4 squares around
		{b: Box{Start: Pos{X: 0, Y: 0}, End: Pos{X: 1, Y: 1}}, p: Pos{X: -1, Y: 0}, r: false},
		{b: Box{Start: Pos{X: 0, Y: 0}, End: Pos{X: 1, Y: 1}}, p: Pos{X: -1, Y: -1}, r: false},
		{b: Box{Start: Pos{X: 0, Y: 0}, End: Pos{X: 1, Y: 1}}, p: Pos{X: 0, Y: -1}, r: false},
		{b: Box{Start: Pos{X: 0, Y: 0}, End: Pos{X: 1, Y: 1}}, p: Pos{X: 1, Y: -1}, r: false},
		{b: Box{Start: Pos{X: 0, Y: 0}, End: Pos{X: 1, Y: 1}}, p: Pos{X: 2, Y: -1}, r: false},
		{b: Box{Start: Pos{X: 0, Y: 0}, End: Pos{X: 1, Y: 1}}, p: Pos{X: 2, Y: 0}, r: false},
		{b: Box{Start: Pos{X: 0, Y: 0}, End: Pos{X: 1, Y: 1}}, p: Pos{X: 2, Y: 1}, r: false},
		{b: Box{Start: Pos{X: 0, Y: 0}, End: Pos{X: 1, Y: 1}}, p: Pos{X: 2, Y: 2}, r: false},
		{b: Box{Start: Pos{X: 0, Y: 0}, End: Pos{X: 1, Y: 1}}, p: Pos{X: 1, Y: 2}, r: false},
		{b: Box{Start: Pos{X: 0, Y: 0}, End: Pos{X: 1, Y: 1}}, p: Pos{X: 0, Y: 2}, r: false},
		{b: Box{Start: Pos{X: 0, Y: 0}, End: Pos{X: 1, Y: 1}}, p: Pos{X: -1, Y: 2}, r: false},
		{b: Box{Start: Pos{X: 0, Y: 0}, End: Pos{X: 1, Y: 1}}, p: Pos{X: -1, Y: 1}, r: false},
	}

	for _, tt := range tbl {
		t.Run(fmt.Sprintf("%v->%v", tt.p, tt.b), func(t *testing.T) {
			res := tt.b.Contains(tt.p)
			assert.Equal(t, tt.r, res)
		})
	}
}
