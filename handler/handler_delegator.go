package handler

import (
	"fmt"

	"github.com/piot/hasty-protocol/channel"
	"github.com/piot/hasty-protocol/commands"
	"github.com/piot/hasty-protocol/packetserializers"
)

// PacketHandler : handle commands
type PacketHandlerDelegator struct {
	handlers []packetserializers.PacketHandler
}

func NewPacketHandlerDelegator() PacketHandlerDelegator {
	return PacketHandlerDelegator{}
}

func (in *PacketHandlerDelegator) AddHandler(v packetserializers.PacketHandler) {
	in.handlers = append(in.handlers, v)
}

func (in *PacketHandlerDelegator) HandleConnect(cmd commands.Connect) error {
	fmt.Printf("Delegator Connect")
	for _, v := range in.handlers {
		v.HandleConnect(cmd)
	}
	return nil
}

func (in *PacketHandlerDelegator) HandlePublishStream(cmd commands.PublishStream) error {
	for _, v := range in.handlers {
		err := v.HandlePublishStream(cmd)
		if err != nil {
			return err
		}

	}
	return nil
}

func (in *PacketHandlerDelegator) HandleSubscribeStream(cmd commands.SubscribeStream) {
	for _, v := range in.handlers {
		v.HandleSubscribeStream(cmd)
	}
}

func (in *PacketHandlerDelegator) HandleUnsubscribeStream(cmd commands.UnsubscribeStream) {
	for _, v := range in.handlers {
		v.HandleUnsubscribeStream(cmd)
	}
}

func (in *PacketHandlerDelegator) HandleCreateStream(cmd commands.CreateStream) (channel.ID, error) {
	for _, v := range in.handlers {
		c, err := v.HandleCreateStream(cmd)
		if err != nil {
			return c, err
		}

	}
	return channel.ID{}, nil
}

func (in *PacketHandlerDelegator) HandleStreamData(cmd commands.StreamData) {
	for _, v := range in.handlers {
		v.HandleStreamData(cmd)
	}
}
