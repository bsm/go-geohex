package geohex

import (
	"math"
)

const VERSION = "3.0.0"

var (
	hChars = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	hIndex = make(map[byte]int, len(hChars))
	hK     = math.Tan(math.Pi / 6.0)
	hD2R   = math.Pi / 180.0
)

const (
	hBase = 20037508.34
	hEr   = 6371007.2
)

// A zoom is a helper for level dimensions
type Zoom struct {
	level int
	size  float64
	scale float64
	w     float64
	h     float64
}

// Cached zooms lookup
var zooms = make(map[int]*Zoom, 20)

// LL is a lat/lon tuple
type LL struct {
	Lat, Lon float64
}

// NewLL creates a new normalised LL
func NewLL(lat, lon float64) *LL {
	if lon < -180 {
		lon += 360
	} else if lon > 180 {
		lon -= 360
	}
	return &LL{Lat: lat, Lon: lon}
}

// Point generates a grid point from a lat/lon
func (ll *LL) Point() *Point {
	e := ll.Lon * hBase / 180.0
	n := math.Log(math.Tan((90+ll.Lat)*hD2R/2)) / math.Pi * hBase

	return &Point{E: e, N: n}
}

// Init zooms
func init() {
	for level := 0; level < 21; level++ {
		size := hBase / math.Pow(3, float64(level+3))
		zooms[level] = &Zoom{level: level, size: size, scale: size / hEr, w: 6 * size, h: 6 * size * hK}
	}

	for i, b := range hChars {
		hIndex[b] = i
	}
}
