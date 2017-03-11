package packetserializers

import (
	"errors"

	"github.com/piot/hasty-protocol/commands"
	"github.com/piot/hasty-protocol/packet"
)

// ToUnsubscribeStream : packet to Unsubscribe
func ToUnsubscribeStream(in packet.Packet) (commands.UnsubscribeStream, error) {
	if in.Type() != packet.UnsubscribeStream {
		return commands.UnsubscribeStream{}, errors.New("Illegal packet type")
	}

	payload := in.Payload()
	channel, _, err := ToChannelID(payload)
	if err != nil {
		return commands.UnsubscribeStream{}, err
	}

	createdUnsubscribe := commands.NewUnsubscribeStream(channel)
	return createdUnsubscribe, nil
}
