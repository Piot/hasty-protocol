package packet

type PacketNotDoneError struct {
	msg string
}

func (e *PacketNotDoneError) Error() string { return e.msg }
