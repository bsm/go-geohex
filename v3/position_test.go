package geohex

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Position", func() {

	It("should decode", func() {
		for _, tc := range testCasesCode2XY {
			Expect(Decode(tc.code)).To(Equal(tc.exp), "for %#v", tc)
		}
	})

	It("should generate centroids", func() {
		cnt := Position{X: 4, Y: -5}.Centroid()
		Expect(cnt.E).To(BeNumerically("~", 20037508.3, 0.1))
		Expect(cnt.N).To(BeNumerically("~", -1285406.8, 0.1))
	})

	It("should encode", func() {
		for _, tc := range testCasesXY2HEX {
			code := Position{X: tc.x, Y: tc.y, Level: tc.level}.Code()
			Expect(code).To(Equal(tc.exp), "for %#v", tc)
		}
	})

})
