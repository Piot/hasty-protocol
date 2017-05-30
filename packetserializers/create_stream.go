package packetserializers

import (
	"bytes"
	"encoding/binary"

	log "github.com/sirupsen/logrus"

	"github.com/piot/hasty-protocol/asciistring"
	"github.com/piot/hasty-protocol/packet"
)

// CreateStreamToOctets : todo
func CreateStreamToOctets(request uint64, path string) []byte {
	buf := new(bytes.Buffer)
	buf.WriteByte(byte(packet.CreateStream))
	binary.Write(buf, binary.BigEndian, request)
	pathOctets, pathOctetsErr := asciistring.ToOctets(path)
	if pathOctetsErr != nil {
		log.Warnf("Couldn't serialize path")
	}
	buf.Write(pathOctets)
	return buf.Bytes()
}
