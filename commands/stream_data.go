package commands

import (
	"encoding/hex"
	"fmt"

	"github.com/piot/hasty-protocol/channel"
	"github.com/piot/hasty-protocol/streamoffset"
)

// StreamData : data inside a stream
type StreamData struct {
	channel         channel.ID
	offset          streamoffset.Offset
	isAtEndPosition bool
	data            []byte
}

// NewStreamData : Creates a publish
func NewStreamData(channel channel.ID, offset streamoffset.Offset, data []byte, isAtEndPosition bool) StreamData {
	return StreamData{channel: channel, offset: offset, data: data, isAtEndPosition: isAtEndPosition}
}

// String : Return human readable
func (in StreamData) String() string {
	hexPayload := hex.Dump(in.data)
	return fmt.Sprintf("[streamdata %s %s %s]", in.channel, in.offset, hexPayload)
}

// Channel : string
func (in StreamData) Channel() channel.ID {
	return in.channel
}

// Offset : Where this data starts
func (in StreamData) Offset() streamoffset.Offset {
	return in.offset
}

// Data : octets
func (in StreamData) Data() []byte {
	return in.data
}
