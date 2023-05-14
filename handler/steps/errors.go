package steps

type UnicornError struct {
	error
	msg string
}

func (e *UnicornError) Error() string {
	return e.msg
}

func NewUnicornError(msg string) *UnicornError {
	return &UnicornError{
		msg: msg,
	}
}
