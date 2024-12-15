package error

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestRestApiError(t *testing.T) {
	restError := RestApiError{HttpCode: 123}

	output := restError.Error()

	assert.Equal(t, "HTTP error 123 on rest API call", output)
}
