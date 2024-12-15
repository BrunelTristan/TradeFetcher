package error

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestJsonError(t *testing.T) {
	restError := JsonError{Message: "something weird"}

	output := restError.Error()

	assert.Equal(t, "Error on json operation : something weird", output)
}
