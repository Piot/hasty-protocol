package packetserializers

import (
	"errors"
	"fmt"

	"github.com/piot/hasty-protocol/commands"
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

	fmt.Printf("numberOfSubscriptions:%d\n", numberOfSubscriptions)
	for i := uint8(0); i < numberOfSubscriptions; i++ {
		channelID, idOctets, err := ToChannelID(payload[pos : pos+4])
		if err != nil {
			return commands.SubscribeStream{}, err
		}
		fmt.Printf("channelID:%s\n", channelID)
		pos += idOctets

		qos, err2 := ToQos(payload[pos])
		if err2 != nil {
			return commands.SubscribeStream{}, err2
		}
		fmt.Printf("qos:%s\n", qos)
		pos++

		offset, offsetOctetsUsed, _ := ToOffset(payload[pos : pos+4])
		pos += offsetOctetsUsed
		info := commands.NewSubscribeStreamInfo(channelID, qos, offset)
		infos = append(infos, info)
	}
	createdSubscribe := commands.NewSubscribeStream(infos)
	fmt.Printf("subscribe %s\n", createdSubscribe)
	return createdSubscribe, nil
}
