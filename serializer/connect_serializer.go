package serializer

import (
	"bytes"

	"github.com/piot/hasty-protocol/packet"
)

// ConnectResultToOctets : todo
func ConnectResultToOctets() []byte {
	buf := new(bytes.Buffer)
	buf.WriteByte(byte(packet.ConnectResult))
	buf.WriteByte(byte(1))
	return buf.Bytes()
}
