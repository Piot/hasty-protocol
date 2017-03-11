package commands

import "fmt"

// ConnectResult : ConnectResult command
type ConnectResult struct {
	protocolVersion Version
	realm           string
}

// NewConnectResult : Creates a new ConnectResult command
func NewConnectResult(realm string, protocolVersion Version) ConnectResult {
	return ConnectResult{protocolVersion: protocolVersion, realm: realm}
}

// String : Returns a human readable string
func (in ConnectResult) String() string {
	return fmt.Sprintf("[connectresult %s %s]", in.realm, in.protocolVersion)
}
