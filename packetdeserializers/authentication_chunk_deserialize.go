package packetdeserializers

import (
	"fmt"
	"log"

	"github.com/piot/hasty-protocol/authentication"
	"github.com/piot/hasty-protocol/deserialize"
)

// OctetsToAuthenticationChunk : todo
func OctetsToAuthenticationChunk(payload []byte) (authentication.Info, []byte, error) {
	pos := 0
	/*
	 */
	authenticationInfo, authenticationInfoCount := OctetsToAuthenticationInfo(payload[pos:])
	log.Printf("UserID %v channel %v ", authenticationInfo.UserID(), authenticationInfo.UserAllocatedChannelID())
	pos += authenticationInfoCount
	commandLength, commandLengthOctetCount, _ := deserialize.ToSmallLength(payload[pos:])
	pos += commandLengthOctetCount
	rest := uint16(len(payload) - pos)
	if rest != commandLength {
		return authentication.Info{}, nil, fmt.Errorf("Illegal length %v %v", rest, commandLength)
	}
	chunk := payload[pos:]
	return authenticationInfo, chunk, nil
}
