package geohex

import (
	"encoding/json"
	"log"
	"os"
)

// Test cases downloaded from http://geohex.net/testcase/v3.2.html

type code2HexTestCase struct {
	code       string
	expectedLL LL
}

func loadCode2HexTestCases() []code2HexTestCase {
	var tcs []code2HexTestCase
	// http://geohex.net/testcase/hex_v3.2_test_code2HEX.json
	loadTestCasesFromJson("hex_v3.2_test_code2HEX.json", func(raw []json.RawMessage) error {
		tc := code2HexTestCase{}
		err := unmarshalRawFields([]tcFieldMapping{
			{raw[0], &tc.code},
			{raw[1], &tc.expectedLL.Lat},
			{raw[2], &tc.expectedLL.Lon},
		})
		tcs = append(tcs, tc)
		return err
	})
	return tcs
}

type ll2hexTestCase struct {
	level        int
	ll           LL
	expectedCode string
}

func loadLL2HexTestCases() []ll2hexTestCase {
	var tcs []ll2hexTestCase
	// http://geohex.net/testcase/hex_v3.2_test_coord2HEX.json
	loadTestCasesFromJson("hex_v3.2_test_coord2HEX.json", func(raw []json.RawMessage) error {
		tc := ll2hexTestCase{}
		err := unmarshalRawFields([]tcFieldMapping{
			{raw[0], &tc.level},
			{raw[1], &tc.ll.Lat},
			{raw[2], &tc.ll.Lon},
			{raw[3], &tc.expectedCode},
		})
		tcs = append(tcs, tc)
		return err
	})
	return tcs
}

type ll2PositionTestCase struct {
	level            int
	ll               LL
	expectedPosition Position
}

func loadLL2PositionTestCases() []ll2PositionTestCase {
	var tcs []ll2PositionTestCase
	// http://geohex.net/testcase/hex_v3.2_test_coord2XY.json
	loadTestCasesFromJson("hex_v3.2_test_coord2XY.json", func(raw []json.RawMessage) error {
		tc := ll2PositionTestCase{}
		err := unmarshalRawFields([]tcFieldMapping{
			{raw[0], &tc.level},
			{raw[1], &tc.ll.Lat},
			{raw[2], &tc.ll.Lon},
			{raw[3], &tc.expectedPosition.X},
			{raw[4], &tc.expectedPosition.Y},
		})
		tcs = append(tcs, tc)
		return err
	})
	return tcs
}

type code2PositionTestCase struct {
	code             string
	expectedPosition Position
}

func loadCode2PositionTestCases() []code2PositionTestCase {
	var tcs []code2PositionTestCase
	// http://geohex.net/testcase/hex_v3.2_test_code2XY.json
	loadTestCasesFromJson("hex_v3.2_test_code2XY.json", func(raw []json.RawMessage) error {
		tc := code2PositionTestCase{}
		err := unmarshalRawFields([]tcFieldMapping{
			{raw[0], &tc.code},
			{raw[1], &tc.expectedPosition.X},
			{raw[2], &tc.expectedPosition.Y},
		})
		tcs = append(tcs, tc)
		return err
	})
	return tcs
}

type position2hexTestCase struct {
	level        int
	position     Position
	expectedCode string
}

func loadPosition2HexTestCases() []position2hexTestCase {
	var tcs []position2hexTestCase
	// http://geohex.net/testcase/hex_v3.2_test_XY2HEX.json
	loadTestCasesFromJson("hex_v3.2_test_XY2HEX.json", func(raw []json.RawMessage) error {
		tc := position2hexTestCase{}
		err := unmarshalRawFields([]tcFieldMapping{
			{raw[0], &tc.level},
			{raw[1], &tc.position.X},
			{raw[2], &tc.position.Y},
			{raw[3], &tc.expectedCode},
		})
		tcs = append(tcs, tc)
		return err
	})
	return tcs
}

type tcFieldMapping struct {
	f json.RawMessage
	v interface{}
}

func unmarshalRawFields(fields []tcFieldMapping) error {
	for _, f := range fields {
		if err := json.Unmarshal(f.f, f.v); err != nil {
			return err
		}
	}
	return nil
}

func loadTestCasesFromJson(filename string, rawUnmarshal func([]json.RawMessage) error) {
	file, err := os.Open("testcases/" + filename)
	if err != nil {
		log.Fatalf("Error: %s", err)
		return
	}
	decoder := json.NewDecoder(file)
	rawTcs := make([][]json.RawMessage, 0)
	err = decoder.Decode(&rawTcs)
	if err != nil {
		log.Fatalf("Error: %s", err)
		return
	}
	for _, rtc := range rawTcs {
		if err := rawUnmarshal(rtc); err != nil {
			log.Fatalf("Error: %s", err)
			return
		}
	}
}
