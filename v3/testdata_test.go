package geohex

import (
	"encoding/json"
	"os"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

// Test cases downloaded from http://geohex.net/testcase/v3.2.html

type (
	testCaseCode2LL struct {
		code string
		exp  LL
	}
	testCaseLL2Code struct {
		level int
		ll    LL
		exp   string
	}
	testCaseLL2Pos struct {
		level int
		ll    LL
		expX  int
		expY  int
	}
	testCaseCode2Pos struct {
		code string
		exp  Position
	}
	testCasePos2Code struct {
		level int
		x     int
		y     int
		exp   string
	}
)

var (
	testCasesCode2LL  []testCaseCode2LL
	testCasesLL2Code  []testCaseLL2Code
	testCasesLL2Pos   []testCaseLL2Pos
	testCasesCode2Pos []testCaseCode2Pos
	testCasesPos2Code []testCasePos2Code
)

var _ = BeforeSuite(func() {

	// http://geohex.net/testcase/hex_v3.2_test_code2HEX.json
	loadTestCasesFromJson("hex_v3.2_test_code2HEX.json", func(raw []json.RawMessage) error {
		tc := testCaseCode2LL{}
		err := unmarshalRawFields([]testCaseFieldMap{
			{raw[0], &tc.code},
			{raw[1], &tc.exp.Lat},
			{raw[2], &tc.exp.Lon},
		})
		testCasesCode2LL = append(testCasesCode2LL, tc)
		return err
	})

	// http://geohex.net/testcase/hex_v3.2_test_coord2HEX.json
	loadTestCasesFromJson("hex_v3.2_test_coord2HEX.json", func(raw []json.RawMessage) error {
		tc := testCaseLL2Code{}
		err := unmarshalRawFields([]testCaseFieldMap{
			{raw[0], &tc.level},
			{raw[1], &tc.ll.Lat},
			{raw[2], &tc.ll.Lon},
			{raw[3], &tc.exp},
		})
		testCasesLL2Code = append(testCasesLL2Code, tc)
		return err
	})

	// http://geohex.net/testcase/hex_v3.2_test_coord2XY.json
	loadTestCasesFromJson("hex_v3.2_test_coord2XY.json", func(raw []json.RawMessage) error {
		tc := testCaseLL2Pos{}
		err := unmarshalRawFields([]testCaseFieldMap{
			{raw[0], &tc.level},
			{raw[1], &tc.ll.Lat},
			{raw[2], &tc.ll.Lon},
			{raw[3], &tc.expX},
			{raw[4], &tc.expY},
		})
		testCasesLL2Pos = append(testCasesLL2Pos, tc)
		return err
	})

	// http://geohex.net/testcase/hex_v3.2_test_code2XY.json
	loadTestCasesFromJson("hex_v3.2_test_code2XY.json", func(raw []json.RawMessage) error {
		tc := testCaseCode2Pos{}
		err := unmarshalRawFields([]testCaseFieldMap{
			{raw[0], &tc.code},
			{raw[1], &tc.exp.X},
			{raw[2], &tc.exp.Y},
		})
		testCasesCode2Pos = append(testCasesCode2Pos, tc)
		return err
	})

	// http://geohex.net/testcase/hex_v3.2_test_XY2HEX.json
	loadTestCasesFromJson("hex_v3.2_test_XY2HEX.json", func(raw []json.RawMessage) error {
		tc := testCasePos2Code{}
		err := unmarshalRawFields([]testCaseFieldMap{
			{raw[0], &tc.level},
			{raw[1], &tc.x},
			{raw[2], &tc.y},
			{raw[3], &tc.exp},
		})
		testCasesPos2Code = append(testCasesPos2Code, tc)
		return err
	})

})

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
	file, err := os.Open("testdata/" + filename)
	Expect(err).NotTo(HaveOccurred())
	defer file.Close()

	var rawTcs [][]json.RawMessage
	Expect(json.NewDecoder(file).Decode(&rawTcs)).To(Succeed())

	for _, rtc := range rawTcs {
		Expect(rawUnmarshal(rtc)).To(Succeed())
	}
}
