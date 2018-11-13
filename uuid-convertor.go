package convertor

import (
	"math/big"
	"strings"
)

var (
	removeDashes = strings.NewReplacer("-", "")
)

// UUIDToIntString takes a connical UUID e.g 8fa7158f-f262-44c2-8633-f02298624c52
// and converts it to a string representation of it's numeric type e.g 190947154307861374080435443858447420498
func UUIDToIntString(uuid string) string {
	var i big.Int
	i.SetString(removeDashes.Replace(uuid), 16)
	return i.String()
}

// UUIDToUInt takes a connical UUID e.g 8fa7158f-f262-44c2-8633-f02298624c52
// and converts it to a uint64 it's numeric type e.g 9670336856270720082
func UUIDToUInt(uuid string) uint64 {
	var i big.Int
	i.SetString(removeDashes.Replace(uuid), 16)
	return i.Uint64()
}
