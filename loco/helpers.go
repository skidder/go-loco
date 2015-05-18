package loco

import (
	"math"
	"strconv"

	"github.com/golang/geo/s2"
)

func GenerateGeohashFromLatLng(lat float64, lng float64) (geohash int64) {
	p := s2.PointFromLatLng(s2.LatLngFromDegrees(lat, lng))
	return GenerateGeohash(p)
}

func GenerateGeohash(p s2.Point) (geohash int64) {
	ll := s2.LatLngFromPoint(p)
	geohash = int64(s2.CellIDFromLatLng(ll))
	return
}

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

func round(f float64) float64 {
	return math.Floor(f + .5)
}
