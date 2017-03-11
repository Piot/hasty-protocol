package commands

import (
	"fmt"

	"github.com/piot/hasty-protocol/channel"
)

// UnsubscribeStream : UnsubscribeStream command
type UnsubscribeStream struct {
	channel channel.ID
}

// NewUnsubscribeStream : Create UnsubscribeStream command
func NewUnsubscribeStream(channel channel.ID) UnsubscribeStream {
	return UnsubscribeStream{channel: channel}
}

// Channel : get channel id
func (in UnsubscribeStream) Channel() channel.ID {
	return in.channel
}

// String : Return human readable string
func (in UnsubscribeStream) String() string {
	return fmt.Sprintf("[unsubscribestream %s]", in.channel)
}
