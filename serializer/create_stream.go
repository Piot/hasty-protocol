package serializer

import (
	"bytes"
	"log"

	"github.com/piot/hasty-protocol/packet"
)

// CreateStreamToOctets : todo
func CreateStreamToOctets(path string) []byte {
	buf := new(bytes.Buffer)
	buf.WriteByte(byte(packet.CreateStream))
	lengthBuf, lengthErr := SmallLengthToOctets(uint16(len(path)))
	if lengthErr != nil {
		log.Fatalf("We couldn't write length")
		return []byte{}
	}
	buf.Write(lengthBuf)
	buf.Write([]byte(path))
	return buf.Bytes()
}
