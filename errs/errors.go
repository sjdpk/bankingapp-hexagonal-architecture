package errs

type AppError struct {
	Code    int
	Message string
}

func HandleError(code int, msg string) *AppError {
	return &AppError{
		Code:    code,
		Message: msg,
	}
}
