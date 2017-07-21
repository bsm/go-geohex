package geohex

import (
	"math"
)

// Point implements geographic Cartesian coordinates
type Point struct {
	E, N float64
}

// Position returns the X/Y grid position of the Point
func (p *Point) Position(z *Zoom) *Position {
	x, y := (p.E+p.N/hK)/z.w, (p.N-hK*p.E)/z.h
	x0, y0 := math.Floor(x), math.Floor(y)
	xd, yd := x-x0, y-y0

	pos := Position{z: z}
	if yd > -xd+1 && yd < 2*xd && yd > 0.5*xd {
		pos.X, pos.Y = int(x0)+1, int(y0)+1
	} else if yd <= -xd+1 && yd > 2*xd-1 && yd < 0.5*xd+0.5 {
		pos.X, pos.Y = int(x0), int(y0)
	} else if yd > xd {
		pos.X, pos.Y = int(x0), int(y0)+1
	} else {
		pos.X, pos.Y = int(x0)+1, int(y0)
	}

	// Not really efficient to do here, or at least should be cached
	cnt := pos.Centroid()
	if hBase-cnt.E <= z.size/2 {
		pos.X, pos.Y = pos.Y, pos.X
	}

	return &pos
}
