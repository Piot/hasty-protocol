package packetserializers

import (
	"github.com/piot/hasty-protocol/channel"
	"github.com/piot/hasty-protocol/commands"
	"github.com/piot/hasty-protocol/packet"
)

// PacketHandler : handle commands
type PacketHandler interface {
	HandleConnect(commands.Connect) error
	HandlePublishStream(commands.PublishStream) error
	HandleSubscribeStream(commands.SubscribeStream)
	HandleUnsubscribeStream(commands.UnsubscribeStream)
	HandleCreateStream(commands.CreateStream) (channel.ID, error)
	HandleStreamData(commands.StreamData)
}

// Deserialize : ss
func Deserialize(in packet.Packet, handler PacketHandler) (err error) {
	switch in.Type() {
	case packet.CreateStream:
		createStream, err := ToCreateStream(in)
		if err != nil {
			return err
		}
		handler.HandleCreateStream(createStream)
	case packet.PublishStream:
		publishStream, err := ToPublishStream(in)
		if err != nil {
			return err
		}
		handler.HandlePublishStream(publishStream)
	case packet.SubscribeStream:
		subscribe, err := ToSubscribeStream(in)
		if err != nil {
			return err
		}
		handler.HandleSubscribeStream(subscribe)
	case packet.UnsubscribeStream:
		unsubscribe, err := ToUnsubscribeStream(in)
		if err != nil {
			return err
		}
		handler.HandleUnsubscribeStream(unsubscribe)
	case packet.StreamData:
		streamdata, err := ToStreamData(in)
		if err != nil {
			return err
		}
		handler.HandleStreamData(streamdata)
	case packet.Connect:
		connect, err := ToConnect(in)
		if err != nil {
			return err
		}
		handler.HandleConnect(connect)
	}

	return nil
}
