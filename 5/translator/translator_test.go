package translator

import (
	"fmt"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestTranslator_Map(t *testing.T) {
	tr := Translator{
		from: "a",
		to:   "b",
		ranges: []idRange{
			{1, 11, 9},
			{45, 55, 5},
		},
	}

	tbl := []struct {
		src, dst int
	}{
		{11, 1},
		{17, 7},
		{19, 9},
		{20, 20},
		{2, 2},
		{45, 45},
		{55, 45},
	}

	for i, tt := range tbl {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			dst := tr.Map(tt.src)

			assert.Equal(t, tt.dst, dst)
		})
	}
}
