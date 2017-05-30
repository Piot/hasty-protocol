package packetdeserializers

import (
	"encoding/binary"
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
	pos := 0
	requestID := binary.BigEndian.Uint64(payload[pos : pos+8])
	pos += 8
	opath, _, err := ToOpath(payload[pos:])
	if err != nil {
		return commands.CreateStream{}, err
	}

	createStream := commands.NewCreateStream(requestID, opath)
	return createStream, nil
}
