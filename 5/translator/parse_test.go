package translator

import (
	"fmt"
	"reflect"
	"testing"
)

func TestLoadAll_seeds(t *testing.T) {
	tbl := []struct {
		seeds      string
		seedRanges []SeedRange
	}{
		{`seeds: 3127166940 109160474 3265086325 86449584 1581539098 205205726 3646327835 184743451 2671979893 17148151 305618297 40401857 2462071712 203075200 358806266 131147346 1802185716 538526744 635790399 705979250`,
			[]SeedRange{
				SeedRange{3127166940, 109160474},
				SeedRange{3265086325, 86449584},
				SeedRange{1581539098, 205205726},
				SeedRange{3646327835, 184743451},
				SeedRange{2671979893, 17148151},
				SeedRange{305618297, 40401857},
				SeedRange{2462071712, 203075200},
				SeedRange{358806266, 131147346},
				SeedRange{1802185716, 538526744},
				SeedRange{635790399, 705979250},
			},
		},
		{`seeds: 79 14 55 13`,
			[]SeedRange{
				{79, 14},
				{55, 13},
			},
		},
	}

	for i, tt := range tbl {
		t.Run(fmt.Sprint(i), func(t *testing.T) {
			_, seedRanges := LoadAll(tt.seeds)

			if len(seedRanges) != len(tt.seedRanges) {
				t.Error("seed range count differs")
			}

			if !reflect.DeepEqual(seedRanges, tt.seedRanges) {
				t.Error("seed ranges differ")
			}
		})
	}
}
