package serializer

import (
	"bytes"
	"encoding/binary"

	"github.com/piot/hasty-protocol/channel"
	"github.com/piot/hasty-protocol/packet"
)

// PublishStreamToOctets : todo
func PublishStreamToOctets(channel channel.ID, chunk []byte) []byte {
	buf := new(bytes.Buffer)
	buf.WriteByte(byte(packet.PublishStream))
	binary.Write(buf, binary.BigEndian, channel.Raw())
	buf.WriteByte(byte(len(chunk)))
	buf.Write(chunk)
	return buf.Bytes()
}
