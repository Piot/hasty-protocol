package packetserializers

import (
	"bytes"
	"encoding/binary"

	"github.com/piot/hasty-protocol/channel"
	"github.com/piot/hasty-protocol/packet"
)

// CreateStreamResultToOctets : todo
func CreateStreamResultToOctets(request uint64, channelID channel.ID) []byte {
	buf := new(bytes.Buffer)
	buf.WriteByte(byte(packet.CreateStreamResult))
	binary.Write(buf, binary.BigEndian, request)
	binary.Write(buf, binary.BigEndian, channelID.Raw())
	return buf.Bytes()
}
