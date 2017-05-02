package packetserializers

import (
	"bytes"
	"encoding/binary"

	"github.com/piot/hasty-protocol/channel"
	"github.com/piot/hasty-protocol/packet"
)

// LoginResultToOctets : todo
func LoginResultToOctets(channelID channel.ID) []byte {
	buf := new(bytes.Buffer)
	buf.WriteByte(byte(packet.LoginResult))
	buf.WriteByte(byte(1))
	binary.Write(buf, binary.BigEndian, channelID.Raw())
	return buf.Bytes()
}
