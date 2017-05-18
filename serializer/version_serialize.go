package serializer

import "github.com/piot/hasty-protocol/version"

// VersionToOctets : convert to octets
func VersionToOctets(ver version.Version) (octets []byte) {
	octets = make([]byte, 3)
	octets[0] = ver.Major()
	octets[1] = ver.Minor()
	octets[2] = ver.Patch()
	return octets
}
