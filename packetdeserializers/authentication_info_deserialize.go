package packetdeserializers

import (
	"encoding/binary"

	"github.com/piot/hasty-protocol/authentication"
	"github.com/piot/hasty-protocol/channel"
	"github.com/piot/hasty-protocol/user"
)

// OctetsToAuthenticationInfo : todo
func OctetsToAuthenticationInfo(octets []byte) (authentication.Info, int) {
	pos := 0
	userIDValue := binary.BigEndian.Uint64(octets[pos : pos+8])
	userID, _ := user.NewID(userIDValue)
	pos += 8
	userAllocatedChannelIDValue := binary.BigEndian.Uint64(octets[pos : pos+8])
	userAllocatedChannelID, _ := channel.NewFromID(uint32(userAllocatedChannelIDValue))
	info := authentication.NewInfo(userID, userAllocatedChannelID)
	return info, pos
}
