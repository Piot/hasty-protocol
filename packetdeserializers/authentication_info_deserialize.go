package packetdeserializers

import (
	"encoding/binary"

	"github.com/piot/hasty-protocol/authentication"
	"github.com/piot/hasty-protocol/user"
)

// OctetsToAuthenticationInfo : todo
func OctetsToAuthenticationInfo(octets []byte) (authentication.Info, int) {
	pos := 0
	userIDValue := binary.BigEndian.Uint64(octets[pos : pos+8])
	pos += 8
	userID, _ := user.NewID(userIDValue)
	info := authentication.NewInfo(userID)
	return info, pos
}
