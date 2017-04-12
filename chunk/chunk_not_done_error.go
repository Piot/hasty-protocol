package chunk

// NotDoneError : todo
type NotDoneError struct {
	msg string
}

func (e *NotDoneError) Error() string { return e.msg }
