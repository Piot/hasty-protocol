package commands

import (
	"fmt"

	"github.com/piot/hasty-protocol/version"
)

// Connect : Connect command
type Connect struct {
	protocolVersion version.Version
	realm           string
}

// NewConnect : Creates a new Connect command
func NewConnect(realm string, protocolVersion version.Version) Connect {
	return Connect{protocolVersion: protocolVersion, realm: realm}
}

// String : Returns a human readable string
func (in Connect) String() string {
	return fmt.Sprintf("[connect realm:%s %s]", in.realm, in.protocolVersion)
}

// Realm : todo
func (in Connect) Realm() string {
	return in.realm
}

// ProtocolVersion : todo
func (in Connect) ProtocolVersion() version.Version {
	return in.protocolVersion
}
