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

    BenchmarkEncodeLevel2-4      5000000         327 ns/op        88 B/op        5 allocs/op
    BenchmarkEncodeLevel6-4      5000000         356 ns/op        96 B/op        5 allocs/op
    BenchmarkEncodeLevel15-4     3000000         428 ns/op       144 B/op        5 allocs/op
    BenchmarkDecodeLevel2-4      5000000         298 ns/op        19 B/op        2 allocs/op
    BenchmarkDecodeLevel6-4      5000000         313 ns/op        19 B/op        2 allocs/op
    BenchmarkDecodeLevel15-4     5000000         347 ns/op        19 B/op        2 allocs/op
