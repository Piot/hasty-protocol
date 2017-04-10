package packetserializers

import (
	"bytes"

	"github.com/piot/hasty-protocol/packet"
)

// LoginResultToOctets : todo
func LoginResultToOctets() []byte {
	buf := new(bytes.Buffer)
	buf.WriteByte(byte(packet.LoginResult))
	buf.WriteByte(byte(1))
	return buf.Bytes()
}
