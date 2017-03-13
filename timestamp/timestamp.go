package timestamp

import (
	"fmt"
	"time"
)

// Time : Time in milliseconds UTC
type Time struct {
	ms uint64
}

const millisNanosFactor = 1000000

// New : Create Version
func New(ms uint64) (out Time) {
	t := Time{ms: ms}
	return t
}

// String : return string representation
func (in Time) String() string {
	nanos := int64(in.ms * millisNanosFactor)
	return fmt.Sprintf("[time %s]", time.Unix(0, nanos))
}

// Raw : return raw representation
func (in Time) Raw() uint64 {
	return in.ms
}

// Now : Current time
func Now() Time {
	now := time.Now()
	nanos := now.UnixNano()
	millis := uint64(nanos / millisNanosFactor)
	return New(millis)
}
