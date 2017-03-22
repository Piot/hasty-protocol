package packetdeserializers

import (
	"errors"

	"github.com/piot/hasty-protocol/commands"
	"github.com/piot/hasty-protocol/packet"
)

// ToPublishStream : Packet to publish
func ToPublishStream(in packet.Packet) (commands.PublishStream, error) {
	if in.Type() != packet.PublishStream {
		return commands.PublishStream{}, errors.New("Illegal packet type")
	}

	payload := in.Payload()
	channel, idOctets, err := ToChannelID(payload)
	if err != nil {
		return commands.PublishStream{}, err
	}

	chunk := payload[idOctets:]

	createdPublishStream := commands.NewPublishStream(channel, chunk)
	return createdPublishStream, nil
}
