package errors

type Error interface {
	Error()	string
}

type errorMsg struct {
	s	string
}

func (e *errorMsg) Error() string {
	return e.s
}

func New(msg string) Error {
	return &errorMsg{msg}
}

