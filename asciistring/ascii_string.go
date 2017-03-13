package asciistring

import "errors"

// FromOctets : Returns string from octets
func FromOctets(octets []byte) (string, int, error) {
	start := 0
	octetLength := len(octets)
	if octetLength < 1 {
		return "", 0, errors.New("too small buffer")
	}
	length := int(octets[start])
	start++
	if length > 127 {
		if octetLength < 2 {
			return "", 0, errors.New("too small buffer")
		}
		length += int(octets[start])
		start++
	}

	if octetLength < start+length {
		return "", 0, errors.New("too small buffer")
	}

	s := string(octets[start : start+length])
	return s, length + start, nil
}

// ToOctets : Converts an ascii string to octets
func ToOctets(str string) ([]byte, error) {
	strLen := len(str)
	encodeLengthSize := 1
	if strLen > 127 {
		encodeLengthSize = 2
	}
	target := make([]byte, encodeLengthSize+strLen)
	target[0] = byte(strLen)
	copy(target[encodeLengthSize:], []byte(str))
	return target, nil
}
