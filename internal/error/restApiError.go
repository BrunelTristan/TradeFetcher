package error

type RestApiError struct {
	ExternalError,
	HttpCode int
}
