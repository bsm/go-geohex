package geohex

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("LL", func() {

	It("should decode", func() {
		for _, tc := range testCasesCode2LL {
			tc := &tc
			pos, err := Decode(tc.code)
			Expect(err).NotTo(HaveOccurred(), "for %s", tc.code)

			act := pos.LL()
			Expect(act.Lat).To(BeNumerically("~", tc.exp.Lat, 0.000001), "for %s", tc.code)
			Expect(act.Lon).To(BeNumerically("~", tc.exp.Lon, 0.000001), "for %s", tc.code)
		}
	})

	It("should encode", func() {
		for _, tc := range testCasesLL2Code {
			pos, err := Encode(tc.ll.Lat, tc.ll.Lon, tc.level)
			Expect(err).NotTo(HaveOccurred(), "for %v %d", tc.ll, tc.level)

			code := pos.Code()
			Expect(code).To(Equal(tc.exp), "for %v %d", tc.ll, tc.level)
		}
	})

})
