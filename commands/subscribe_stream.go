package commands

import (
	"fmt"

	"github.com/piot/hasty-protocol/channel"
	"github.com/piot/hasty-protocol/qos"
)

// SubscribeStreamInfo : SubscribeStreamInfo command
type SubscribeStreamInfo struct {
	id     channel.ID
	qos    qos.QoS
	offset uint32
}

// Channel : todo
func (in SubscribeStreamInfo) Channel() channel.ID {
	return in.id
}

// QoS : todo
func (in SubscribeStreamInfo) QoS() qos.QoS {
	return in.qos
}

func (in SubscribeStreamInfo) Offset() uint32 {
	return in.offset
}

// String : Return human readable string
func (in SubscribeStreamInfo) String() string {
	return fmt.Sprintf("[subinfo %s %s]", in.id, in.qos)
}

// NewSubscribeStreamInfo : create subscribe info
func NewSubscribeStreamInfo(channelID channel.ID, qos qos.QoS, offset uint32) SubscribeStreamInfo {
	return SubscribeStreamInfo{id: channelID, qos: qos, offset: offset}
}

// SubscribeStream : SubscribeStream command
type SubscribeStream struct {
	infos []SubscribeStreamInfo
}

// NewSubscribeStream : Creates a new ubscribeStream command
func NewSubscribeStream(infos []SubscribeStreamInfo) SubscribeStream {
	return SubscribeStream{infos: infos}
}

// Infos : todo
func (in SubscribeStream) Infos() []SubscribeStreamInfo {
	return in.infos
}

// String : Return human readable string
func (in SubscribeStream) String() string {
	return fmt.Sprintf("[subscribestream %s]", in.infos)
}
