package chunk

import (
	"bytes"
	"encoding/hex"
	"errors"

	log "github.com/sirupsen/logrus"

	"github.com/piot/hasty-protocol/channel"
	"github.com/piot/hasty-protocol/packet"
)

// Stream : The input octet stream
type Stream struct {
	buffer       bytes.Buffer
	connectionID packet.ConnectionID
	channelID    channel.ID
}

// NewChunkStream : Creates and input stream
func NewChunkStream(connectionID packet.ConnectionID, channelID channel.ID) *Stream {
	stream := &Stream{connectionID: connectionID, channelID: channelID}
	return stream
}

// Feed : Adds octets to stream
func (stream *Stream) Feed(octets []byte) error {
	lengthWritten, err := stream.buffer.Write(octets)
	if err != nil {
		return err
	}

	if lengthWritten != len(octets) {
		err := errors.New("Couldn't write all octets")
		return err
	}
	if true {
		hexPayload := hex.Dump(stream.buffer.Bytes())
		log.Debugf("Buffer is now:%s", hexPayload)
	}
	return nil
}

func (stream *Stream) octets() []byte {
	return stream.buffer.Bytes()
}

// FetchChunk : Converts from a stream to a packet if possible
func (stream *Stream) FetchChunk() (chunk Chunk, err error) {
	packetHeader, packetHeaderIsDone, packetHeaderErr := CheckIfWeHaveChunkHeader(stream.octets())

	if packetHeaderErr != nil {
		return Chunk{}, packetHeaderErr
	}

	if !packetHeaderIsDone {
		return Chunk{}, &NotDoneError{msg: "Header is not done"}
	}

	availableOctets := stream.buffer.Len() - packetHeader.headerOctetSize
	if availableOctets < packetHeader.payloadSize {
		return Chunk{}, &NotDoneError{msg: "Payload is not done"}
	}
	completeChunkPayloadOctetCount := packetHeader.payloadSize + packetHeader.headerOctetSize
	packetPayload := make([]byte, completeChunkPayloadOctetCount)
	stream.buffer.Read(packetPayload)

	foundPacket := NewChunk(packetHeader, packetPayload)
	log.Debugf("%s %s Received Chunk %s", stream.connectionID, stream.channelID, foundPacket)
	return foundPacket, nil
}
