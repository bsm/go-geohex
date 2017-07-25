package geohex

import (
	"encoding/json"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// Test cases downloaded from http://geohex.net/testcase/v3.2.html

type (
	testCaseCode2HEX struct {
		code string
		exp  LL
	}
	testCaseCoord2HEX struct {
		level int
		ll    LL
		exp   string
	}
	testCaseCoord2XY struct {
		level int
		ll    LL
		exp   Position
	}
	testCaseCode2XY struct {
		code string
		exp  Position
	}
	testCaseXY2HEX struct {
		level int
		x     int
		y     int
		exp   string
	}
)

var (
	testCasesCode2HEX  []testCaseCode2HEX
	testCasesCoord2HEX []testCaseCoord2HEX
	testCasesCoord2XY  []testCaseCoord2XY
	testCasesCode2XY   []testCaseCode2XY
	testCasesXY2HEX    []testCaseXY2HEX
)

var _ = BeforeSuite(func() {

	// http://geohex.net/testcase/hex_v3.2_test_code2HEX.json
	loadTestCasesFromJson("hex_v3.2_test_code2HEX.json", func(raw []json.RawMessage) error {
		tc := testCaseCode2HEX{}
		err := unmarshalRawFields([]testCaseFieldMap{
			{raw[0], &tc.code},
			{raw[1], &tc.exp.Lat},
			{raw[2], &tc.exp.Lon},
		})
		testCasesCode2HEX = append(testCasesCode2HEX, tc)
		return err
	})

	// http://geohex.net/testcase/hex_v3.2_test_coord2HEX.json
	loadTestCasesFromJson("hex_v3.2_test_coord2HEX.json", func(raw []json.RawMessage) error {
		tc := testCaseCoord2HEX{}
		err := unmarshalRawFields([]testCaseFieldMap{
			{raw[0], &tc.level},
			{raw[1], &tc.ll.Lat},
			{raw[2], &tc.ll.Lon},
			{raw[3], &tc.exp},
		})
		testCasesCoord2HEX = append(testCasesCoord2HEX, tc)
		return err
	})

	// http://geohex.net/testcase/hex_v3.2_test_coord2XY.json
	loadTestCasesFromJson("hex_v3.2_test_coord2XY.json", func(raw []json.RawMessage) error {
		tc := testCaseCoord2XY{}
		err := unmarshalRawFields([]testCaseFieldMap{
			{raw[0], &tc.level},
			{raw[1], &tc.ll.Lat},
			{raw[2], &tc.ll.Lon},
			{raw[3], &tc.exp.X},
			{raw[4], &tc.exp.Y},
		})
		tc.exp.Level = tc.level
		testCasesCoord2XY = append(testCasesCoord2XY, tc)
		return err
	})

	// http://geohex.net/testcase/hex_v3.2_test_code2XY.json
	loadTestCasesFromJson("hex_v3.2_test_code2XY.json", func(raw []json.RawMessage) error {
		tc := testCaseCode2XY{}
		err := unmarshalRawFields([]testCaseFieldMap{
			{raw[0], &tc.code},
			{raw[1], &tc.exp.X},
			{raw[2], &tc.exp.Y},
		})
		tc.exp.Level = len(tc.code) - 2
		testCasesCode2XY = append(testCasesCode2XY, tc)
		return err
	})

	// http://geohex.net/testcase/hex_v3.2_test_XY2HEX.json
	loadTestCasesFromJson("hex_v3.2_test_XY2HEX.json", func(raw []json.RawMessage) error {
		tc := testCaseXY2HEX{}
		err := unmarshalRawFields([]testCaseFieldMap{
			{raw[0], &tc.level},
			{raw[1], &tc.x},
			{raw[2], &tc.y},
			{raw[3], &tc.exp},
		})
		testCasesXY2HEX = append(testCasesXY2HEX, tc)
		return err
	})

})

// --------------------------------------------------------------------

type testCaseFieldMap struct {
	f json.RawMessage
	v interface{}
}

func unmarshalRawFields(fields []testCaseFieldMap) error {
	for _, f := range fields {
		if err := json.Unmarshal(f.f, f.v); err != nil {
			return err
		}
	}
	return nil
}

func loadTestCasesFromJson(filename string, rawUnmarshal func([]json.RawMessage) error) {
	file, err := os.Open("../testdata/" + filename)
	Expect(err).NotTo(HaveOccurred())
	defer file.Close()

	var rawTcs [][]json.RawMessage
	Expect(json.NewDecoder(file).Decode(&rawTcs)).To(Succeed())

	for _, rtc := range rawTcs {
		Expect(rawUnmarshal(rtc)).To(Succeed())
	}
}
