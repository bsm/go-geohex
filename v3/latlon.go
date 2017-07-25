package geohex

import (
	"fmt"
	"math"
)

// LL is a lat/lon tuple
type LL struct {
	Lat, Lon float64
}

// NewLL creates a new normalised LL
func NewLL(lat, lon float64) LL {
	if lon < -180 {
		lon += 360
	} else if lon >= 180 {
		lon -= 360
	}
	return LL{Lat: lat, Lon: lon}
}

// Point generates a grid point from a lat/lon
func (ll LL) Point() Point {
	return Point{
		E: ll.Lon / 360.0,
		N: math.Log(math.Tan((ll.Lat*deg2Rad+pio2)/2)) / math.Pi / 2,
	}
}

// String returns a string representation of this coordinates
func (ll LL) String() string {
	return fmt.Sprintf("[%f, %f]", ll.Lat, ll.Lon)
}
