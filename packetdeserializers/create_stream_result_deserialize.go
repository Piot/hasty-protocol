package packetdeserializers

import (
	"encoding/binary"
	"errors"

	"github.com/piot/hasty-protocol/channel"
	"github.com/piot/hasty-protocol/commands"
	"github.com/piot/hasty-protocol/packet"
)

// ToCreateStreamResult : Path to create
func ToCreateStreamResult(in packet.Packet) (commands.CreateStreamResult, error) {
	if in.Type() != packet.CreateStreamResult {
		return commands.CreateStreamResult{}, errors.New("Illegal packet type")
	}

	payload := in.Payload()
	pos := 0
	requestID := binary.BigEndian.Uint64(payload[pos : pos+8])
	pos += 8
	channelIDValue := binary.BigEndian.Uint32(payload[pos : pos+4])

	channelID, _ := channel.NewFromID(channelIDValue)
	createStreamResult := commands.NewCreateStreamResult(requestID, channelID)
	return createStreamResult, nil
}
