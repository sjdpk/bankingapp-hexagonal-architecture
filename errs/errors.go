package errs

type AppError struct {
	Code    int    `json:"-"`
	Message string `json:"message"`
}

func HandleError(code int, msg string) *AppError {
	return &AppError{
		Code:    code,
		Message: msg,
	}
}
