package geohex

import (
	"fmt"
	"math"
)

const VERSION = "3.2.0"

// MaxLevel is the maximum encoding level that this implementation supports
const MaxLevel = 20

// Error types
var (
	ErrLevelInvalid = fmt.Errorf("geohex: level invalid")
	ErrCodeInvalid  = fmt.Errorf("geohex: code invalid")
)

// Encode encodes a lat/lon/level into a code.
// Can return an ErrLevelInvalid code if level is not valid.
func Encode(lat, lon float64, level int) (string, error) {
	pnt := NewLL(lat, lon).Point()  // Point at lat/lon
	pos, err := pnt.Position(level) // Tile position
	if err != nil {
		return "", err
	}

	return pos.Code(), nil
}

// Decode decodes a string code into Lat/Lon coordinates
// Can return ErrCodeInvalid if code is  not valid
func Decode(code string) (LL, error) {
	pos, err := DecodePosition(code)
	if err != nil {
		return LL{}, err
	}
	return pos.LL(), nil
}

var (
	hChars = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	hIndex = make(map[byte]int, len(hChars))
	hK     = math.Tan(math.Pi / 6.0)
)

const (
	deg2Rad = math.Pi / 180.0
	pio2    = math.Pi / 2
)

var (
	// Precalculated math stuff
	pow3     [MaxLevel + 3]int
	halfPow3 [MaxLevel + 3]int
)

// zoom is a helper for level dimensions
type zoom struct {
	size int
}

// Cached zooms lookup
var zooms [MaxLevel + 1]*zoom

// Init zooms
func init() {
	for i := 0; i < MaxLevel+3; i++ {
		pow := math.Pow(3, float64(i))
		pow3[i] = int(pow)
		halfPow3[i] = int(math.Ceil(pow / 2))
	}

	for level := 0; level <= MaxLevel; level++ {
		zooms[level] = &zoom{size: pow3[level+2]}
	}

	for i, b := range hChars {
		hIndex[b] = i
	}
}
