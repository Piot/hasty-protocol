package deserialize

import "github.com/piot/hasty-protocol/version"

// ToVersion : convert from octets
func ToVersion(octets []byte) (version.Version, int, error) {
	major := octets[0]
	minor := octets[1]
	patch := octets[2]
	version, err := version.New(major, minor, patch)
	return version, 3, err
}
