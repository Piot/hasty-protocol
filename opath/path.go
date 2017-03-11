package opath

import (
	"errors"
	"fmt"
	"regexp"
)

// OPath : Path to an object or stream
type OPath struct {
	path string
}

func NewFromString(in string) (out OPath, err error) {
	verifyErr := verifyPath(in)
	if verifyErr != nil {
		return OPath{}, verifyErr
	}

	path := OPath{path: in}

	return path, nil
}

func verifyPath(path string) (err error) {
	matched, err2 := regexp.MatchString("^((\\/[@]?[a-z_0-9]+)+)$", path)

	if err2 != nil {
		return err2
	}

	if !matched {
		return errors.New("No match")
	}

	return nil
}

func (in OPath) ToString() string {
	return in.path
}

func (in OPath) String() string {
	return fmt.Sprintf("[opath %s]", in.path)
}
