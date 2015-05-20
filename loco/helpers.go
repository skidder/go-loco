package loco

import (
	"math"
	"strconv"

	"github.com/golang/geo/s2"
)

func FindCellIDs(area s2.Rect) (cellUnion s2.CellUnion) {
	cellQueue := make([]s2.CellID, 0)
	c := s2.CellIDFromFacePosLevel(0, 0, 0).ChildBeginAtLevel(0)
	endCellId := s2.CellIDFromFacePosLevel(5, 0, 0).ChildEndAtLevel(0)
	for c != endCellId {
		if containsGeodataToFind(c, area) {
			cellQueue = append(cellQueue, c)
		}
	}
	return s2.CellUnion{}
}

// Merge continuous cells in CellUnion and return a list of merged GeohashRanges.
func MergeCells(cellUnion s2.CellUnion) (ranges []GeohashRange) {
	ranges = make([]GeohashRange, 0)
	var cellId s2.CellID
	for _, cellId = range cellUnion {
		currentRange := GeohashRange{RangeMin: cellId.RangeMin(), RangeMax: cellId.RangeMax()}
		wasMerged := false

		for _, r := range ranges {
			if wasMerged = r.TryMerge(&currentRange); wasMerged == true {
				break
			}
		}

		if !wasMerged {
			ranges = append(ranges, currentRange)
		}
	}
	return
}

// Generate a geohash for the given latitude & longitude
func GenerateGeohashFromLatLng(lat float64, lng float64) (geohash int64) {
	p := s2.PointFromLatLng(s2.LatLngFromDegrees(lat, lng))
	return GenerateGeohash(p)
}

// Generate a geohash for the supplied point
func GenerateGeohash(p s2.Point) (geohash int64) {
	ll := s2.LatLngFromPoint(p)
	geohash = int64(s2.CellIDFromLatLng(ll))
	return
}

// Generate a DynamoDB hashkey of the given length for the supplied geohash
func GenerateHashKey(geohash int64, hashKeyLength int) (hashkey int64) {
	if geohash < 0 {
		// Counteract "-" at beginning of geohash
		hashKeyLength = hashKeyLength + 1
	}
	geohashString := strconv.FormatInt(geohash, 10)
	denominator := round(math.Pow(10, float64(len(geohashString)-hashKeyLength)))
	hashkey = int64(round(float64(geohash) / denominator))
	return
}

func processQueue(queue []s2.CellID, r s2.Rect) (cells []s2.CellID) {
	cellIds := make([]s2.CellID, 0)
	for _, c := range queue {
		if !c.IsValid() {
			break
		}

		cellIds = append(cellIds, processChildren(c, r, queue)...)
	}
	return
}

func processChildren(parent s2.CellID, r s2.Rect, queue []s2.CellID) (cellIds []s2.CellID) {
	children := make([]s2.CellID, 0)
	for c := parent.ChildBegin(); c != parent.ChildEnd(); c = c.Next() {
		if containsGeodataToFind(c, r) {
			children = append(children, c)
		}
	}

	cellIds = make([]s2.CellID, 0)
	switch len(children) {
	case 1:
	case 2:
		for _, child := range children {
			if child.IsLeaf() {
				cellIds = append(cellIds, child)
			} else {
				queue = append(queue, child)
			}
		}
		break
	case 3:
		cellIds = append(cellIds, children...)
		break
	case 4:
		cellIds = append(cellIds, parent)
		break
	default:
	}
	return
}

func containsGeodataToFind(c s2.CellID, r s2.Rect) (intersects bool) {
	for edgeID := 0; edgeID <= 3; edgeID += 1 {
		if r.ContainsLatLng(s2.LatLngFromPoint(s2.CellFromCellID(c).Vertex(edgeID))) {
			return true
		}
	}
	return false
}

func round(f float64) float64 {
	return math.Floor(f + .5)
}
