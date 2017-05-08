package packetdeserializers

import (
	"errors"
	"fmt"

	"github.com/piot/hasty-protocol/commands"
	"github.com/piot/hasty-protocol/deserialize"
	"github.com/piot/hasty-protocol/packet"
)

// ToPublishStreamUser : Packet to publish
func ToPublishStreamUser(in packet.Packet) (commands.PublishStreamUser, error) {
	if in.Type() != packet.PublishStreamUser {
		return commands.PublishStreamUser{}, errors.New("Illegal packet type")
	}

	payload := in.Payload()
	pos := 0
	userID, idOctets, err := ToUserID(payload[pos:])
	if err != nil {
		return commands.PublishStreamUser{}, err
	}
	pos += idOctets

	payloadLength, payloadLengthOctetsUsed, payloadLengthErr := deserialize.ToSmallLength(payload[pos:])
	if payloadLengthErr != nil {
		return commands.PublishStreamUser{}, payloadLengthErr
	}
	pos += payloadLengthOctetsUsed

	chunk := payload[pos:]
	if len(chunk) != int(payloadLength) {
		return commands.PublishStreamUser{}, fmt.Errorf("Illegal length %d expected %d", len(chunk), int(payloadLength))
	}

	createdPublishStreamUser := commands.NewPublishStreamUser(userID, chunk)
	return createdPublishStreamUser, nil
}
