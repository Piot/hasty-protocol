package packetserializers

import (
	"encoding/binary"

	"github.com/piot/hasty-protocol/commands"
	"github.com/piot/hasty-protocol/packet"
	"github.com/piot/hasty-protocol/timestamp"
)

// ToPongFromOctets : convert from octets
func ToPongFromOctets(octets []byte) (uint64, uint64, int, error) {
	remoteNow := binary.BigEndian.Uint64(octets)
	myTimeEchoedBack := binary.BigEndian.Uint64(octets)
	return remoteNow, myTimeEchoedBack, 16, nil
}

// ToPong : convert from packet
func ToPong(in packet.Packet) (commands.Pong, error) {
	payload := in.Payload()
	nowValue, echoedValue, _, toPongErr := ToPongFromOctets(payload)
	if toPongErr != nil {
		return commands.Pong{}, toPongErr
	}
	now := timestamp.New(nowValue)
	echoedTime := timestamp.New(echoedValue)
	return commands.NewPong(now, echoedTime), nil
}
