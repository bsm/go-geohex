package geohex

import (
	"fmt"
	"math"
)

// Position implements a grid tile position
type Position struct {
	X, Y int
	z    *Zoom

	code string
}

// Centroid returns the centroid point of the tile
func (p *Position) Centroid() *Point {
	x := float64(p.X)
	y := float64(p.Y)
	n := (hK*x*p.z.w + y*p.z.h) / 2
	e := (n - y*p.z.h) / hK
	return &Point{E: e, N: n}
}

// LL converts the position into a LL
func (p *Position) LL() *LL {
	c := p.Centroid()
	lat := 180 / math.Pi * (2*math.Atan(math.Exp(c.N/hBase*180*hD2R)) - math.Pi/2)

	var lon float64
	if math.Abs(-hBase-c.E) <= p.z.size/2 {
		lon = -180
	} else {
		lon = c.E / hBase * 180
	}

	return NewLL(lat, lon)
}

// Code returns string Code of this position
func (p *Position) Code() string {
	if p.code != "" {
		return p.code
	}

	x, y := float64(p.X), float64(p.Y)
	base, num, code := 0, 0, make([]byte, p.z.level+2)

	for i := 0; i < p.z.level+3; i++ {
		pow := pow3f[p.z.level+2-i]
		p2c := halfCeilPow3f[p.z.level+2-i]

		if x >= p2c {
			x -= pow
			num = 6
		} else if x <= -p2c {
			x += pow
			num = 0
		} else {
			num = 3
		}

		if y >= p2c {
			y -= pow
			num += 2
		} else if y <= -p2c {
			y += pow
			// num += 0
		} else {
			num += 1
		}

		if i >= 3 {
			code[i-1] = '0' + byte(num)
		} else if i == 2 {
			base += num
		} else if i == 1 {
			base += 10 * num
		} else {
			base += 100 * num
		}
	}

	code[0] = hChars[base/30]
	code[1] = hChars[base%30]

	p.code = string(code)

	return p.code
}

// String returns a String representation of this position (without taking in account zoom level)
func (p *Position) String() string {
	level := -1
	if p.z != nil {
		level = p.z.level
	}
	return fmt.Sprintf("[%d, %d]@d", p.X, p.Y, level)
}
