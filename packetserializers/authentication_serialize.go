package packetserializers

import (
	"bytes"
	"encoding/binary"

	"github.com/piot/hasty-protocol/authentication"
)

// AuthenticationInfoToOctets : todo
func AuthenticationInfoToOctets(info authentication.Info) []byte {
	buf := new(bytes.Buffer)
	binary.Write(buf, binary.BigEndian, info.UserID())
	binary.Write(buf, binary.BigEndian, info.UserAllocatedChannelID().Raw())
	return buf.Bytes()
}
