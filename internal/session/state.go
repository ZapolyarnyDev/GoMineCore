package session

type State int

const (
	Handshake State = iota
	Status
	Login
	Play
)
