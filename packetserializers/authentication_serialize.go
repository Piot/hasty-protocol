package packetserializers

import (
	"bytes"
	"encoding/binary"

	"github.com/piot/hasty-protocol/authentication"
)

// AuthenticationInfoToOctets : todo
func AuthenticationInfoToOctets(info authentication.Info) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, info.UserID().Raw())
	binary.Write(buf, binary.BigEndian, uint64(info.UserAllocatedChannelID().Raw()))
	return buf.Bytes()
}
