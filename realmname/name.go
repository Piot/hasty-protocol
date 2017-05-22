package realmname

import (
	"fmt"
	"regexp"
)

// Name : todo
type Name struct {
	value string
}

func verify(value string) error {
	//
	// ^([a-zA-Z0-9]([a-zA-Z0-9\\-]{0,61}[a-zA-Z0-9])?\\.)+[a-zA-Z\\-]{0,61}$
	matched, err2 := regexp.MatchString("^([a-zA-Z0-9]([a-zA-Z0-9\\-\\.]{0,61}[a-zA-Z0-9])?[\\-\\.])+[a-zA-Z0-9]{0,61}$", value)

	if err2 != nil {
		return err2
	}

	if !matched {
		return fmt.Errorf("Name is not ok '%s'", value)
	}

	return nil
}

// NewName : todo
func NewName(value string) (Name, error) {
	verifyErr := verify(value)
	if verifyErr != nil {
		return Name{}, verifyErr
	}
	return Name{value: value}, nil
}

// String : Returns human readable string
func (in Name) String() string {
	return fmt.Sprintf("[realm %s]", in.value)
}

// Name : todo
func (in Name) Name() string {
	return in.value
}
