package geohex

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
	"math/rand"
	"testing"
)

const testItems = 10000

var (
	randomGenerator = rand.New(rand.NewSource(0))
	points          [testItems][2]float64
)

func init() {
	for i := 0; i < testItems; i++ {
		points[i] = [2]float64{randomGenerator.Float64()*180 - 90, randomGenerator.Float64()*360 - 180}
	}
}

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "geohex")
}

func BenchmarkEncodeLevel2(b *testing.B) {
	benchmarkEncode(b, 2)
}

func BenchmarkEncodeLevel6(b *testing.B) {
	benchmarkEncode(b, 6)
}

func BenchmarkEncodeLevel15(b *testing.B) {
	benchmarkEncode(b, 15)
}

func benchmarkEncode(b *testing.B, level int) {
	b.ResetTimer()
	for i := 0; i < b.N; i += 1 {
		p := points[i%testItems]
		Encode(p[0], p[1], level)
	}
}

func BenchmarkDecodeLevel2(b *testing.B) {
	benchmarkDecode(b, 2)
}

func BenchmarkDecodeLevel6(b *testing.B) {
	benchmarkDecode(b, 6)
}

func BenchmarkDecodeLevel15(b *testing.B) {
	benchmarkDecode(b, 15)
}

func benchmarkDecode(b *testing.B, level int) {
	codes := [testItems]string{}
	for i := 0; i < testItems; i++ {
		codes[i], _ = Encode(points[i][0], points[i][1], level)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i += 1 {
		Decode(codes[i%testItems])
	}
}

func BenchmarkDecodeTileLevel2(b *testing.B) {
	benchmarkDecodeTile(b, 2)
}

func BenchmarkDecodeTileLevel6(b *testing.B) {
	benchmarkDecodeTile(b, 6)
}

func BenchmarkDecodeTileLevel15(b *testing.B) {
	benchmarkDecodeTile(b, 15)
}

func benchmarkDecodeTile(b *testing.B, level int) {
	codes := [testItems]string{}
	for i := 0; i < testItems; i++ {
		codes[i], _ = Encode(points[i][0], points[i][1], level)
	}
	b.ResetTimer()
	for i := 0; i < b.N; i += 1 {
		DecodeTile(codes[i%testItems])
	}
}

var _ = Describe("Encode Code from LatLon", func() {
	for _, tc := range loadLL2CodeTestCases() {
		tc := tc
		It("should encode "+tc.ll.String()+" to "+tc.expectedCode, func() {
			code, err := Encode(tc.ll.Lat, tc.ll.Lon, tc.level)
			Expect(err).To(BeNil())
			Expect(code).To(Equal(tc.expectedCode))
		})
	}
})

var _ = Describe("Decode LatLon from Code", func() {

	for _, tc := range loadCode2LLTestCases() {
		tc := tc
		It("should decode latlon "+tc.expectedLL.String()+" from "+tc.code, func() {
			act, err := Decode(tc.code)
			Expect(err).To(BeNil())

			Expect(act.Lat).To(BeNumerically("~", tc.expectedLL.Lat))
			Expect(act.Lon).To(BeNumerically("~", tc.expectedLL.Lon))
		})
	}

})

var _ = Describe("Valid levels", func() {

	It("Should encode with level 0", func() {
		_, err := Encode(0, 0, 0)
		Expect(err).To(BeNil())
	})

	It("Should encode with level MaxLevel", func() {
		_, err := Encode(0, 0, MaxLevel)
		Expect(err).To(BeNil())
	})

})
