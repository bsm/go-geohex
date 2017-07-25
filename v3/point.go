package geohex

import (
	"math"
)

// Point implements geographic Cartesian coordinates of a Mercator projection to 1x1 square
type Point struct {
	E, N float64
}

// Position returns the X/Y grid position of the Point
func (p Point) Position(level int) (Position, error) {
	if level < 0 || level > MaxLevel {
		return Position{}, ErrLevelInvalid
	}
	z := zooms[level]

	x := (p.N/hK + p.E) * float64(z.size)
	y := (p.N/hK - p.E) * float64(z.size)

	x0, y0 := math.Floor(x), math.Floor(y)
	xd, yd := x-x0, y-y0

	pos := Position{Level: level}
	if yd > -xd+1 && yd < 2*xd && yd > 0.5*xd {
		pos.X, pos.Y = int(x0)+1, int(y0)+1
	} else if yd <= -xd+1 && yd > 2*xd-1 && yd < 0.5*xd+0.5 {
		pos.X, pos.Y = int(x0), int(y0)
	} else if yd > xd {
		pos.X, pos.Y = int(x0), int(y0)+1
	} else {
		pos.X, pos.Y = int(x0)+1, int(y0)
	}

	// pos.X-pos.Y == z.size means that we've wrapped through the eastmost border
	// For example, there's no x=0,y=-9 position on level 0, that's x=-9,y=0
	if pos.X-pos.Y == z.size {
		pos.X, pos.Y = pos.Y, pos.X
	}

	return pos, nil
}
