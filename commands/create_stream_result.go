package commands

import (
	"fmt"

	"github.com/piot/hasty-protocol/channel"
	"github.com/piot/hasty-protocol/opath"
)

// CreateStreamResult : CreateStreamResult command
type CreateStreamResult struct {
	path    opath.OPath
	channel channel.ID
}

// NewCreateStreamResult : Creates a new CreateStreamResult command
func NewCreateStreamResult(path opath.OPath, channel channel.ID) CreateStreamResult {
	return CreateStreamResult{path: path, channel: channel}
}

// String : Returns a human readable string
func (in CreateStreamResult) String() string {
	return fmt.Sprintf("[createstreamresult %s %s]", in.path, in.channel)
}
