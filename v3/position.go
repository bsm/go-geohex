package geohex

import (
	"math"
	"strconv"
)

// Position implements a grid tile Position
type Position struct {
	X, Y  int
	Level uint8
}

// NewPosition instantiates a new normalized tile, but it does not check for level to be valid
func NewPosition(x, y int, level uint8) Position {
	// x-y  == size means that we've wrapped through the eastmost border
	// For example, there's no x=0,y=-9 tile on level 0, that's x=-9,y=0
	if x-y == sizes[level] {
		x, y = y, x
	}

	return Position{X: x, Y: y, Level: level}
}

// Decode decodes a string code into a Position.
func Decode(code string) (Position, error) {
	lnc := len(code)
	level := uint8(lnc - 2)

	size, ok := sizes[level]
	if !ok {
		return Position{}, ErrLevelInvalid
	}

	var n1, n2 int
	if n1, ok = hIndex[code[0]]; !ok {
		return Position{}, ErrCodeInvalid
	} else if n2, ok = hIndex[code[1]]; !ok {
		return Position{}, ErrCodeInvalid
	}

	base := n1*30 + n2
	if base < 100 {
		code = "0" + strconv.Itoa(base) + code[2:]
	} else {
		code = strconv.Itoa(base) + code[2:]
	}

	var x, y int
	for i, digit := range code {
		n := uint8(digit - '0')
		if n < 0 || n > 9 {
			return Position{}, ErrCodeInvalid
		}

		pow := pow3[lnc-i]
		switch n / 3 {
		case 0:
			x -= pow
		case 2:
			x += pow
		}
		switch n % 3 {
		case 0:
			y -= pow
		case 2:
			y += pow
		}
	}

	overflow := y - x - size
	if overflow > 0 {
		if x > y {
			x, y = y+overflow, x-overflow
		} else if x < y {
			x, y = y-overflow, x+overflow
		}
	}

	return Position{X: x, Y: y, Level: level}, nil
}

// Encode encodes a lat/lon/level into a Position
func Encode(lat, lon float64, level uint8) (Position, error) {
	return NewLL(lat, lon).Position(level)
}

// LL converts the Position into a LL
func (p Position) LL() LL {
	size, ok := sizes[p.Level]
	if !ok {
		return LL{}
	}

	// First of all, calculate coordinates of the projection
	// e, n are coordinates of Mercator projection of lat/lon to a [0,1]x[0,1] square
	hX := float64(p.X) / 2
	hY := float64(p.Y) / 2
	e := (hX - hY) / float64(size)
	n := (hX + hY) / float64(size) * hK

	lat := (2*math.Atan(math.Exp(360*n*deg2Rad)) - pio2) / deg2Rad

	var lon float64
	// p.Y - p.X == size means that we are on the westmost border.
	// We have some precision errors here and we often calculate those as -179.99999... , so just fix that
	if p.Y-p.X == size {
		lon = -180
	} else {
		lon = 360 * e
	}

	return NewLL(lat, lon)
}

// Code returns string Code of this Position
func (p Position) Code() string {
	x, y := p.X, p.Y

	var code [22]byte
	var bx, by [3]uint8
	var c3x, c3y uint8

	for i := uint8(0); i < p.Level+3; i++ {
		n := int(p.Level + 2 - i)
		pow := pow3[n]
		p2c := halfPow3[n]

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
			code[i-1] = '0' + 3*c3x + c3y
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

	base := 3*(100*int(bx[0])+10*int(bx[1])+int(bx[2])) + (100*int(by[0]) + 10*int(by[1]) + int(by[2]))
	code[0] = hChars[base/30]
	code[1] = hChars[base%30]
	return string(code[:p.Level+2])
}

// Neighbours returns the Positions of Hexagons that have at least one common side with one.
// It works correctly on the longitudes close to 180º/-180º, returning the normalized Positions,
// but it is not expected to work correctly close to poles, where the projection distorts everything anyways.
func (p Position) Neighbours() []Position {
	return []Position{
		NewPosition(p.X+1, p.Y+1, p.Level),
		NewPosition(p.X-1, p.Y-1, p.Level),
		NewPosition(p.X+1, p.Y, p.Level),
		NewPosition(p.X-1, p.Y, p.Level),
		NewPosition(p.X, p.Y+1, p.Level),
		NewPosition(p.X, p.Y-1, p.Level),
	}
}
