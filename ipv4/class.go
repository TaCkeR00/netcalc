package ipv4

type Class uint8

const (
	A Class = iota + 1
	B
	C
	D
	E
	INVALID
)

func (c Class) String() string {
	switch c {
	case A:
		return "A"
	case B:
		return "B"
	case C:
		return "C"
	case D:
		return "D"
	case E:
		return "E"
	default:
		return "invalid class"
	}
}
