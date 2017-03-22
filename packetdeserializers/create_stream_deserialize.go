package packetdeserializers

import (
	"errors"

	"github.com/piot/hasty-protocol/commands"
	"github.com/piot/hasty-protocol/packet"
)

// ToCreateStream : Path to create
func ToCreateStream(in packet.Packet) (commands.CreateStream, error) {
	if in.Type() != packet.CreateStream {
		return commands.CreateStream{}, errors.New("Illegal packet type")
	}

	payload := in.Payload()
	opath, _, err := ToOpath(payload)
	if err != nil {
		return commands.CreateStream{}, err
	}

	createStream := commands.NewCreateStream(opath)
	return createStream, nil
}
