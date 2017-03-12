package serializer

import (
	"bytes"
	"encoding/binary"
	"log"

	"github.com/piot/hasty-protocol/channel"
	"github.com/piot/hasty-protocol/packet"
)

// PublishStreamToOctets : todo
func PublishStreamToOctets(channel channel.ID, chunk []byte) []byte {
	buf := new(bytes.Buffer)
	buf.WriteByte(byte(packet.PublishStream))
	binary.Write(buf, binary.BigEndian, channel.Raw())
	lengthBuf, lengthErr := SmallLengthToOctets(uint16(len(chunk)))
	if lengthErr != nil {
		log.Fatalf("We couldn't write length")
		return []byte{}
	}
	buf.Write(lengthBuf)
	buf.Write(chunk)
	return buf.Bytes()
}
