package packetserializers

import (
	"bytes"

	log "github.com/sirupsen/logrus"

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
		log.Warnf("We couldn't write length")
		return nil, lengthErr
	}
	buf.Write(lengthBuf)
	buf.Write(infoOctets)
	buf.Write(payload)
	encodedPayload := buf.Bytes()
	return encodedPayload, nil
}
