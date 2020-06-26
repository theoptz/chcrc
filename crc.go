package chcrc

import "unsafe"

const (
	// The ecma polynomial, defined in ECMA 182.
	ecma = 0xC96C5795D7870F42
)

// table is a 256-word table representing the polynomial for efficient processing.
type table [256]uint64

// cached table
var ecmaTab *table

func init() {
	ecmaTab = makeTable(ecma)
}

func makeTable(poly uint64) *table {
	t := new(table)
	for i := 0; i < 256; i++ {
		crc := uint64(i)
		for j := 0; j < 8; j++ {
			if crc&1 == 1 {
				crc = (crc >> 1) ^ poly
			} else {
				crc >>= 1
			}
		}
		t[i] = crc
	}
	return t
}

// GetCrc64 returns crc64 checksum of provided string
// see for details: https://github.com/ClickHouse/ClickHouse/pull/7480/commits/2d2e738085f71731eab61571f592b3d4fcebfdc2
func GetCrc64(v string) uint64 {
	var crc uint64
	bytesArr := *(*[]byte)(unsafe.Pointer(&v))

	for _, by := range bytesArr {
		crc = ecmaTab[(byte(crc)^by)&0xff] ^ (crc >> 8)
	}

	return crc
}
