package serializer

import (
	"bytes"
	"encoding/binary"

	"github.com/piot/hasty-protocol/channel"
	"github.com/piot/hasty-protocol/packet"
)

// StreamDataToOctets : todo
func StreamDataToOctets(channel channel.ID, offset uint32, data []byte) []byte {
	buf := new(bytes.Buffer)
	buf.WriteByte(byte(packet.StreamData))
	binary.Write(buf, binary.BigEndian, channel.Raw())
	binary.Write(buf, binary.BigEndian, offset)
	buf.WriteByte(byte(len(data)))
	buf.Write(data)
	return buf.Bytes()
}
