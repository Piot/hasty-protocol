package commands

import (
	"fmt"

	"github.com/piot/hasty-protocol/channel"
	"github.com/piot/hasty-protocol/user"
)

// Authenticated : Authenticated command
type Authenticated struct {
	userID                 user.ID
	userAllocatedChannelID channel.ID
	realname               string
}

// NewAuthenticated : Creates a new Authenticated command
func NewAuthenticated(userID user.ID, userAllocatedChannelID channel.ID, realname string) Authenticated {
	return Authenticated{userID: userID, userAllocatedChannelID: userAllocatedChannelID, realname: realname}
}

// String : Returns a human readable string
func (in Authenticated) String() string {
	return fmt.Sprintf("[authenticated_cmd id:%s realname:%s user_channel:%s]", in.userID, in.realname, in.userAllocatedChannelID)
}

// Realname : todo
func (in Authenticated) Realname() string {
	return in.realname
}

// UserID : todo
func (in Authenticated) UserID() user.ID {
	return in.userID
}

// UserID : todo
func (in Authenticated) UserAllocatedChannelID() channel.ID {
	return in.userAllocatedChannelID
}
