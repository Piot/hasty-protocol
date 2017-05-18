package commands

import (
	"fmt"

	"github.com/piot/hasty-protocol/realmname"
	"github.com/piot/hasty-protocol/version"
)

// Connect : Connect command
type Connect struct {
	protocolVersion version.Version
	realm           realmname.Name
}

// NewConnect : Creates a new Connect command
func NewConnect(realm realmname.Name, protocolVersion version.Version) Connect {
	return Connect{protocolVersion: protocolVersion, realm: realm}
}

// String : Returns a human readable string
func (in Connect) String() string {
	return fmt.Sprintf("[connect realm:%s %s]", in.realm, in.protocolVersion)
}

// Realm : todo
func (in Connect) Realm() realmname.Name {
	return in.realm
}

// ProtocolVersion : todo
func (in Connect) ProtocolVersion() version.Version {
	return in.protocolVersion
}
