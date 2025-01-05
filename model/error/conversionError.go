package error

import (
	"fmt"
)

type ConversionError struct {
	InputField   string
	OutputField  string
	InputStruct  string
	OutputStruct string
	Message      string
}

func (e *ConversionError) Error() string {
	return fmt.Sprintf(
		"Conversion error from %s.%s to %s.%s : %s",
		e.InputStruct,
		e.InputField,
		e.OutputStruct,
		e.OutputField,
		e.Message,
	)
}
