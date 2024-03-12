package error

type UserErr struct {
	err     error
	userMsg string
}

// NewUserErr returns a new error that wraps an existing error with an error message
// dedicated to the user
func NewUserErr(userMsg string, err error) UserErr {
	return UserErr{
		err:     err,
		userMsg: userMsg,
	}
}

// UserMsg returns the message dedicated to the user
func (e UserErr) UserMsg() string {
	return e.userMsg
}

// Error returns the error message of the underlying Error
func (e UserErr) Error() string {
	if e.err == nil {
		return e.userMsg
	}
	return e.err.Error()
}

// Unwrap returns the underlying error
func (e UserErr) Unwrap() error {
	return e.err
}
