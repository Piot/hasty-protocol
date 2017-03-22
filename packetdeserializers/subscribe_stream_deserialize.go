package packetdeserializers

import (
	"errors"

	"github.com/piot/hasty-protocol/commands"
	"github.com/piot/hasty-protocol/deserialize"
	"github.com/piot/hasty-protocol/packet"
)

// ToSubscribeStream : Channel to subscribe to
func ToSubscribeStream(in packet.Packet) (commands.SubscribeStream, error) {
	if in.Type() != packet.SubscribeStream {
		return commands.SubscribeStream{}, errors.New("Illegal packet type")
	}

	payload := in.Payload()
	numberOfSubscriptions := payload[0]
	pos := 1
	var infos []commands.SubscribeStreamInfo

	for i := uint8(0); i < numberOfSubscriptions; i++ {
		channelID, idOctets, err := ToChannelID(payload[pos : pos+4])
		if err != nil {
			return commands.SubscribeStream{}, err
		}
		pos += idOctets

		qos, err2 := deserialize.ToQos(payload[pos])
		if err2 != nil {
			return commands.SubscribeStream{}, err2
		}
		pos++

		offset, offsetOctetsUsed, _ := deserialize.ToOffset(payload[pos : pos+4])
		pos += offsetOctetsUsed
		info := commands.NewSubscribeStreamInfo(channelID, qos, offset)
		infos = append(infos, info)
	}
	createdSubscribe := commands.NewSubscribeStream(infos)
	return createdSubscribe, nil
}
