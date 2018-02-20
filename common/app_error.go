package common

type AppError struct {
	Error          error
	ErrorCode      int
	Message        string
	HttpStatusCode int
}
