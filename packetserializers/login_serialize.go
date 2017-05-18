package packetserializers

import (
	"bytes"

	"github.com/piot/hasty-protocol/packet"
	"github.com/piot/hasty-protocol/unicodestring"
)

// LoginToOctets : todo
func LoginToOctets(username string, password string) []byte {
	buf := new(bytes.Buffer)
	buf.WriteByte(byte(packet.Login))
	usernameOctets, _ := unicodestring.ToOctets(username)
	buf.Write(usernameOctets)
	passwordOctets, _ := unicodestring.ToOctets(password)
	buf.Write(passwordOctets)
	return buf.Bytes()
}
