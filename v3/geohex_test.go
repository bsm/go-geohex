package geohex

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/rand"
	"testing"
)

const testItems = 30

var (
	randomGenerator = rand.New(rand.NewSource(0))
	points          [testItems][2]float64
	geohex2         [testItems]string
	geohex6         [testItems]string
	geohex15        [testItems]string
)

func init() {
	for i := 0; i < testItems; i++ {
		points[i] = [2]float64{randomGenerator.Float64()*180 - 90, randomGenerator.Float64()*360 - 180}
		zone2, _ := Encode(points[i][0], points[i][1], 2)
		geohex2[i] = zone2.String()
		zone6, _ := Encode(points[i][0], points[i][1], 6)
		geohex6[i] = zone6.String()
		zone15, _ := Encode(points[i][0], points[i][1], 15)
		geohex15[i] = zone15.String()
	}
}

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "geohex")
}

func BenchmarkEncodeLevel2(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		p := points[i%testItems]
		Encode(p[0], p[1], 2)
	}
}

func BenchmarkEncodeLevel6(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		p := points[i%testItems]
		Encode(p[0], p[1], 6)
	}
}

func BenchmarkEncodeLevel15(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		p := points[i%testItems]
		Encode(p[0], p[1], 15)
	}
}

func BenchmarkDecodeLevel2(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		Decode(geohex2[i%testItems])
	}
}

func BenchmarkDecodeLevel6(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		Decode(geohex6[i%testItems])
	}
}

func BenchmarkDecodeLevel15(b *testing.B) {
	for i := 0; i < b.N; i += 1 {
		Decode(geohex15[i%testItems])
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
		Expect(pt).To(BeAssignableToTypeOf(Point{}))
	})

})
