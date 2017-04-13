package packetdeserializers

import (
	"errors"

	"github.com/piot/hasty-protocol/commands"
	"github.com/piot/hasty-protocol/deserialize"
	"github.com/piot/hasty-protocol/packet"
)

// ToStreamData : Packet to publish
func ToStreamData(in packet.Packet) (commands.StreamData, error) {
	if in.Type() != packet.StreamData {
		return commands.StreamData{}, errors.New("Illegal packet type")
	}
	pos := 0
	payload := in.Payload()
	channel, idOctets, err := ToChannelID(payload)
	if err != nil {
		return commands.StreamData{}, err
	}
	pos += idOctets
	offset, offsetOctets, offsetErr := deserialize.ToStreamOffset(payload[pos:])
	if offsetErr != nil {
		return commands.StreamData{}, offsetErr
	}
	pos += offsetOctets

	isEndPosition := payload[pos] != 0

	chunk := payload[pos+1:]

	createdStreamData := commands.NewStreamData(channel, offset, chunk, isEndPosition)

	return createdStreamData, nil
}
