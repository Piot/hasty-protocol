package asciistring

import (
	"errors"

	"github.com/piot/hasty-protocol/deserialize"
	"github.com/piot/hasty-protocol/serializer"
)

// FromOctets : Returns string from octets
func FromOctets(octets []byte) (string, int, error) {
	pos := 0
	octetLength := len(octets)
	if octetLength < 1 {
		return "", 0, errors.New("too small buffer")
	}
	length, octetsUsedForLength, lengthErr := deserialize.ToSmallLength(octets)
	if lengthErr != nil {
		return "", 0, lengthErr
	}
	intLength := int(length)
	pos += octetsUsedForLength

	if octetLength < pos+intLength {
		return "", 0, errors.New("too small buffer")
	}

	s := string(octets[pos : pos+intLength])
	return s, intLength + pos, nil
}

// ToOctets : Converts an ascii string to octets
func ToOctets(str string) ([]byte, error) {
	strLen := len(str)
	lengthOctets, lengthErr := serializer.SmallLengthToOctets(uint16(strLen))
	if lengthErr != nil {
		return nil, lengthErr
	}
	encodeLengthSize := len(lengthOctets)
	target := make([]byte, encodeLengthSize+strLen)
	copy(target, lengthOctets)
	copy(target[encodeLengthSize:], []byte(str))
	return target, nil
}
