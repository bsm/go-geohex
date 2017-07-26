package geohex

import (
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Position", func() {

	It("should decode", func() {
		for _, tc := range testCasesCode2XY {
			Expect(Decode(tc.code)).To(Equal(tc.exp), "for %#v", tc)
		}
	})

	It("should encode", func() {
		for _, tc := range testCasesXY2HEX {
			code := Position{X: tc.x, Y: tc.y, Level: tc.level}.Code()
			Expect(code).To(Equal(tc.exp), "for %#v", tc)
		}
	})

})

var _ = DescribeTable("Should return neighbors of Positions with correct coordinates",

	func(Position Position, expectedNeighbours []interface{}) {
		Expect(Position.Neighbours()).To(ConsistOf(expectedNeighbours...))
	},

	Entry("OY (0, 0 - center)",
		Position{X: 0, Y: 0, Level: 0},
		[]interface{}{
			Position{X: 1, Y: 1, Level: 0},   // Oc
			Position{X: 1, Y: 0, Level: 0},   // Ob
			Position{X: 0, Y: -1, Level: 0},  // OX
			Position{X: -1, Y: -1, Level: 0}, // OU
			Position{X: -1, Y: 0, Level: 0},  // OV
			Position{X: 0, Y: 1, Level: 0},   // OZ
		},
	),

	Entry("XL (5,-3 - wrapping over lon 180º)",
		Position{X: 5, Y: -3, Level: 0},
		[]interface{}{
			Position{X: 6, Y: -2, Level: 0}, // XP
			Position{X: -3, Y: 6, Level: 0}, // QY
			Position{X: -4, Y: 5, Level: 0}, // QU
			Position{X: 4, Y: -4, Level: 0}, // PQ
			Position{X: 4, Y: -3, Level: 0}, // PR
			Position{X: 5, Y: -2, Level: 0}, // XM
		},
	),
)
