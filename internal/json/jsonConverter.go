package json

import (
	"encoding/json"
	"errors"
	customError "tradeFetcher/internal/error"
)

type JsonConverter[T any] struct {
}

func NewJsonConverter[T any]() IJsonConverter[T] {
	return &JsonConverter[T]{}
}

func (j JsonConverter[T]) Import(jsonText string) (*T, error) {
	var object T

	if err := json.Unmarshal([]byte(jsonText), &object); err != nil {
		jsonError := &customError.JsonError{}

		if errors.As(err, new(*json.UnmarshalTypeError)) {
			jsonError.Message = "Bad field format"
		} else if errors.As(err, new(*json.SyntaxError)) {
			jsonError.Message = "Bad json format"
		} else {
			jsonError.Message = err.Error()
		}

		return nil, jsonError
	}

	return &object, nil
}

func (j JsonConverter[T]) Export(object *T) (string, error) {
	if object == nil {
		return "", nil
	}

	bytes, err := json.Marshal(object)

	return string(bytes), err
}
