package serializer

import (
	"bytes"
	"encoding/binary"
	"log"

	"github.com/piot/hasty-protocol/channel"
	"github.com/piot/hasty-protocol/packet"
)

// StreamDataToOctets : todo
func StreamDataToOctets(channel channel.ID, offset uint32, data []byte) []byte {
	buf := new(bytes.Buffer)
	buf.WriteByte(byte(packet.StreamData))
	binary.Write(buf, binary.BigEndian, channel.Raw())
	binary.Write(buf, binary.BigEndian, offset)
	lengthBuf, lengthErr := SmallLengthToOctets(uint16(len(data)))
	if lengthErr != nil {
		log.Fatalf("We couldn't write length")
		return []byte{}
	}
	buf.Write(lengthBuf)
	buf.Write(data)
	return buf.Bytes()
}
