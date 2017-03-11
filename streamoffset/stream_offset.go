package streamoffset

import "fmt"

// Offset : The id of the channel
type Offset struct {
	position uint32
}

// NewFromOffset : Create ChannelId
func NewFromOffset(position uint32) (out Offset, err error) {
	offset := Offset{position: position}
	return offset, nil
}

// Raw : raw value
func (in Offset) Raw() uint32 {
	return in.position
}

// String : return string representation
func (in Offset) String() string {
	return fmt.Sprintf("[streamoffset %08X]", in.position)
}

// ToHex : Returns hex string
func (in Offset) ToHex() string {
	return fmt.Sprintf("%08x", in.position)
}
