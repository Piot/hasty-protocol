package deserialize

import "encoding/binary"

// ToOffset : convert from octets
func ToOffset(octets []byte) (uint32, int, error) {
	value := binary.BigEndian.Uint32(octets)
	return value, 4, nil
}
