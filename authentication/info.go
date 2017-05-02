package authentication

import (
	"github.com/piot/hasty-protocol/channel"
	"github.com/piot/hasty-protocol/user"
)

// Info : todo
type Info struct {
	userID    user.ID
	channelID channel.ID
}

// NewInfo : todo
func NewInfo(userID user.ID, channelID channel.ID) Info {
	return Info{userID: userID, channelID: channelID}
}

// UserID : todo
func (in Info) UserID() user.ID {
	return in.userID
}

// UserAllocatedChannelID : todo
func (in Info) UserAllocatedChannelID() channel.ID {
	return in.channelID
}
