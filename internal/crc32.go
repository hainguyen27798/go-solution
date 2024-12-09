package internal

import "hash/crc32"

func StringToCRC32(s string) uint32 {
	// Create a CRC32 hash table with the IEEE polynomial
	table := crc32.MakeTable(crc32.IEEE)
	// Calculate the checksum for the input string
	return crc32.Checksum([]byte(s), table)
}
