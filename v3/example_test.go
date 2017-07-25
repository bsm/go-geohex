package geohex_test

import (
	"fmt"

	geohex "github.com/bsm/go-geohex/v3"
)

func ExampleEncode() {
	code, _ := geohex.Encode(35.647401, 139.716911, 6)
	fmt.Println(code)

	// Output:
	// XM488541
}

func ExampleDecode() {
	pos, _ := geohex.Decode("XM488541")
	ll := pos.LL()
	fmt.Println(ll.Lat, ll.Lon)

	// Output:
	// 35.63992106908978 139.7256515775034
}
