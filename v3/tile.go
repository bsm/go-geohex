package geohex

import (
	"fmt"
	"math"
	"strconv"
)

// Tile implements a grid tile tile
type Tile struct {
	X, Y, Level int
}

// LL converts the tile into a LL
func (t Tile) LL() LL {
	if t.Level < 0 || t.Level > MaxLevel {
		return LL{}
	}
	size := sizes[t.Level]

	hX := float64(t.X) / 2
	hY := float64(t.Y) / 2

	e := (hX - hY) / float64(size)
	n := (hX + hY) / float64(size) * hK

	lat := (2*math.Atan(math.Exp(360*n*deg2Rad)) - pio2) / deg2Rad

	var lon float64
	// t.Y - t.X == size means that we are on the westmost border.
	// We have some precision errors here and we often calculate those as -179.99999... , so just fix that
	if t.Y-t.X == size {
		lon = -180
	} else {
		lon = 360 * e
	}

	return newLL(lat, lon)
}

// Code returns string Code of this tile
func (t Tile) Code() string {
	x, y := t.X, t.Y
	bx, by, base := make([]int, 3), make([]int, 3), 0
	c3x, c3y := 0, 0
	code := make([]byte, t.Level+2)

	for i := 0; i < t.Level+3; i++ {
		pow := pow3[t.Level+2-i]
		p2c := halfPow3[t.Level+2-i]

		if x >= p2c {
			x -= pow
			c3x = 2
		} else if x <= -p2c {
			x += pow
			c3x = 0
		} else {
			c3x = 1
		}

		if y >= p2c {
			y -= pow
			c3y = 2
		} else if y <= -p2c {
			y += pow
			c3y = 0
		} else {
			c3y = 1
		}

		if i >= 3 {
			code[i-1] = '0' + byte(3*c3x+c3y)
		} else {
			bx[i] = c3x
			by[i] = c3y
		}
	}

	// y <= x is the same condition as
	//  centroid.Lon >= 0 || centroid.Lon == -180
	// in the original library code. Without this, we'd generate other first two characters.
	// They would still be correctly decoded to the same tile though
	if y <= x {
		if bx[1] == by[1] && bx[2] == by[2] {
			if (bx[0] == 2 && by[0] == 1) || (bx[0] == 1 && by[0] == 0) {
				bx[0], by[0] = by[0], bx[0]
			}
		}
	}

	base = 3*(100*bx[0]+10*bx[1]+bx[2]) + (100*by[0] + 10*by[1] + by[2])

	code[0] = hChars[base/30]
	code[1] = hChars[base%30]

	return string(code)
}

// DecodeTile decodes a string code into a Tile,
// useful for further operations without having to decode it into a Lat/Lon, like calculating neighbours
func DecodeTile(code string) (Tile, error) {
	lnc := len(code)
	level := lnc - 2

	if level < 0 || level > MaxLevel {
		return Tile{}, ErrCodeInvalid
	}

	var (
		n1, n2 int
		ok     bool
	)

	if n1, ok = hIndex[code[0]]; !ok {
		return Tile{}, ErrCodeInvalid
	} else if n2, ok = hIndex[code[1]]; !ok {
		return Tile{}, ErrCodeInvalid
	}

	base := n1*30 + n2
	if base < 100 {
		code = "0" + strconv.Itoa(base) + code[2:]
	} else {
		code = strconv.Itoa(base) + code[2:]
	}

	var x, y int
	for i, digit := range code {
		n := int64(digit - '0')
		if n < 0 || n > 9 {
			return Tile{}, ErrCodeInvalid
		}

		pow := pow3[lnc-i]
		c3x := n / 3
		c3y := n % 3
		switch c3x {
		case 0:
			x -= pow
		case 2:
			x += pow
		}
		switch c3y {
		case 0:
			y -= pow
		case 2:
			y += pow
		}
	}

	overflow := y - x - sizes[level]
	if overflow > 0 {
		if x > y {
			x, y = y+overflow, x-overflow
		} else if x < y {
			x, y = y-overflow, x+overflow
		}
	}

	return Tile{X: x, Y: y, Level: level}, nil
}

// String returns a String representation of this tile
func (t Tile) String() string {
	return fmt.Sprintf("[%d, %d]@%d", t.X, t.Y, t.Level)
}
