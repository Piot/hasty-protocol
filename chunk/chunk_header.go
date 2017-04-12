package chunk

import (
	"fmt"

	"github.com/piot/hasty-protocol/deserialize"
)

// Header : A Packet header
type Header struct {
	timestamp       uint64
	userID          uint64
	deviceID        uint64
	payloadSize     int
	headerOctetSize int
}

func (in Header) String() string {
	return fmt.Sprintf("[ChunkHeader payloadSize:%d]", in.payloadSize)
}

func tryReadLength(buf []byte) (int, int, error) {
	length, octetsUsed, err := deserialize.ToSmallLength(buf)
	return int(length), octetsUsed, err
}

// CheckIfWeHaveChunkHeader : Returns a chunk if it is ready
func CheckIfWeHaveChunkHeader(buf []byte) (packet Header, packetWasReady bool, err error) {
	bufSize := len(buf)
	if bufSize < 1 {
		return Header{}, false, nil
	}
	packetLength, octetsUsedForLengthEncoding, _ := tryReadLength(buf)
	octetSize := len(buf)
	if octetSize < packetLength+octetsUsedForLengthEncoding {
		return Header{}, false, nil
	}
	headerOctetSize := octetsUsedForLengthEncoding

	return Header{payloadSize: packetLength - 1, headerOctetSize: headerOctetSize}, true, nil
}
