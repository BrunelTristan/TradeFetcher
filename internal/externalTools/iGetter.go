package externalTools

import (
	"tradeFetcher/internal/error"
)

type IGetter interface {
	Get(parameters interface{}) (interface{}, error.ExternalError)
}
