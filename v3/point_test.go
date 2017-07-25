package geohex

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Point", func() {
	var p1, p2 Point

	BeforeEach(func() {
		p1 = NewLL(-2.7315738409448347, 178.9405262207031).Point()
		p2 = NewLL(82.27244849463305, 172.87607309570308).Point()
	})

	It("should create points", func() {
		Expect(p1.E).To(BeNumerically("~", 19919568.3, 0.1))
		Expect(p1.N).To(BeNumerically("~", -304192.7, 0.1))
		Expect(p2.E).To(BeNumerically("~", 19244476.4, 0.1))
		Expect(p2.N).To(BeNumerically("~", 17189491.4, 0.1))
	})

	It("should create positions", func() {
		for _, tc := range testCasesLL2Pos {
			tc := tc
			pos := tc.ll.Point().Position(zooms[tc.level])
			Expect(pos.X).To(Equal(tc.expX), "for %s", tc.ll)
			Expect(pos.Y).To(Equal(tc.expY), "for %s", tc.ll)
		}
	})

})
