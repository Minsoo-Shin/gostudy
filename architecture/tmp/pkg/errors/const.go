package errors

import "strconv"

type Error int

const (
	ErrBadRequest Error = 4000
)

func (e Error) Error() string {
	return strconv.Itoa(int(e))
}
