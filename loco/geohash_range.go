package loco

import (
	"github.com/golang/geo/s2"
)

type GeohashRange struct {
	RangeMin s2.CellID
	RangeMax s2.CellID
}

const (
	MERGE_THRESHOLD = 2
)

func (g *GeohashRange) TryMerge(otherRange *GeohashRange) (merged bool) {
	merged = false
	if otherRange.RangeMin-g.RangeMax < MERGE_THRESHOLD && otherRange.RangeMin-g.RangeMax > 0 {
		g.RangeMax = otherRange.RangeMax
		merged = true
		return
	}

	if g.RangeMin-otherRange.RangeMax <= MERGE_THRESHOLD && g.RangeMin-otherRange.RangeMax > 0 {
		g.RangeMin = otherRange.RangeMin
		merged = true
	}
	return
}
