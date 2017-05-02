package packetserializers

import (
	"bytes"

	"github.com/piot/hasty-protocol/packet"
)

// LoginToOctets : todo
func LoginToOctets(username string, password string) []byte {
	buf := new(bytes.Buffer)
	buf.WriteByte(byte(packet.Login))
	buf.WriteString(username)
	buf.WriteString(password)
	return buf.Bytes()
}
