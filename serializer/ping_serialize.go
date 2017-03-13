package serializer

import (
	"bytes"
	"encoding/binary"

	"github.com/piot/hasty-protocol/packet"
)

// PingToOctets : todo
func PingToOctets(milliseconds uint64) []byte {
	buf := new(bytes.Buffer)
	buf.WriteByte(byte(packet.Ping))
	binary.Write(buf, binary.BigEndian, milliseconds)
	return buf.Bytes()
}
