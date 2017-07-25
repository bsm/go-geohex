package geohex

import (
	"fmt"
	"math"
)

// LL is a lat/lon tuple
type LL struct {
	Lat, Lon float64
}

// newLL creates a new normalised LL
func newLL(lat, lon float64) LL {
	if lon < -180 {
		lon += 360
	} else if lon >= 180 {
		lon -= 360
	}
	return LL{Lat: lat, Lon: lon}
}

// point generates a grid point from a lat/lon
func (ll LL) tile(level int) (Tile, error) {
	if level < 0 || level > MaxLevel {
		return Tile{}, ErrLevelInvalid
	}
	size := sizes[level]

	// First of all, calculate coordinates of the projection
	// e, n are coordinates of Mercator projection of lat/lon to a 1x1 square
	e := ll.Lon / 360.0
	n := math.Log(math.Tan((ll.Lat*deg2Rad+pio2)/2)) / math.Pi / 2

	// fX, fY are float coordinates over the tiles, but we have to check which tile they belong to
	fX := (n/hK + e) * float64(size)
	fY := (n/hK - e) * float64(size)

	x0, y0 := math.Floor(fX), math.Floor(fY)
	xd, yd := fX-x0, fY-y0

	var x, y int
	if yd > -xd+1 && yd < 2*xd && yd > 0.5*xd {
		x, y = int(x0)+1, int(y0)+1
	} else if yd <= -xd+1 && yd > 2*xd-1 && yd < 0.5*xd+0.5 {
		x, y = int(x0), int(y0)
	} else if yd > xd {
		x, y = int(x0), int(y0)+1
	} else {
		x, y = int(x0)+1, int(y0)
	}

	return NewTile(x, y, level), nil
}

// String returns a string representation of this coordinates
func (ll LL) String() string {
	return fmt.Sprintf("[%f, %f]", ll.Lat, ll.Lon)
}
