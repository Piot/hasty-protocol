package packetserializers

import (
	"bytes"
	"encoding/binary"

	"github.com/piot/hasty-protocol/channel"
	"github.com/piot/hasty-protocol/packet"
)

// LoginResultToOctets : todo
func LoginResultToOctets(successful bool, channelID channel.ID) []byte {
	buf := new(bytes.Buffer)
	buf.WriteByte(byte(packet.LoginResult))
	var loginSuccessfulOctet uint8
	if successful {
		loginSuccessfulOctet = 1
	}
	buf.WriteByte(loginSuccessfulOctet)
	binary.Write(buf, binary.BigEndian, channelID.Raw())
	return buf.Bytes()
}
