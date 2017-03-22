package packetdeserializers

import (
	"fmt"

	"github.com/piot/hasty-protocol/asciistring"
	"github.com/piot/hasty-protocol/opath"
)

// ToOpath : creating opath
func ToOpath(octets []byte) (opath.OPath, int, error) {
	str, length, errFrom := asciistring.FromOctets(octets)
	if errFrom != nil {
		return opath.OPath{}, 0, errFrom
	}

	fmt.Printf("OPATH:'%s'\n", str)
	path, pathError := opath.NewFromString(str)
	return path, length, pathError
}
