package exception

type DataConflictError struct {
	Message string
}

func (e DataConflictError) Error() string {
	return e.Message
}

func NewDataConflictError(message string) DataConflictError {
	return DataConflictError{
		Message: message,
	}
}
