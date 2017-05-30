package packetdeserializers

import (
	"encoding/binary"
	"encoding/hex"
	"fmt"

	log "github.com/sirupsen/logrus"

	"github.com/piot/hasty-protocol/channel"
	"github.com/piot/hasty-protocol/commands"
	"github.com/piot/hasty-protocol/packet"
	"github.com/piot/hasty-protocol/unicodestring"
	"github.com/piot/hasty-protocol/user"
)

// OctetsToAuthenticated : todo
func OctetsToAuthenticated(octets []byte) (user.ID, channel.ID, string, int, error) {
	pos := 0
	userIDValue := binary.BigEndian.Uint64(octets[pos : pos+8])
	userID, _ := user.NewID(userIDValue)
	pos += 8

	channelIDValue := binary.BigEndian.Uint32(octets[pos : pos+4])
	channelID, _ := channel.NewFromID(channelIDValue)
	pos += 4

	realname, _, realnameErr := unicodestring.FromOctets(octets[pos:])
	if realnameErr != nil {
		return user.ID{}, channel.ID{}, "", 0, fmt.Errorf("Illegal password string %s", realnameErr)
	}
	return userID, channelID, realname, pos, nil
}

// ToAuthenticated : todo
func ToAuthenticated(in packet.Packet) (commands.Authenticated, error) {
	return ToAuthenticatedEx(in.Payload())
}

// ToAuthenticatedEx : todo
func ToAuthenticatedEx(payload []byte) (commands.Authenticated, error) {
	log.Debugf("Before decode:%s", hex.EncodeToString(payload))
	userID, userAllocatedChannelID, realname, _, infoErr := OctetsToAuthenticated(payload)
	if infoErr != nil {
		return commands.Authenticated{}, infoErr
	}

	authenticatedCommand := commands.NewAuthenticated(userID, userAllocatedChannelID, realname)

	return authenticatedCommand, nil
}
