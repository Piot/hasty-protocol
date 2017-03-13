package commands

import (
	"fmt"

	"github.com/piot/hasty-protocol/timestamp"
)

// Ping : Ping command
type Ping struct {
	sentTime timestamp.Time
}

// NewPing : Creates a new Ping command
func NewPing(sentTime timestamp.Time) Ping {
	return Ping{sentTime: sentTime}
}

// String : Returns a human readable string
func (in Ping) String() string {
	return fmt.Sprintf("[ping was sent:%s]", in.sentTime)
}

// SentTime : Current time it was sent
func (in Ping) SentTime() timestamp.Time {
	return in.sentTime
}
