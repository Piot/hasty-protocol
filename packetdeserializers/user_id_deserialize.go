package packetdeserializers

import (
	"encoding/binary"

	"github.com/piot/hasty-protocol/user"
)

// ToUserID : convert from octets
func ToUserID(octets []byte) (user.ID, int, error) {
	value := binary.BigEndian.Uint64(octets)
	userID, err := user.NewID(value)
	return userID, 8, err
}
