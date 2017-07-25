package geohex

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Position", func() {

	It("should generate centroids", func() {
		pos := Position{X: 4, Y: -5, Level: 0}
		cnt := pos.Centroid()
		Expect(cnt.E).To(BeNumerically("~", 0.5, 0.0001))
		Expect(cnt.N).To(BeNumerically("~", -0.0321, 0.0001))
	})

	for _, tc := range loadPosition2HexTestCases() {
		tc := tc
		It(fmt.Sprintf("should encode position [%d, %d] to %s", tc.x, tc.y, tc.expectedCode), func() {
			pos := Position{X: tc.x, Y: tc.y, Level: tc.level}
			code := pos.Code()
			Expect(code).To(Equal(tc.expectedCode))
		})
	}

})

var _ = Describe("Decode Position from Code", func() {

	for _, tc := range loadCode2PositionTestCases() {
		tc := tc
		It(fmt.Sprintf("should decode position [%d, %d] from %s", tc.expectedPosition.X, tc.expectedPosition.Y, tc.code), func() {
			act, err := DecodePosition(tc.code)
			Expect(err).To(BeNil())

			Expect(act.X).To(Equal(tc.expectedPosition.X))
			Expect(act.Y).To(Equal(tc.expectedPosition.Y))
		})
	}

})
