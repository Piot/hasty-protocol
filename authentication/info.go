package authentication

import (
	"fmt"

	"github.com/piot/hasty-protocol/user"
)

// Info : todo
type Info struct {
	userID user.ID
}

// NewInfo : todo
func NewInfo(userID user.ID) Info {
	return Info{userID: userID}
}

// UserID : todo
func (in Info) UserID() user.ID {
	return in.userID
}

// String : todo
func (in Info) String() string {
	return fmt.Sprintf("[authentication ID:%s ]", in.userID)
}
