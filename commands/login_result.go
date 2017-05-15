package commands

import "fmt"

// LoginResult : LoginResult command
type LoginResult struct {
	userID   uint64
	realname string
}

// NewLoginResult : Creates a new Login command
func NewLoginResult(userID uint64, realname string) LoginResult {
	return LoginResult{userID: userID, realname: realname}
}

// String : Returns a human readable string
func (in LoginResult) String() string {
	return fmt.Sprintf("[loginresult realname:%s]", in.realname)
}

// Realname : todo
func (in LoginResult) Realname() string {
	return in.realname
}

// UserID : todo
func (in LoginResult) UserID() uint64 {
	return in.userID
}
