# GeoHex

[![Build Status](https://travis-ci.org/bsm/go-geohex.png)](https://travis-ci.org/bsm/go-geohex)
[![GoDoc](https://godoc.org/github.com/bsm/go-geohex?status.png)](http://godoc.org/github.com/bsm/go-geohex)
[![License](https://img.shields.io/badge/License-MIT-blue.svg)](https://opensource.org/licenses/MIT)

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

    BenchmarkEncodeLevel2-4   3000000  477 ns/op  120 B/op  6 allocs/op
    BenchmarkEncodeLevel6-4   3000000  515 ns/op  128 B/op  6 allocs/op
    BenchmarkEncodeLevel15-4  2000000  595 ns/op  176 B/op  6 allocs/op
    BenchmarkDecodeLevel2-4   5000000  374 ns/op   72 B/op  4 allocs/op
    BenchmarkDecodeLevel6-4   3000000  403 ns/op   80 B/op  4 allocs/op
    BenchmarkDecodeLevel15-4  3000000  455 ns/op   99 B/op  4 allocs/op
