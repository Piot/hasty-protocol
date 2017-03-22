package packetserializers

import (
	"bytes"
	"encoding/binary"

	"github.com/piot/hasty-protocol/packet"
	"github.com/piot/hasty-protocol/timestamp"
)

// PongToOctets : todo
func PongToOctets(millisecondsNow timestamp.Time, remoteMilliseconds timestamp.Time) []byte {
	buf := new(bytes.Buffer)
	buf.WriteByte(byte(packet.Pong))
	binary.Write(buf, binary.BigEndian, millisecondsNow.Raw())
	binary.Write(buf, binary.BigEndian, remoteMilliseconds.Raw())
	return buf.Bytes()
}
