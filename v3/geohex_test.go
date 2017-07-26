package geohex

import (
	"math/rand"
	"testing"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Sizes", func() {

	It("should preload sizes", func() {
		Expect(sizes).To(HaveLen(21))
	})
})

func TestSuite(t *testing.T) {
	RegisterFailHandler(Fail)
	RunSpecs(t, "geohex")
}

// --------------------------------------------------------------------

func BenchmarkEncode(b *testing.B) {
	seeds := benchmarkSeedLL(300)

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		ll := seeds[i%len(seeds)]
		if _, err := Encode(ll.Lat, ll.Lon, 15); err != nil {
			b.Fatalf("expected no error but got: %v", err)
		}
	}
}

func BenchmarkDecode(b *testing.B) {
	seeds, err := benchmarkSeedPos(300)
	if err != nil {
		b.Fatalf("expected no error but got: %v", err)
	}

	codes := make([]string, 0)
	for _, pos := range seeds {
		codes = append(codes, pos.Code())
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		if _, err := Decode(codes[i%len(codes)]); err != nil {
			b.Fatalf("expected no error but got: %v", err)
		}
	}
}

func BenchmarkPosition_Code(b *testing.B) {
	seeds, err := benchmarkSeedPos(300)
	if err != nil {
		b.Fatalf("expected no error but got: %v", err)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		pos := seeds[i%len(seeds)]
		pos.Code()
	}
}

func benchmarkSeedLL(n int) []LL {
	rnd := rand.New(rand.NewSource(0))
	lls := make([]LL, 0)
	for i := 0; i < n; i++ {
		lls = append(lls, LL{
			Lat: rnd.Float64()*180 - 90,
			Lon: rnd.Float64()*360 - 180,
		})
	}
	return lls
}

func benchmarkSeedPos(n int) ([]Position, error) {
	res := make([]Position, 0)
	for _, ll := range benchmarkSeedLL(n) {
		pos, err := Encode(ll.Lat, ll.Lon, 15)
		if err != nil {
			return nil, err
		}
		res = append(res, pos)
	}
	return res, nil
}
