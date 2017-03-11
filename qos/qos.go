package qos

import "fmt"

// Enum : The QoS enum
type Enum int

// Nop : The different types
const (
	Normal Enum = 1 + iota
	Prioritized
)

// QoS : The qos for a subscribe
type QoS struct {
	v Enum
}

func NewQoS(v Enum) QoS {
	return QoS{v: v}
}

// String : return string representation
func (in QoS) String() string {
	return fmt.Sprintf("[qos %v]", in.v)
}
