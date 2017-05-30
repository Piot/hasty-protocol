package packetserializers

import (
	"bytes"
	"encoding/binary"

	log "github.com/sirupsen/logrus"

	"github.com/piot/hasty-protocol/channel"
	"github.com/piot/hasty-protocol/packet"
	"github.com/piot/hasty-protocol/serializer"
)

// StreamDataToOctets : todo
func StreamDataToOctets(channel channel.ID, offset uint32, data []byte) []byte {
	buf := new(bytes.Buffer)
	buf.WriteByte(byte(packet.StreamData))
	binary.Write(buf, binary.BigEndian, channel.Raw())
	binary.Write(buf, binary.BigEndian, offset)
	buf.WriteByte(0)
	lengthBuf, lengthErr := serializer.SmallLengthToOctets(uint16(len(data)))
	if lengthErr != nil {
		log.Warnf("We couldn't write length")
		return []byte{}
	}
	buf.Write(lengthBuf)
	buf.Write(data)
	return buf.Bytes()
}
