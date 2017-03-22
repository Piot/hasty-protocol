package packetdeserializers

import (
	"errors"
	"fmt"

	"github.com/piot/hasty-protocol/commands"
	"github.com/piot/hasty-protocol/deserialize"
	"github.com/piot/hasty-protocol/packet"
	"github.com/piot/shadow-broker/asciistring"
)

// ToConnect : Channel to subscribe to
func ToConnect(in packet.Packet) (commands.Connect, error) {
	if in.Type() != packet.Connect {
		return commands.Connect{}, errors.New("Illegal packet type")
	}

	payload := in.Payload()
	pos := 0
	versionObject, versionOctetCount, versionErr := deserialize.ToVersion(payload[0:3])
	if versionErr != nil {
		return commands.Connect{}, errors.New("Illegal version")
	}
	pos += versionOctetCount
	realm, _, realmErr := asciistring.FromOctets(payload[pos:])
	if realmErr != nil {
		return commands.Connect{}, fmt.Errorf("Illegal realm %s", realmErr)
	}
	connect := commands.NewConnect(realm, versionObject)
	fmt.Printf("connect %s\n", connect)
	return connect, nil
}
