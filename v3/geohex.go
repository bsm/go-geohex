package geohex

import "fmt"

const VERSION = "3.2.0"

// MaxLevel is the maximum encoding level that this implementation supports
const MaxLevel = 20

// Error types
var (
	ErrLevelInvalid = fmt.Errorf("geohex: level invalid")
	ErrCodeInvalid  = fmt.Errorf("geohex: code invalid")
)

// Encode encodes a lat/lon/level into a code.
// Can return an ErrLevelInvalid code if level is not valid.
func Encode(lat, lon float64, level int) (string, error) {
	tile, err := newLL(lat, lon).tile(level)
	if err != nil {
		return "", err
	}

	return tile.Code(), nil
}

// Decode decodes a string code into Lat/Lon coordinates
// Can return ErrCodeInvalid if code is  not valid
func Decode(code string) (LL, error) {
	pos, err := DecodeTile(code)
	if err != nil {
		return LL{}, err
	}
	return pos.LL(), nil
}
