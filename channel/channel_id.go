package channel

import "fmt"

// ID : The id of the channel
type ID struct {
	id uint32
}

// NewFromID : Create ChannelId
func NewFromID(in uint32) (out ID, err error) {
	id := ID{id: in}
	return id, nil
}

// Raw : raw value
func (in ID) Raw() uint32 {
	return in.id
}

// String : return string representation
func (in ID) String() string {
	return fmt.Sprintf("[channel %08X]", in.id)
}

// ToHex : Returns hex string
func (in ID) ToHex() string {
	return fmt.Sprintf("%08x", in.id)
}
