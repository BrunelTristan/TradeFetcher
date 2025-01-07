package bitget

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestNewSpotGetFillQueryParametersBuilder(t *testing.T) {
	fakeObject := NewSpotGetFillQueryParametersBuilder("")

	assert.NotNil(t, fakeObject)
}

func TestBuildParameters(t *testing.T) {
	builder := NewSpotGetFillQueryParametersBuilder("TOTO")

	params, err := builder.BuildQueryParameters()

	assert.Nil(t, err)
	assert.NotNil(t, params)

	assert.Equal(t, "TOTO", params.Symbol)
}
