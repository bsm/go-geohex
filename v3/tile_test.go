package geohex

import (
	"fmt"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/ginkgo/extensions/table"
	. "github.com/onsi/gomega"
)

var _ = Describe("Tile", func() {

	for _, tc := range loadPosition2HexTestCases() {
		tc := tc
		It(fmt.Sprintf("should encode tile [%d, %d] to %s", tc.x, tc.y, tc.expectedCode), func() {
			pos := Tile{X: tc.x, Y: tc.y, Level: tc.level}
			code := pos.Code()
			Expect(code).To(Equal(tc.expectedCode))
		})
	}

})

var _ = Describe("Decode Tile from Code", func() {

	for _, tc := range loadCode2PositionTestCases() {
		tc := tc
		It(fmt.Sprintf("should decode tile [%d, %d] from %s", tc.expectedPosition.X, tc.expectedPosition.Y, tc.code), func() {
			act, err := DecodeTile(tc.code)
			Expect(err).To(BeNil())

			Expect(act.X).To(Equal(tc.expectedPosition.X))
			Expect(act.Y).To(Equal(tc.expectedPosition.Y))
		})
	}

})

var _ = DescribeTable("Should return neighbors of Tiles with correct codes",
	func(tileCode string, expectedNeighbours []interface{}) {
		tile, _ := DecodeTile(tileCode)
		neighbourCodes := make([]string, 6)
		for i, n := range tile.Neighbours() {
			neighbourCodes[i] = n.Code()
		}
		Expect(neighbourCodes).To(ConsistOf(expectedNeighbours...))
	},
	Entry("OY (center)", "OY", []interface{}{"Oc", "Ob", "OX", "OU", "OV", "OZ"}),
	Entry("XL (wrapping)", "XL", []interface{}{"XP", "QY", "QU", "PQ", "PR", "XM"}),
)

var _ = DescribeTable("Should return neighbors of Tiles with correct coordinates",
	func(tile Tile, expectedNeighbours []interface{}) {
		Expect(tile.Neighbours()).To(ConsistOf(expectedNeighbours...))
	},
	Entry("OY (0, 0 - center)",
		Tile{X: 0, Y: 0, Level: 0},
		[]interface{}{
			Tile{X: 1, Y: 1, Level: 0},   // Oc
			Tile{X: 1, Y: 0, Level: 0},   // Ob
			Tile{X: 0, Y: -1, Level: 0},  // OX
			Tile{X: -1, Y: -1, Level: 0}, // OU
			Tile{X: -1, Y: 0, Level: 0},  // OV
			Tile{X: 0, Y: 1, Level: 0},   // OZ
		},
	),
	Entry("XL (5,-3 - wrapping)",
		Tile{X: 5, Y: -3, Level: 0},
		[]interface{}{
			Tile{X: 6, Y: -2, Level: 0}, // XP
			Tile{X: -3, Y: 6, Level: 0}, // QY
			Tile{X: -4, Y: 5, Level: 0}, // QU
			Tile{X: 4, Y: -4, Level: 0}, // PQ
			Tile{X: 4, Y: -3, Level: 0}, // PR
			Tile{X: 5, Y: -2, Level: 0}, // XM
		},
	),
)
