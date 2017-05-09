package packetserializers

import (
	"bytes"
	"encoding/binary"
	"log"

	"github.com/piot/hasty-protocol/packet"
	"github.com/piot/hasty-protocol/serializer"
	"github.com/piot/hasty-protocol/user"
)

// PublishStreamUserToOctets : todo
func PublishStreamUserToOctets(userID user.ID, chunk []byte) []byte {
	buf := new(bytes.Buffer)
	buf.WriteByte(byte(packet.PublishStreamUser))
	binary.Write(buf, binary.BigEndian, userID.Raw())
	lengthBuf, lengthErr := serializer.SmallLengthToOctets(uint16(len(chunk)))
	if lengthErr != nil {
		log.Printf("We couldn't write length")
		return []byte{}
	}
	buf.Write(lengthBuf)
	buf.Write(chunk)
	return buf.Bytes()
}
