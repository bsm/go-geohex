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
