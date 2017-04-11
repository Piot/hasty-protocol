package packet

import (
	"bytes"
	"encoding/hex"
	"errors"
	"log"
)

// Stream : The input octet stream
type Stream struct {
	buffer       bytes.Buffer
	connectionID ConnectionID
}

// NewPacketStream : Creates and input stream
func NewPacketStream(connectionID ConnectionID) Stream {
	stream := Stream{connectionID: connectionID}
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
	if false {
		hexPayload := hex.Dump(stream.buffer.Bytes())
		log.Printf("Buffer is now:%s", hexPayload)
	}
	return nil
}

func (stream *Stream) octets() []byte {
	return stream.buffer.Bytes()
}

// FetchPacket : Converts from a stream to a packet if possible
func (stream *Stream) FetchPacket() (packet Packet, err error) {
	packetHeader, packetHeaderIsDone, packetHeaderErr := CheckIfWeHavePacketHeader(stream.octets())

	if packetHeaderErr != nil {
		return Packet{}, packetHeaderErr
	}

	if !packetHeaderIsDone {
		return Packet{}, &PacketNotDoneError{msg: "Header is not done"}
	}

	availableOctets := stream.buffer.Len() - packetHeader.headerOctetSize
	if availableOctets < packetHeader.payloadSize {
		return Packet{}, &PacketNotDoneError{msg: "Payload is not done"}
	}

	skip := make([]byte, packetHeader.headerOctetSize)
	stream.buffer.Read(skip)

	packetPayload := make([]byte, packetHeader.payloadSize)
	stream.buffer.Read(packetPayload)

	foundPacket := NewPacket(packetHeader, packetPayload)
	log.Printf("%s Received Packet %s", stream.connectionID, foundPacket)
	return foundPacket, nil
}
