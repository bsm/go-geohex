package geohex

import "math"

var (
	hChars = []byte{
		'A', 'B', 'C', 'D', 'E', 'F', 'G', 'H', 'I', 'J', 'K', 'L', 'M', 'N', 'O', 'P', 'Q', 'R', 'S', 'T', 'U', 'V', 'W', 'X', 'Y', 'Z',
		'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z',
	}
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

	// Cached sizes
	sizes [MaxLevel + 1]int
)

// zoom is a helper for level dimensions
type zoom struct {
	size int
}

// Init sizes
func init() {
	for i := 0; i < MaxLevel+3; i++ {
		pow := math.Pow(3, float64(i))
		pow3[i] = int(pow)
		halfPow3[i] = int(math.Ceil(pow / 2))
	}

	for level := 0; level <= MaxLevel; level++ {
		sizes[level] = pow3[level+2]
	}

	for i, b := range hChars {
		hIndex[b] = i
	}
}
