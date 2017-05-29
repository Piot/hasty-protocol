package packetdeserializers

import (
	"log"

	"github.com/piot/hasty-protocol/handler"
	"github.com/piot/hasty-protocol/packet"
)

// Deserialize : ss
func Deserialize(in packet.Packet, handler handler.PacketHandler) (err error) {
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
	case packet.PublishStreamUser:
		publishStream, err := ToPublishStreamUser(in)
		if err != nil {
			return err
		}
		handler.HandlePublishStreamUser(publishStream)
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
	case packet.Authenticated:
		authenticated, err := ToAuthenticated(in)
		if err != nil {
			return err
		}
		handler.HandleAuthenticated(authenticated)
	case packet.Connect:
		connect, err := ToConnect(in)
		if err != nil {
			return err
		}
		handler.HandleConnect(connect)
	case packet.Ping:
		ping, err := ToPing(in)
		if err != nil {
			return err
		}
		handler.HandlePing(ping)
	case packet.Pong:
		pong, err := ToPong(in)
		if err != nil {
			return err
		}
		handler.HandlePong(pong)
	case packet.Login:
		login, err := ToLogin(in)
		if err != nil {
			log.Printf("Login err: %s", err)
			return err
		}
		loginErr := handler.HandleLogin(login)
		if loginErr != nil {
			log.Printf("Authentication err: %s", loginErr)
			return loginErr
		}
	}

	return nil
}
