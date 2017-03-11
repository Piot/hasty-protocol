package commands

import (
	"fmt"

	"github.com/piot/hasty-protocol/channel"
)

// PublishStream : publish path
type PublishStream struct {
	channel channel.ID
	chunk   []byte
}

// NewPublishStream : Creates a publish
func NewPublishStream(channel channel.ID, chunk []byte) PublishStream {
	return PublishStream{channel: channel, chunk: chunk}
}

// String : Return human readable
func (in PublishStream) String() string {
	return fmt.Sprintf("[publishstream %s %s]", in.channel, in.chunk)
}

// Channel : string
func (in PublishStream) Channel() channel.ID {
	return in.channel
}

// Chunk : octets
func (in PublishStream) Chunk() []byte {
	return in.chunk
}
