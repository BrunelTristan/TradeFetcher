package error

import (
	"fmt"
)

type RestApiError struct {
	HttpCode int
}

func (e *RestApiError) Error() string {
	return fmt.Sprintf("HTTP error %d on rest API call", e.HttpCode)
}
