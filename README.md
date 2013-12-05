# GeoHex

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

    $ go test

## Licence

GeoHex is originally licenced by @sa2da under a Creative Commons BY-SA 2.1
License.

    This code is covered by MIT License.
    (The MIT License)

    Copyright (c) 2013 Black Square Media Ltd

    Permission is hereby granted, free of charge, to any person obtaining
    a copy of this software and associated documentation files (the
    'Software'), to deal in the Software without restriction, including
    without limitation the rights to use, copy, modify, merge, publish,
    distribute, sublicense, and/or sell copies of the Software, and to
    permit persons to whom the Software is furnished to do so, subject to
    the following conditions:

    The above copyright notice and this permission notice shall be
    included in all copies or substantial portions of the Software.

    THE SOFTWARE IS PROVIDED 'AS IS', WITHOUT WARRANTY OF ANY KIND,
    EXPRESS OR IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF
    MERCHANTABILITY, FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT.
    IN NO EVENT SHALL THE AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY
    CLAIM, DAMAGES OR OTHER LIABILITY, WHETHER IN AN ACTION OF CONTRACT,
    TORT OR OTHERWISE, ARISING FROM, OUT OF OR IN CONNECTION WITH THE
    SOFTWARE OR THE USE OR OTHER DEALINGS IN THE SOFTWARE.
