package handler

// PacketHandler : handle commands
import (
	"github.com/piot/hasty-protocol/channel"
	"github.com/piot/hasty-protocol/commands"
)

type PacketHandler interface {
	HandleConnect(commands.Connect) error
	HandlePing(commands.Ping)
	HandlePong(commands.Pong)
	HandlePublishStream(commands.PublishStream) error
	HandlePublishStreamUser(commands.PublishStreamUser) error
	HandleSubscribeStream(commands.SubscribeStream)
	HandleUnsubscribeStream(commands.UnsubscribeStream)
	HandleCreateStream(commands.CreateStream) (channel.ID, error)
	HandleCreateStreamResult(commands.CreateStreamResult) error
	HandleStreamData(commands.StreamData)
	HandleAuthenticated(commands.Authenticated)
	HandleLogin(commands.Login) error
	HandleTransportDisconnect()
}
