package packetserializers

import (
	"bytes"
	"encoding/hex"
	"log"

	"github.com/piot/hasty-protocol/authentication"
	"github.com/piot/hasty-protocol/serializer"
)

// AuthenticationChunkToOctets : todo
func AuthenticationChunkToOctets(info authentication.Info, payload []byte) ([]byte, error) {
	buf := new(bytes.Buffer)
	payloadLength := len(payload)
	infoOctets := AuthenticationInfoToOctets(info)
	chunkLength := uint16(payloadLength + len(infoOctets))
	lengthBuf, lengthErr := serializer.SmallLengthToOctets(chunkLength)
	if lengthErr != nil {
		log.Printf("We couldn't write length")
		return nil, lengthErr
	}
	log.Printf("Writing authentication length: %v user:%v channel:%v", chunkLength, info.UserID(), info.UserAllocatedChannelID())
	buf.Write(lengthBuf)
	buf.Write(infoOctets)
	buf.Write(payload)
	encodedPayload := buf.Bytes()
	log.Printf("Authentication serialize len %d payload:%s", chunkLength, hex.Dump(encodedPayload))
	return encodedPayload, nil
}
