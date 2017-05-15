package handler

import (
	"github.com/piot/hasty-protocol/channel"
	"github.com/piot/hasty-protocol/commands"
)

// PacketHandler : handle commands
type PacketHandlerDelegator struct {
	handlers []PacketHandler
}

func NewPacketHandlerDelegator() PacketHandlerDelegator {
	return PacketHandlerDelegator{}
}

func (in *PacketHandlerDelegator) AddHandler(v PacketHandler) {
	in.handlers = append(in.handlers, v)
}

func (in *PacketHandlerDelegator) HandleConnect(cmd commands.Connect) error {
	for _, v := range in.handlers {
		v.HandleConnect(cmd)
	}
	return nil
}

func (in *PacketHandlerDelegator) HandlePublishStreamUser(cmd commands.PublishStreamUser) error {
	for _, v := range in.handlers {
		err := v.HandlePublishStreamUser(cmd)
		if err != nil {
			return err
		}

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

func (in *PacketHandlerDelegator) HandlePing(cmd commands.Ping) {
	for _, v := range in.handlers {
		v.HandlePing(cmd)
	}
}

func (in *PacketHandlerDelegator) HandlePong(cmd commands.Pong) {
	for _, v := range in.handlers {
		v.HandlePong(cmd)
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

func (in *PacketHandlerDelegator) HandleLogin(cmd commands.Login) error {
	for _, v := range in.handlers {
		err := v.HandleLogin(cmd)
		if err != nil {
			return err
		}

	}
	return nil
}

func (in *PacketHandlerDelegator) HandleAuthenticated(cmd commands.Authenticated) {
	for _, v := range in.handlers {
		v.HandleAuthenticated(cmd)
	}
}

func (in *PacketHandlerDelegator) HandleTransportDisconnect() {
	for _, v := range in.handlers {
		v.HandleTransportDisconnect()
	}
}
