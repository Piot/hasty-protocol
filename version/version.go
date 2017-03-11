package version

import "fmt"

// Version : The id of the channel
type Version struct {
	major uint8
	minor uint8
	patch uint8
}

// New : Create Version
func New(major uint8, minor uint8, patch uint8) (out Version, err error) {
	v := Version{major: major, minor: minor, patch: patch}
	return v, nil
}

// String : return string representation
func (in Version) String() string {
	return fmt.Sprintf("[version %d.%d.%d]", in.major, in.minor, in.patch)
}
