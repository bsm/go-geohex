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
	pow3f [MaxLevel + 3]float64
	pow3i [MaxLevel + 3]int
)

// String returns the zone code
func (z *Zone) String() string {
	return z.Code
}

// Level returns the level
func (z *Zone) Level() int {
	return len(z.Code) - 2
}

// Encode encodes a lat/lon/level into a Zone
func Encode(lat, lon float64, level int) (_ *Zone, err error) {
	zoom, ok := zooms[level]
	if !ok {
		return nil, ErrLevelInvalid
	}

	pnt := NewLL(lat, lon).Point() // Point at lat/lon
	pos := pnt.Position(zoom)      // Tile position
	cnt := pos.Centroid()          // Centroid of pos

	x, y := float64(pos.X), float64(pos.Y)
	if hBase-cnt.E < zoom.size {
		x, y = y, x
	}
	base, code := 0, make([]byte, level+2)

	for i := 0; i < level+3; i++ {
		pow := pow3f[level+2-i]
		p2c := math.Ceil(pow / 2)
		c3x, c3y := 1, 1

		if x >= p2c {
			x -= pow
			c3x = 2
		} else if x <= -p2c {
			x += pow
			c3x = 0
		}

		if y >= p2c {
			y -= pow
			c3y = 2
		} else if y <= -p2c {
			y += pow
			c3y = 0
		}

		num := c3x*3 + c3y
		if i == 0 {
			base += 100 * num
		} else if i == 1 {
			base += 10 * num
		} else if i == 2 {
			base += num
		} else {
			code[i-1] = '0' + byte(num)
		}
	}

	code[0] = hChars[base/30]
	code[1] = hChars[base%30]

	return &Zone{Code: string(code), Pos: pos}, nil
}

// Decode decodes a string code into Point
func Decode(code string) (_ *LL, err error) {
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

	pos := &Position{z: zoom}
	base := n1*30 + n2
	if base < 100 {
		code = "0" + strconv.Itoa(base) + code[2:]
	} else {
		code = strconv.Itoa(base) + code[2:]
	}
	for i, digit := range code {
		n := int64(digit - '0')
		if n < 0 || n > 9 {
			err = fmt.Errorf("expected a digit, got '%s'", digit)
			return
		}

		pow := pow3i[lnc-i]
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

	return pos.LL(), nil
}

func init() {
	for i := 0; i < MaxLevel+3; i++ {
		pow3f[i] = math.Pow(3, float64(i))
		pow3i[i] = int(math.Pow(3, float64(i)))
	}
}
