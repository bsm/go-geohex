package geohex

import "math"

// Position implements a grid tile position
type Position struct {
	X, Y int
	z    *Zoom
}

// Centroid returns the centroidpoint of the tile
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
