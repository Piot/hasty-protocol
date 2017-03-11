package packetserializers

import (
	"encoding/binary"

	"github.com/piot/hasty-protocol/channel"
)

// ToChannelID : convert from octets
func ToChannelID(octets []byte) (channel.ID, int, error) {
	value := binary.BigEndian.Uint32(octets)
	channelID, err := channel.NewFromID(value)
	return channelID, 4, err
}
