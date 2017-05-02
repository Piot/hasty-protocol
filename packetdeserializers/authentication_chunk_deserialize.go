package packetdeserializers

import (
	"fmt"

	"github.com/piot/hasty-protocol/authentication"
	"github.com/piot/hasty-protocol/deserialize"
)

// OctetsToAuthenticationChunk : todo
func OctetsToAuthenticationChunk(payload []byte) (authentication.Info, []byte, error) {
	pos := 0
	authenticationLength, authenticationLengthOctetCount, _ := deserialize.ToSmallLength(payload)
	pos += authenticationLengthOctetCount
	authenticationInfo, authenticationInfoCount := OctetsToAuthenticationInfo(payload[pos:])
	pos += authenticationInfoCount
	totalLength := len(payload)
	foundLength := int(totalLength) - pos
	if foundLength != int(authenticationLength) {
		return authentication.Info{}, nil, fmt.Errorf("Authentication length is wrong:%d detected %d", authenticationLength, foundLength)
	}
	chunk := payload[pos:]
	return authenticationInfo, chunk, nil
}
