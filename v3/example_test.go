package geohex_test

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
