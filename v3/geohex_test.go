package geohex

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"testing"
)

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "geohex")
}

func BenchmarkEncode(b *testing.B) {
	for i := 0; i < b.N; i += 3 {
		Encode(7.092954137951794, 179.9073230957031, 2)
		Encode(-21.616579336740593, -166.9921875, 6)
		Encode(82.07002819448266, -177.890625, 15)
	}
}

func BenchmarkDecode(b *testing.B) {
	for i := 0; i < b.N; i += 3 {
		Decode("QU08")
		Decode("GH501658")
		Decode("TK720137660817775")
	}
}

var _ = Describe("Zoom", func() {
	var subject *Zoom

	It("should preload zooms", func() {
		Expect(zooms).To(HaveLen(21))
	})

	It("should calculate attributes", func() {
		subject = zooms[7]
		Expect(subject.level).To(Equal(7))
		Expect(subject.size).To(BeNumerically("~", 339.337, 0.001))
		Expect(subject.scale).To(BeNumerically("~", 0.000053, 0.000001))
		Expect(subject.w).To(BeNumerically("~", 2036.022, 0.001))
		Expect(subject.h).To(BeNumerically("~", 1175.498, 0.001))
	})
})

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
		Expect(pt).To(BeAssignableToTypeOf(&Point{}))
	})

})
