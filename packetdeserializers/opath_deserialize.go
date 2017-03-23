package packetdeserializers

import (
	"github.com/piot/hasty-protocol/asciistring"
	"github.com/piot/hasty-protocol/opath"
)

// ToOpath : creating opath
func ToOpath(octets []byte) (opath.OPath, int, error) {
	str, length, errFrom := asciistring.FromOctets(octets)
	if errFrom != nil {
		return opath.OPath{}, 0, errFrom
	}

	path, pathError := opath.NewFromString(str)
	return path, length, pathError
}
