package geohex

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("LL", func() {

	It("should create new LLs", func() {
		ll1 := NewLL(66.68, -87.98)
		Expect(ll1.Lat).To(Equal(66.68))
		Expect(ll1.Lon).To(Equal(-87.98))

		ll2 := NewLL(0.0, 370.5)
		Expect(ll2.Lat).To(Equal(0.0))
		Expect(ll2.Lon).To(Equal(10.5))
	})

	It("should create points", func() {
		pt := NewLL(66.68, -87.98).Point()
		Expect(pt).To(BeAssignableToTypeOf(Point{}))
	})

	It("should decode", func() {
		for _, tc := range testCasesCode2HEX {
			pos, err := Decode(tc.code)
			Expect(err).NotTo(HaveOccurred(), "for %#v", tc)

			act := pos.LL()
			Expect(act.Lat).To(BeNumerically("~", tc.exp.Lat, 0.000001), "for %#v", tc)
			Expect(act.Lon).To(BeNumerically("~", tc.exp.Lon, 0.000001), "for %#v", tc)
		}
	})

	It("should encode", func() {
		for _, tc := range testCasesCoord2HEX {
			pos, err := Encode(tc.ll.Lat, tc.ll.Lon, tc.level)
			Expect(err).NotTo(HaveOccurred(), "for %#v", tc)

			code := pos.Code()
			Expect(code).To(Equal(tc.exp), "for %#v", tc)
		}
	})

})
