package chunk

import (
	"encoding/hex"
	"fmt"
)

// Chunk :
type Chunk struct {
	header  Header
	payload []byte
}

// NewChunk : Creates a new chunk
func NewChunk(header Header, payload []byte) Chunk {
	return Chunk{header: header, payload: payload}
}

// Payload : Returns the actual octets
func (in Chunk) Payload() []byte {
	return in.payload
}

// String : A human readable representation of a chunk
func (in Chunk) String() string {
	hexPayload := hex.Dump(in.Payload())
	return fmt.Sprintf("[Chunk %s %s]", in.header, hexPayload)
}
