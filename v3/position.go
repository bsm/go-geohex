package geohex

import (
	"fmt"
	"math"
	"strconv"
)

// Position implements a grid tile position
type Position struct {
	X, Y, Level int
}

// Centroid returns the centroid point of the tile
func (p Position) Centroid() Point {
	z, err := getZoom(p.Level)
	if err != nil {
		return Point{}
	}

	hX := float64(p.X) / 2
	hY := float64(p.Y) / 2

	return Point{
		E: (hX - hY) / z.size,
		N: (hX + hY) / z.size * hK,
	}
}

// LL converts the position into a LL
func (p Position) LL() LL {
	z, err := getZoom(p.Level)
	if err != nil {
		return LL{}
	}

	c := p.Centroid()
	lat := (2*math.Atan(math.Exp(360*c.N/equatorLen*deg2Rad)) - pio2) / deg2Rad

	var lon float64
	// p.Y - p.X == z.wrap means that we are on the westmost border.
	// We have some precision errors here and we often calculate those as -179.99999... , so just fix that
	if p.Y-p.X == z.wrap {
		lon = -180
	} else {
		lon = 360 * c.E / equatorLen
	}

	return NewLL(lat, lon)
}

// Code returns string Code of this position
func (p Position) Code() string {
	x, y := p.X, p.Y
	bx, by, base := make([]int, 3), make([]int, 3), 0
	c3x, c3y := 0, 0
	code := make([]byte, p.Level+2)

	for i := 0; i < p.Level+3; i++ {
		pow := pow3[p.Level+2-i]
		p2c := halfPow3[p.Level+2-i]

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
	// They would still be correctly decoded to the same position though
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

// DecodePosition decodes a string code into a Position,
// useful for further operations without having to decode it into a Lat/Lon, like calculating neighbours
func DecodePosition(code string) (*Position, error) {
	lnc := len(code)
	level := lnc - 2
	_, ok := zooms[level]
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

	pos := &Position{Level: level}
	for i, digit := range code {
		n := int64(digit - '0')
		if n < 0 || n > 9 {
			return nil, ErrCodeInvalid
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

// String returns a String representation of this position
func (p Position) String() string {
	return fmt.Sprintf("[%d, %d]@%d", p.X, p.Y, p.Level)
}
