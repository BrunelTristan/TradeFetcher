package bitget_IT

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tradeFetcher/internal/bitget"
	bitgetModel "tradeFetcher/model/bitget"
)

func TestCallApiCommandSimpleGet(t *testing.T) {
	api := bitget.NewBitgetApiCommand(nil, nil)
	parameters := &bitgetModel.ApiCommandParameters{
		Route: "/api/v2/public/time",
	}

	output, err := api.Get(parameters)

	assert.Nil(t, err)
	assert.NotNil(t, output)

	assert.NotEmpty(t, output)
	assert.Equal(t, "{\"code\":\"00000\",\"msg\":\"success\"", output.(string)[0:31])
}

/* TODO enable when signature is OK
func TestCallApiCommandGetWithSignature(t *testing.T) {
	api := bitget.NewBitgetApiCommand(nil, nil)
	parameters := &bitgetModel.ApiCommandParameters{
		Route: "/api/v2/spot/account/assets",
	}

	output, err := api.Get(parameters)

	assert.Nil(t, err)
	assert.NotNil(t, output)

	assert.NotEmpty(t, output)
	assert.Equal(t, "{\"code\":\"00000\",\"msg\":\"success\"", output.(string)[0:31])
} */
