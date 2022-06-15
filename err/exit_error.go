package err

type ExitError struct {
	msg string
}

func NewExitError(msg string) *ExitError {
	return &ExitError{
		msg: msg,
	}
}

func IsExitError(err error) bool {
	if err == nil {
		return false
	}
	_, ok := err.(*ExitError)
	return ok
}

func (e *ExitError) Error() string {
	return e.msg
}
