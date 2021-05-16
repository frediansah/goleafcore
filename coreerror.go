package goleafcore

type CoreError struct {
	ErrorCode        string
	ErrorMessage     string
	ErrorFullMessage string
	ErrorArgs        []interface{}
}

func (e *CoreError) Error() string {
	return e.ErrorMessage
}

func NewCoreError(code, message, fullMessage string, args ...interface{}) *CoreError {
	coreError := CoreError{
		ErrorCode:        code,
		ErrorMessage:     message,
		ErrorFullMessage: fullMessage,
		ErrorArgs:        args,
	}

	return &coreError
}
