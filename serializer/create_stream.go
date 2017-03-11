package serializer

import (
	"bytes"

	"github.com/piot/hasty-protocol/packet"
)

// CreateStreamToOctets : todo
func CreateStreamToOctets(path string) []byte {
	buf := new(bytes.Buffer)
	buf.WriteByte(byte(packet.CreateStream))
	buf.WriteByte(byte(len(path)))
	buf.Write([]byte(path))
	return buf.Bytes()
}
