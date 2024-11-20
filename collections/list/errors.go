package list

import "fmt"

// sentinelError is a type of error which indicates a state of the list.
// Ex - Index out boundary (ErrIndexOutOfBounds), No such elements (ErrNoSuchElement)
type sentinelError string

func (e sentinelError) Error() string {
	return string(e)
}

const (
	ErrNoSuchElement    sentinelError = "ErrNoSuchElement"
	ErrIndexOutOfBounds sentinelError = "ErrIndexOutOfBounds"
)

func errIndexOutOfBounds(index, len int) error {
	return fmt.Errorf("%w: index %d is out of bound. collection size %d", ErrIndexOutOfBounds, index, len)
}

func errNoSuchElement() error {
	return fmt.Errorf("%w: empty collection", ErrNoSuchElement)
}
