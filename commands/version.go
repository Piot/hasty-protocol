package commands

import "fmt"

// Version : data inside a stream
type Version struct {
	major byte
	minor byte
	patch byte
}

// NewVersion : Creates a version
func NewVersion(major byte, minor byte, patch byte) Version {
	return Version{major: major, minor: minor, patch: patch}
}

// String : Return human readable
func (in Version) String() string {
	return fmt.Sprintf("[version %d.%d.%d]", in.major, in.minor, in.patch)
}

// Major : todo
func (in Version) Major() byte {
	return in.major
}

// Minor : todo
func (in Version) Minor() byte {
	return in.minor
}

// Patch : todo
func (in Version) Patch() byte {
	return in.patch
}
