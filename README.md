# GeoHex

[![Build Status](https://travis-ci.org/bsm/go-geohex.png)](https://travis-ci.org/bsm/go-geohex)
[![GoDoc](https://godoc.org/github.com/bsm/go-geohex?status.png)](http://godoc.org/github.com/bsm/go-geohex)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

GeoHex implementation in Go

## Quick Start

```go
import (
	"fmt"

	geohex "github.com/cabify/go-geohex/v3"
)

func ExampleEncode() {
	pos, _ := geohex.Encode(35.647401, 139.716911, 6)
	fmt.Println(pos.Code())

	// Output:
	// XM488541
}

func ExampleDecode() {
	pos, _ := geohex.Decode("XM488541")
	ll := pos.LL()
	fmt.Println(ll.Lat, ll.Lon)

	// Output:
	// 35.63992106908978 139.72565157750344
}

func ExampleNeighbours() {
	pos, _ := geohex.Decode("XM488541")
	for _, n := range pos.Neighbours() {
		fmt.Println(n.Code())
	}

	// Output:
	// XM488545
	// XM488516
	// XM488544
	// XM488517
	// XM488542
	// XM488540
}
```

## Running tests

You need to install Ginkgo & Gomega to run tests. Please see
http://onsi.github.io/ginkgo/ for more details.

    $ make testdeps

To run tests, call:

    $ make test

To run benchmarks, call:

    $ make bench

## Latest benchmarks

    BenchmarkEncode-4        10000000   191 ns/op   0 B/op   0 allocs/op
    BenchmarkDecode-4         3000000   478 ns/op   3 B/op   1 allocs/op
    BenchmarkPosition_Code-4  5000000   328 ns/op   32 B/op  1 allocs/op

## Encoding details

A quick explanation of how lat/lon coordinates are encoded to hexagon positions:

Lat/Lon coordinates are projected to a [0,1)x[0x1) square map using
[Mercator projection](https://en.wikipedia.org/wiki/Mercator_projection), note
that this means that hexagons at different latitudes cover different Earth areas.
The coordinates on this square are called `e`, `n` in the code (east, north).

Then we transform those coordinates into `x`, `y` coordinates on the hexagons map:
we turn the axes 45 degrees and we stretch one of them by a factor of tan(pi/6),
this makes each four adjacent hexagons centers' be equidistant. Then we decide
which hexagon contains the desired coordinates. This is probably the trickiest part,
we use this condition to decide:

```go
if yd > -xd+1 && yd < 2*xd && yd > 0.5*xd {
  x, y = int(x0)+1, int(y0)+1
} else if yd <= -xd+1 && yd > 2*xd-1 && yd < 0.5*xd+0.5 {
  x, y = int(x0), int(y0)
} else if yd > xd {
  x, y = int(x0), int(y0)+1
} else {
  x, y = int(x0)+1, int(y0)
}
```

Where `x0` and `y0` are the integer part of the `x`, `y` coordinates after changing
the base and `xd`, `yd` are the decimal part of those. If we assume that both `x0` and
`y0` are zero, you can see how those conditions define limit lines for each of
four hexagons on the next drawing, while our `x`, `y` point lays somewhere inside of the square:

![Hexagons on the 1x1 square](/doc/hexagons.png?raw=true)

This way, previous condtions mean:

```go
if yd > -xd+1 && yd < 2*xd && yd > 0.5*xd {
  // red hexagon
  x, y = int(x0)+1, int(y0)+1
} else if yd <= -xd+1 && yd > 2*xd-1 && yd < 0.5*xd+0.5 {
  // purple hexagon
  x, y = int(x0), int(y0)
} else if yd > xd {
  // green hexagon
  x, y = int(x0), int(y0)+1
} else {
  // blue hexagon
  x, y = int(x0)+1, int(y0)
}
```

Once we have our integer position of the hexagon, they are encoded into string,
approximating in powers of 3 level+2 times: each character is a two-digit base-3
number, where first digit means the relative position of previous `x` approximation,
and the second one defines the same for `y`.
First three approximations are encoded into two special characters, and after all
of that, some fixes are applied for world-wrapping and consistency: we maintain all
of them to make this implementation compatible with the original one from geohex.org.
