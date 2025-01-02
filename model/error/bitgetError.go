package error

import (
	"fmt"
)

type BitgetError struct {
	Code    int
	Message string
}

func (e *BitgetError) Error() string {
	return fmt.Sprintf("Bitget error %d : %s", e.Code, e.Message)
}
