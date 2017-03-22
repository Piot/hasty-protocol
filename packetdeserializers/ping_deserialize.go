package packetdeserializers

import (
	"encoding/binary"

	"github.com/piot/hasty-protocol/commands"
	"github.com/piot/hasty-protocol/packet"
	"github.com/piot/hasty-protocol/timestamp"
)

// ToPingFromOctets : convert from octets
func ToPingFromOctets(octets []byte) (uint64, int, error) {
	value := binary.BigEndian.Uint64(octets)
	return value, 8, nil
}

// ToPing : convert from packet
func ToPing(in packet.Packet) (commands.Ping, error) {
	payload := in.Payload()
	v, _, toPingErr := ToPingFromOctets(payload)
	if toPingErr != nil {
		return commands.Ping{}, toPingErr
	}
	ts := timestamp.New(v)
	return commands.NewPing(ts), nil
}
