package packet

import "fmt"

// ConnectionID : todo
type ConnectionID struct {
	id uint
}

// NewConnectionID : Create connectionId
func NewConnectionID(id uint) ConnectionID {
	return ConnectionID{id: id}
}

// String : Human readable string
func (in ConnectionID) String() string {
	return fmt.Sprintf("[connection %d]", in.id)
}

// Raw : todo
func (in ConnectionID) Raw() uint {
	return in.id
}
