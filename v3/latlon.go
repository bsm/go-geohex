package geohex

import (
	"fmt"
	"math"
)

// LL is a lat/lon tuple
type LL struct{ Lat, Lon float64 }

// NewLL creates a new normalised LL
func NewLL(lat, lon float64) LL {
	if lon < -180 {
		lon += 360
	} else if lon > 180 {
		lon -= 360
	}
	return LL{Lat: lat, Lon: lon}
}

// Point generates a grid point from a lat/lon
func (ll LL) Point() Point {
	e := ll.Lon * hBase / 180.0
	n := math.Log(math.Tan((90+ll.Lat)*hD2R/2)) / math.Pi * hBase

	return Point{E: e, N: n}
}

// Position encodes the position from a lat/lon
func (ll LL) Position(level int) (Position, error) {
	return ll.Point().Position(level)
}

// String returns a string representation of this coordinates
func (ll LL) String() string {
	return fmt.Sprintf("[%f, %f]", ll.Lat, ll.Lon)
}
