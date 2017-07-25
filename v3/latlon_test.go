package geohex

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("LL", func() {

	It("should create new LLs", func() {
		ll1 := newLL(66.68, -87.98)
		Expect(ll1.Lat).To(Equal(66.68))
		Expect(ll1.Lon).To(Equal(-87.98))

		ll2 := newLL(0.0, 370.5)
		Expect(ll2.Lat).To(Equal(0.0))
		Expect(ll2.Lon).To(Equal(10.5))
	})

})

var _ = Describe("Point to tile", func() {
	for _, tc := range loadLL2PositionTestCases() {
		tc := tc
		It("should create tile from "+tc.ll.String(), func() {
			pos, _ := tc.ll.tile(tc.level)
			Expect(pos.X).To(Equal(tc.expectedX))
			Expect(pos.Y).To(Equal(tc.expectedY))
		})
	}
})
