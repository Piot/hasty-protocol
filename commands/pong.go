package commands

import (
	"fmt"

	"github.com/piot/hasty-protocol/timestamp"
)

// Pong : Ping command
type Pong struct {
	sentTime   timestamp.Time
	echoedTime timestamp.Time
}

// NewPong : Creates a new Pong command
func NewPong(sentTime timestamp.Time, echoedTime timestamp.Time) Pong {
	return Pong{sentTime: sentTime, echoedTime: echoedTime}
}

// String : Returns a human readable string
func (in Pong) String() string {
	return fmt.Sprintf("[pong was sent:%s echoed:%s]", in.sentTime, in.echoedTime)
}

// SentTime : SentTime
func (in Pong) SentTime() timestamp.Time {
	return in.sentTime
}

// EchoedTime : Echoed Time from remote
func (in Pong) EchoedTime() timestamp.Time {
	return in.echoedTime
}
