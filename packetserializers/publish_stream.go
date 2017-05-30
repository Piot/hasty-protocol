package packetserializers

import (
	"bytes"
	"encoding/binary"

	log "github.com/sirupsen/logrus"

	"github.com/piot/hasty-protocol/channel"
	"github.com/piot/hasty-protocol/packet"
	"github.com/piot/hasty-protocol/serializer"
)

// PublishStreamToOctets : todo
func PublishStreamToOctets(channel channel.ID, chunk []byte) []byte {
	buf := new(bytes.Buffer)
	buf.WriteByte(byte(packet.PublishStream))
	binary.Write(buf, binary.BigEndian, channel.Raw())
	lengthBuf, lengthErr := serializer.SmallLengthToOctets(uint16(len(chunk)))
	if lengthErr != nil {
		log.Warnf("We couldn't write length")
		return []byte{}
	}
	buf.Write(lengthBuf)
	buf.Write(chunk)
	return buf.Bytes()
}
