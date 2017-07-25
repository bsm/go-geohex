package geohex

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tile", func() {

	for _, tc := range loadPosition2HexTestCases() {
		tc := tc
		It(fmt.Sprintf("should encode tile [%d, %d] to %s", tc.x, tc.y, tc.expectedCode), func() {
			pos := Tile{X: tc.x, Y: tc.y, Level: tc.level}
			code := pos.Code()
			Expect(code).To(Equal(tc.expectedCode))
		})
	}

})

var _ = Describe("Decode Tile from Code", func() {

	for _, tc := range loadCode2PositionTestCases() {
		tc := tc
		It(fmt.Sprintf("should decode tile [%d, %d] from %s", tc.expectedPosition.X, tc.expectedPosition.Y, tc.code), func() {
			act, err := DecodeTile(tc.code)
			Expect(err).To(BeNil())

			Expect(act.X).To(Equal(tc.expectedPosition.X))
			Expect(act.Y).To(Equal(tc.expectedPosition.Y))
		})
	}

})
