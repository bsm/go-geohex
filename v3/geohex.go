package geohex

import (
	"math"
)

const VERSION = "3.0.0"

// MaxLevel is the maximum encoding level that this implementation supports
const MaxLevel = 20

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

// Init zooms
func init() {
	for level := 0; level <= MaxLevel; level++ {
		size := hBase / math.Pow(3, float64(level+3))
		zooms[level] = &Zoom{level: level, size: size, scale: size / hEr, w: 6 * size, h: 6 * size * hK}
	}

	for i, b := range hChars {
		hIndex[b] = i
	}
}
