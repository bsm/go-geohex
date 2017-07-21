package geohex

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Position", func() {

	It("should generate centroids", func() {
		pos := &Position{X: 4, Y: -5, z: zooms[0]}
		cnt := pos.Centroid()
		Expect(cnt.E).To(BeNumerically("~", 20037508.3, 0.1))
		Expect(cnt.N).To(BeNumerically("~", -1285406.8, 0.1))
	})

	for _, tc := range loadPosition2HexTestCases() {
		tc := tc
		It(fmt.Sprintf("should encode position [%d, %d] to %s", tc.x, tc.y, tc.expectedCode), func() {
			pos := Position{X: tc.x, Y: tc.y, z: zooms[tc.level]}
			code := pos.Code()
			Expect(code).To(Equal(tc.expectedCode))
		})
	}

})
