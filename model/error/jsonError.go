package error

import (
	"fmt"
)

type JsonError struct {
	Message string
}

func (e *JsonError) Error() string {
	return fmt.Sprintf("Error on json operation : %s", e.Message)
}
