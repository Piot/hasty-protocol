package packet

import (
	"encoding/hex"
	"fmt"
)

// Packet : An incoming packet
type Packet struct {
	header  Header
	payload []byte
}

// NewPacket : Creates a new incoming packet
func NewPacket(header Header, payload []byte) Packet {
	return Packet{header: header, payload: payload}
}

// Type : Returns the packet type
func (in Packet) Type() Type {
	return in.header.packetType
}

// Payload : Returns the actual octets
func (in Packet) Payload() []byte {
	return in.payload
}

// String : A human readable representation of a packet
func (in Packet) String() string {
	hexPayload := hex.Dump(in.Payload())
	return fmt.Sprintf("[Packet %s %s]", in.header, hexPayload)
}
