package deserialize

import "fmt"

func ToSmallLength(buf []byte) (uint16, int, error) {
	if len(buf) == 0 {
		return 0, 0, fmt.Errorf("Buffer too small")
	}
	packetLength := uint16(buf[0])
	octetsUsed := 1

	if packetLength > 127 {
		if len(buf) < 2 {
			return 0, 0, fmt.Errorf("Buffer too small")
		}
		packetLength = uint16(buf[0]&0x7f)*0x100 + uint16(buf[1])
		octetsUsed++
	}

	return packetLength, octetsUsed, nil
}
