package geohex

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Position", func() {

	It("should decode", func() {
		for _, tc := range testCasesCode2Pos {
			tc := &tc
			Expect(Decode(tc.code)).To(Equal(tc.exp), "%v from %s", tc.exp, tc.code)
		}
	})

	It("should generate centroids", func() {
		pos := &Position{X: 4, Y: -5, z: zooms[0]}
		cnt := pos.Centroid()
		Expect(cnt.E).To(BeNumerically("~", 20037508.3, 0.1))
		Expect(cnt.N).To(BeNumerically("~", -1285406.8, 0.1))
	})

	It("should encode", func() {
		for _, tc := range testCasesPos2Code {
			pos := Position{X: tc.x, Y: tc.y, z: zooms[tc.level]}
			code := pos.Code()
			Expect(code).To(Equal(tc.exp), "[%d, %d] to %s", tc.x, tc.y, tc.exp)
		}
	})

})
