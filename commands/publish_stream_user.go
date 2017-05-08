package commands

import (
	"fmt"

	"github.com/piot/hasty-protocol/user"
)

// PublishStreamUser : publish path
type PublishStreamUser struct {
	user  user.ID
	chunk []byte
}

// NewPublishStreamUser : Creates a publish
func NewPublishStreamUser(user user.ID, chunk []byte) PublishStreamUser {
	return PublishStreamUser{user: user, chunk: chunk}
}

// String : Return human readable
func (in PublishStreamUser) String() string {
	return fmt.Sprintf("[publishstreamuser %s %s]", in.user, in.chunk)
}

// User : string
func (in PublishStreamUser) User() user.ID {
	return in.user
}

// Chunk : octets
func (in PublishStreamUser) Chunk() []byte {
	return in.chunk
}
