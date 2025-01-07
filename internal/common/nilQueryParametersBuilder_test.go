package common

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewNilQueryParametersBuilder(t *testing.T) {
	fakeObject := NewNilQueryParametersBuilder[int]()

	assert.NotNil(t, fakeObject)
}

func TestNilQueryParametersBuilderBuildQueryParameters(t *testing.T) {
	builder := NewNilQueryParametersBuilder[string]()

	assert.NotNil(t, builder)

	params, err := builder.BuildQueryParameters()

	assert.Nil(t, params)
	assert.Nil(t, err)
}
