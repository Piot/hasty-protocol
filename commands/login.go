package commands

import "fmt"

// Login : Login command
type Login struct {
	username string
	password string
}

// NewLogin : Creates a new Login command
func NewLogin(username string, password string) Login {
	return Login{username: username, password: password}
}

// String : Returns a human readable string
func (in Login) String() string {
	return fmt.Sprintf("[login username:%s]", in.username)
}

// Username : todo
func (in Login) Username() string {
	return in.username
}

// Password : todo
func (in Login) Password() string {
	return in.password
}
