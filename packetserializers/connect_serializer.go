package packetserializers

import (
	"bytes"

	"github.com/piot/hasty-protocol/packet"
	"github.com/piot/hasty-protocol/serializer"
	"github.com/piot/hasty-protocol/unicodestring"
	"github.com/piot/hasty-protocol/version"
)

// ConnectResultToOctets : todo
func ConnectResultToOctets() []byte {
	buf := new(bytes.Buffer)
	buf.WriteByte(byte(packet.ConnectResult))
	buf.WriteByte(byte(1))
	return buf.Bytes()
}

// ConnectToOctets : todo
func ConnectToOctets(version version.Version, realm string) []byte {
	buf := new(bytes.Buffer)
	buf.WriteByte(byte(packet.Connect))
	buf.Write(serializer.VersionToOctets(version))
	realmOctets, _ := unicodestring.ToOctets(realm)
	buf.Write(realmOctets)

	return buf.Bytes()
}
