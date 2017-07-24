package geohex

import (
	"fmt"
	"math"
	"strconv"
)

// Error types
var (
	ErrLevelInvalid = fmt.Errorf("geohex: level invalid")
	ErrCodeInvalid  = fmt.Errorf("geohex: code invalid")
)

// A Zone represents a GeoHex tile
type Zone struct {
	Code string
	Pos  *Position
}

var (
	// Precalculated math stuff
	pow3     [MaxLevel + 3]int
	halfPow3 [MaxLevel + 3]int
)

// String returns the zone code
func (z *Zone) String() string {
	return z.Code
}

// Level returns the level
func (z *Zone) Level() int {
	return z.Pos.z.level
}

// Encode encodes a lat/lon/level into a Zone
func Encode(lat, lon float64, level int) (_ *Zone, err error) {
	zoom, ok := zooms[level]
	if !ok {
		return nil, ErrLevelInvalid
	}

	pnt := NewLL(lat, lon).Point() // Point at lat/lon
	pos := pnt.Position(zoom)      // Tile position

	return &Zone{Code: pos.Code(), Pos: pos}, nil
}

// Decode decodes a string code into Lat/Lon coordinates
func Decode(code string) (LL, error) {
	pos, err := DecodePosition(code)
	if err != nil {
		return LL{}, err
	}
	return pos.LL(), nil
}

// DecodePosition decodes a string code into a Position,
// useful for further operations without having to decode it into a Lat/Lon, like calculating neighbours
func DecodePosition(code string) (*Position, error) {
	lnc := len(code)
	zoom, ok := zooms[lnc-2]
	if !ok {
		return nil, ErrCodeInvalid
	}

	var n1, n2 int
	if n1, ok = hIndex[code[0]]; !ok {
		return nil, ErrCodeInvalid
	} else if n2, ok = hIndex[code[1]]; !ok {
		return nil, ErrCodeInvalid
	}

	base := n1*30 + n2
	if base < 100 {
		code = "0" + strconv.Itoa(base) + code[2:]
	} else {
		code = strconv.Itoa(base) + code[2:]
	}

	pos := &Position{z: zoom}
	for i, digit := range code {
		n := int64(digit - '0')
		if n < 0 || n > 9 {
			return nil, fmt.Errorf("expected a digit, got '%b'", digit)
		}

		pow := pow3[lnc-i]
		c3x := n / 3
		c3y := n % 3
		switch c3x {
		case 0:
			pos.X -= pow
		case 2:
			pos.X += pow
		}
		switch c3y {
		case 0:
			pos.Y -= pow
		case 2:
			pos.Y += pow
		}
	}
	return pos, nil

}

func init() {
	for i := 0; i < MaxLevel+3; i++ {
		pow := math.Pow(3, float64(i))
		pow3[i] = int(pow)
		halfPow3[i] = int(math.Ceil(pow / 2))
	}
}
