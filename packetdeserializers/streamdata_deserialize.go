package packetdeserializers

import (
	"errors"
	"fmt"

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
	pos++

	payloadLength, payloadLengthOctetsUsed, payloadLengthErr := deserialize.ToSmallLength(payload[pos:])
	if payloadLengthErr != nil {
		return commands.StreamData{}, payloadLengthErr
	}
	pos += payloadLengthOctetsUsed

	chunk := payload[pos:]
	if len(chunk) != int(payloadLength) {
		return commands.StreamData{}, fmt.Errorf("Illegal length:%d lengthOfChunk:%d", payloadLength, len(chunk))
	}
	createdStreamData := commands.NewStreamData(channel, offset, chunk, isEndPosition)

	return createdStreamData, nil
}
