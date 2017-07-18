# GeoHex

[![Build Status](https://travis-ci.org/bsm/geohex.go.png)](https://travis-ci.org/bsm/geohex.go)

GeoHex implementation in Go

## Quick Start

    import (
      geohex "github.com/bsm/geohex/v3"
    )

    func main() {
      geohex.Encode(35.647401, 139.716911, 6)
      // "XM488541"

      geohex.Decode("XM488541")
      // &LL{Lat: 35.63992106909, Lon: 139.7256515775}
    }

## Running tests

You need to install Ginkgo & Gomega to run tests. Please see
http://onsi.github.io/ginkgo/ for more details.

    $ make testdeps

To run tests, call:

    $ make test

To run benchmarks, call:

    $ make bench

## Latest benchmarks

    BenchmarkEncode-4        5000000           387 ns/op         109 B/op          5 allocs/op
    BenchmarkDecode-4        5000000           302 ns/op          19 B/op          2 allocs/op
