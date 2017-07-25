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

})
