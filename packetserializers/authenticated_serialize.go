package packetserializers

import (
	"bytes"
	"encoding/binary"

	"github.com/piot/hasty-protocol/authentication"
	"github.com/piot/hasty-protocol/packet"
	"github.com/piot/hasty-protocol/unicodestring"
)

// AuthenticatedToOctets : todo
func AuthenticatedToOctets(info authentication.Authenticated) []byte {
	buf := new(bytes.Buffer)
	buf.WriteByte(byte(packet.Authenticated))
	binary.Write(buf, binary.BigEndian, info.UserID().Raw())
	binary.Write(buf, binary.BigEndian, info.UserAllocatedChannelID().Raw())
	realnameOctets, _ := unicodestring.ToOctets(info.Realname())
	buf.Write(realnameOctets)
	return buf.Bytes()
}
