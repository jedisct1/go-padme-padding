package padme

import (
	"math/bits"
)

// PaddedLen - Return the padded length
func PaddedLen(len int) int {
	if len <= 0 {
		return 0
	}
	e := 63 - bits.LeadingZeros64(uint64(len))
	s := 64 - bits.LeadingZeros64(uint64(e))
	z := e - s
	mask := (uint64(1) << z) - 1
	return int((uint64(len) + mask) & ^mask)
}

// PaddingLen - Return the padding length
func PaddingLen(len int) int {
	return PaddedLen(len) - len
}
