package deserialize

import (
	"encoding/binary"

	"github.com/piot/hasty-protocol/streamoffset"
)

// ToStreamOffset : convert from octets
func ToStreamOffset(octets []byte) (streamoffset.Offset, int, error) {
	value := binary.BigEndian.Uint32(octets)
	offset, err := streamoffset.NewFromOffset(value)
	return offset, 4, err
}
