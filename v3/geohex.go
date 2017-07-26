package geohex

import (
	"errors"
	"math"
)

const VERSION = "3.2.0"

// MaxLevel is the maximum encoding level that this implementation supports
const MaxLevel = 20

// Error types
var (
	ErrLevelInvalid = errors.New("geohex: level invalid")
	ErrCodeInvalid  = errors.New("geohex: code invalid")
)

var (
	hChars = []byte{'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z', 'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	hIndex = make(map[byte]int, len(hChars))
	hK     = math.Tan(math.Pi / 6.0)

	// Precalculated math stuff
	pow3     [MaxLevel + 3]int
	halfPow3 [MaxLevel + 3]int
)

const (
	deg2Rad = math.Pi / 180.0
	pio2    = math.Pi / 2
)

// Cached zooms lookup
var sizes = make(map[uint8]int, 20)

// Init cache
func init() {
	for i := 0; i < MaxLevel+3; i++ {
		pow := math.Pow(3, float64(i))
		pow3[i] = int(pow)
		halfPow3[i] = int(math.Ceil(pow / 2))
	}

	for i, b := range hChars {
		hIndex[b] = i
	}

	for level := uint8(0); level <= MaxLevel; level++ {
		sizes[level] = pow3[level+2]
	}
}
