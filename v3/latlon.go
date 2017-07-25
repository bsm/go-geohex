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

	// x, y are coordinates over the tiles, but we have to check which tile they belong to
	x := (n/hK + e) * float64(size)
	y := (n/hK - e) * float64(size)

	x0, y0 := math.Floor(x), math.Floor(y)
	xd, yd := x-x0, y-y0

	tile := Tile{Level: level}
	if yd > -xd+1 && yd < 2*xd && yd > 0.5*xd {
		tile.X, tile.Y = int(x0)+1, int(y0)+1
	} else if yd <= -xd+1 && yd > 2*xd-1 && yd < 0.5*xd+0.5 {
		tile.X, tile.Y = int(x0), int(y0)
	} else if yd > xd {
		tile.X, tile.Y = int(x0), int(y0)+1
	} else {
		tile.X, tile.Y = int(x0)+1, int(y0)
	}

	// tile.X-tile.Y == size means that we've wrapped through the eastmost border
	// For example, there's no x=0,y=-9 tile on level 0, that's x=-9,y=0
	if tile.X-tile.Y == size {
		tile.X, tile.Y = tile.Y, tile.X
	}

	return tile, nil
}

// String returns a string representation of this coordinates
func (ll LL) String() string {
	return fmt.Sprintf("[%f, %f]", ll.Lat, ll.Lon)
}
