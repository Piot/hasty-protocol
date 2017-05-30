package packetdeserializers

import (
	"errors"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/piot/hasty-protocol/commands"
	"github.com/piot/hasty-protocol/packet"
	"github.com/piot/hasty-protocol/unicodestring"
)

// ToLogin : Channel to subscribe to
func ToLogin(in packet.Packet) (commands.Login, error) {
	if in.Type() != packet.Login {
		return commands.Login{}, errors.New("Illegal packet type")
	}

	payload := in.Payload()
	pos := 0

	username, usernameOctetCountConsumed, usernameErr := unicodestring.FromOctets(payload[pos:])
	if usernameErr != nil {
		return commands.Login{}, fmt.Errorf("Illegal username %s", usernameErr)
	}

	password, _, passwordErr := unicodestring.FromOctets(payload[pos+usernameOctetCountConsumed:])
	if passwordErr != nil {
		return commands.Login{}, fmt.Errorf("Illegal password string %s", passwordErr)
	}

	login := commands.NewLogin(username, password)
	log.Infof("login %s", login)
	return login, nil
}
