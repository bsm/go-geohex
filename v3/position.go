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

// Decode decodes a string code into a Position,
// useful for further operations without having to decode it into a Lat/Lon, like calculating neighbours
func Decode(code string) (Position, error) {
	lnc := len(code)
	pos := Position{Level: lnc - 2}

	_, ok := zooms[lnc-2]
	if !ok {
		return pos, ErrCodeInvalid
	}

	var n1, n2 int
	if n1, ok = hIndex[code[0]]; !ok {
		return pos, ErrCodeInvalid
	} else if n2, ok = hIndex[code[1]]; !ok {
		return pos, ErrCodeInvalid
	}

	base := n1*30 + n2
	if base < 100 {
		code = "0" + strconv.Itoa(base) + code[2:]
	} else {
		code = strconv.Itoa(base) + code[2:]
	}

	for i, digit := range code {
		n := int64(digit - '0')
		if n < 0 || n > 9 {
			return pos, fmt.Errorf("expected a digit, got '%b'", digit)
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

// Encode encodes a lat/lon/level into a Position
func Encode(lat, lon float64, level int) (Position, error) {
	return NewLL(lat, lon).Position(level)
}

// Centroid returns the centroid point of the tile
func (p Position) Centroid() Point {
	z := zooms[p.Level]
	x := float64(p.X)
	y := float64(p.Y)
	n := (hK*x*z.w + y*z.h) / 2
	e := (n - y*z.h) / hK
	return Point{E: e, N: n}
}

// LL converts the position into a LL
func (p Position) LL() LL {
	c := p.Centroid()
	z := zooms[p.Level]
	lat := 180 / math.Pi * (2*math.Atan(math.Exp(c.N/hBase*180*hD2R)) - math.Pi/2)

	var lon float64
	if math.Abs(-hBase-c.E) <= z.size/2 {
		lon = -180
	} else {
		lon = c.E / hBase * 180
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

	ll := p.LL()
	// Magic time. Unoptimized so far.
	if ll.Lon == -180 || ll.Lon >= 0 {
		if bx[1] == by[1] && bx[2] == by[2] {
			if bx[0] == 2 && by[0] == 1 {
				bx[0], by[0] = 1, 2
			} else if bx[0] == 1 && by[0] == 0 {
				bx[0], by[0] = 0, 1
			}
		}
	}

	base = 3*(100*bx[0]+10*bx[1]+bx[2]) + (100*by[0] + 10*by[1] + by[2])

	code[0] = hChars[base/30]
	code[1] = hChars[base%30]

	return string(code)
}

// String returns a String representation of this position (without taking in account zoom level)
func (p Position) String() string {
	return fmt.Sprintf("[%d, %d]@%d", p.X, p.Y, p.Level)
}
