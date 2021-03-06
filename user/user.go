package user

import "fmt"

// ID : todo
type ID struct {
	value uint64
}

// NewID : todo
func NewID(value uint64) (ID, error) {
	return ID{value: value}, nil
}

// String : Returns human readable string
func (in ID) String() string {
	return fmt.Sprintf("[userid %d]", in.value)
}

// Raw : todo
func (in ID) Raw() uint64 {
	return in.value
}

// ToHex : Returns hex string
func (in ID) ToHex() string {
	return fmt.Sprintf("%08x", in.value)
}

// ToName : Returns name string
func (in ID) ToName() string {
	return fmt.Sprintf("%d", in.value)
}
