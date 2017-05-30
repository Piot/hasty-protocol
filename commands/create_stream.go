package commands

import (
	"fmt"

	"github.com/piot/hasty-protocol/opath"
)

// CreateStream : CreateStream command
type CreateStream struct {
	path          opath.OPath
	requestNumber uint64
}

// NewCreateStream : Creates a new CreateStream command
func NewCreateStream(requestNumber uint64, path opath.OPath) CreateStream {
	return CreateStream{requestNumber: requestNumber, path: path}
}

// String : Returns a human readable string
func (in CreateStream) String() string {
	return fmt.Sprintf("[createstream %08X : %s]", in.requestNumber, in.path)
}

// Path : todo
func (in CreateStream) Path() opath.OPath {
	return in.path
}

// RequestNumber : todo
func (in CreateStream) RequestNumber() uint64 {
	return in.requestNumber
}
