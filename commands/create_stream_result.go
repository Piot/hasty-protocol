package commands

import (
	"fmt"

	"github.com/piot/hasty-protocol/channel"
)

// CreateStreamResult : CreateStreamResult command
type CreateStreamResult struct {
	channel   channel.ID
	requestID uint64
}

// NewCreateStreamResult : Creates a new CreateStreamResult command
func NewCreateStreamResult(requestID uint64, channel channel.ID) CreateStreamResult {
	return CreateStreamResult{requestID: requestID, channel: channel}
}

// String : Returns a human readable string
func (in CreateStreamResult) String() string {
	return fmt.Sprintf("[createstreamresult %d %s]", in.requestID, in.channel)
}

// RequestID : todo
func (in CreateStreamResult) RequestID() uint64 {
	return in.requestID
}

func (in CreateStreamResult) ChannelID() channel.ID {
	return in.channel
}
