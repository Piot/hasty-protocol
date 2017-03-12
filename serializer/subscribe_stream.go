package serializer

import (
	"bytes"
	"encoding/binary"

	"github.com/piot/hasty-protocol/channel"
	"github.com/piot/hasty-protocol/packet"
	"github.com/piot/hasty-protocol/qos"
)

// SubscribeStreamToOctets : todo
func SubscribeStreamToOctets(channel channel.ID, offset uint32) []byte {
	buf := new(bytes.Buffer)
	buf.WriteByte(byte(packet.SubscribeStream))
	buf.WriteByte(byte(1))
	binary.Write(buf, binary.BigEndian, channel.Raw())
	buf.WriteByte(byte(qos.Normal))
	binary.Write(buf, binary.BigEndian, offset)
	return buf.Bytes()
}
