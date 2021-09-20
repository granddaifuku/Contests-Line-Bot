package errors

type AppError struct {
	message string
}

func (ae *AppError) Error() string {
	return ae.message
}
