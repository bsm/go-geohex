package geohex

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("Encode", func() {
	var zone *Zone
	var err error

	for _, tc := range loadLL2HexTestCases() {
		tc := tc
		It("should encode "+tc.ll.String()+" to "+tc.expectedCode, func() {
			zone, err = Encode(tc.ll.Lat, tc.ll.Lon, tc.level)
			Expect(err).To(BeNil())
			Expect(zone.Code).To(Equal(tc.expectedCode))
		})
	}

	It("should wrap the position and level", func() {
		zone, err := Encode(85.04354094565655, 89.2529296875, 3)
		Expect(err).To(BeNil())
		Expect(zone.Level()).To(Equal(3))
		Expect(zone.Pos.X).To(Equal(271))
		Expect(zone.Pos.Y).To(Equal(150))
	})

})

var _ = Describe("Decode", func() {

	for _, tc := range loadCode2HexTestCases() {
		tc := tc
		It("should decode "+tc.expectedLL.String()+" from "+tc.code, func() {
			act, err := Decode(tc.code)
			Expect(err).To(BeNil())

			Expect(act.Lat).To(BeNumerically("~", tc.expectedLL.Lat))
			Expect(act.Lon).To(BeNumerically("~", tc.expectedLL.Lon))
		})
	}

})
