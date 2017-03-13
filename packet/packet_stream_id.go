package packet

import "fmt"

type ConnectionID struct {
	id uint
}

// New : Create connectionId
func NewConnectionID(id uint) ConnectionID {
	return ConnectionID{id: id}
}

// String : Human readable string
func (in ConnectionID) String() string {
	return fmt.Sprintf("[connection %d]", in.id)
}
