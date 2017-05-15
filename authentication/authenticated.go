package authentication

import (
	"fmt"

	"github.com/piot/hasty-protocol/channel"
	"github.com/piot/hasty-protocol/user"
)

// Authenticated : todo
type Authenticated struct {
	userID                 user.ID
	userAllocatedChannelID channel.ID
	realname               string
}

// NewAuthenticated : todo
func NewAuthenticated(userID user.ID, userAllocatedChannelID channel.ID, realname string) Authenticated {
	return Authenticated{userID: userID, userAllocatedChannelID: userAllocatedChannelID, realname: realname}
}

// UserID : todo
func (in Authenticated) UserID() user.ID {
	return in.userID
}

// UserAllocatedChannelID : todo
func (in Authenticated) UserAllocatedChannelID() channel.ID {
	return in.userAllocatedChannelID
}

// Realname : todo
func (in Authenticated) Realname() string {
	return in.realname
}

// String : todo
func (in Authenticated) String() string {
	return fmt.Sprintf("[authenticated ID:%s ]", in.userID)
}
