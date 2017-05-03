package packetdeserializers

import (
	"errors"
	"fmt"

	"github.com/piot/hasty-protocol/commands"
	"github.com/piot/hasty-protocol/deserialize"
	"github.com/piot/hasty-protocol/packet"
)

// ToPublishStream : Packet to publish
func ToPublishStream(in packet.Packet) (commands.PublishStream, error) {
	if in.Type() != packet.PublishStream {
		return commands.PublishStream{}, errors.New("Illegal packet type")
	}

	payload := in.Payload()
	pos := 0
	channel, idOctets, err := ToChannelID(payload[pos:])
	if err != nil {
		return commands.PublishStream{}, err
	}
	pos += idOctets

	payloadLength, payloadLengthOctetsUsed, payloadLengthErr := deserialize.ToSmallLength(payload[pos:])
	if payloadLengthErr != nil {
		return commands.PublishStream{}, payloadLengthErr
	}
	pos += payloadLengthOctetsUsed

	chunk := payload[pos:]
	if len(chunk) != int(payloadLength) {
		return commands.PublishStream{}, fmt.Errorf("Illegal length %d expected %d", len(chunk), int(payloadLength))
	}

	createdPublishStream := commands.NewPublishStream(channel, chunk)
	return createdPublishStream, nil
}
