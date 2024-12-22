package bitget_IT

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tradeFetcher/internal/bitget"
	"tradeFetcher/internal/configuration"
	"tradeFetcher/internal/externalTools"
	"tradeFetcher/internal/security"
	bitgetModel "tradeFetcher/model/bitget"
	configModel "tradeFetcher/model/configuration"
)

func TestCallApiCommandSimpleGet(t *testing.T) {
	configLoader := configuration.NewConfigurationLoaderFromJsonFile[configModel.GlobalConfiguration]("../files/globalConfig.json")
	globalConfig, err := configLoader.Load()

	assert.Nil(t, err)
	assert.NotNil(t, globalConfig)
	assert.NotEmpty(t, globalConfig)
	assert.NotNil(t, globalConfig.BitgetAccount)

	crypter := security.NewSha256Crypter()
	encoder := externalTools.NewBase64Encoder()

	signBuilder := bitget.NewBitgetApiSignatureBuilder(
		globalConfig.BitgetAccount,
		crypter,
		encoder)

	api := bitget.NewBitgetApiCommand(
		globalConfig.BitgetAccount,
		signBuilder)

	parameters := &bitgetModel.ApiCommandParameters{
		Route: "/api/v2/public/time",
	}

	output, err := api.Get(parameters)

	assert.Nil(t, err)
	assert.NotNil(t, output)

	assert.NotEmpty(t, output)
	assert.Equal(t, "{\"code\":\"00000\",\"msg\":\"success\"", output.(string)[0:31])
}

func TestCallApiCommandGetWithSignature(t *testing.T) {
	configLoader := configuration.NewConfigurationLoaderFromJsonFile[configModel.GlobalConfiguration]("/src/integrationTests/files/globalConfig.json")
	globalConfig, _ := configLoader.Load()

	signBuilder := bitget.NewBitgetApiSignatureBuilder(
		globalConfig.BitgetAccount,
		security.NewSha256Crypter(),
		externalTools.NewBase64Encoder())

	api := bitget.NewBitgetApiCommand(
		globalConfig.BitgetAccount,
		signBuilder)

	parameters := &bitgetModel.ApiCommandParameters{
		Route: "/api/v2/spot/account/assets",
	}

	output, err := api.Get(parameters)

	assert.Nil(t, err)
	assert.NotNil(t, output)

	assert.NotEmpty(t, output)
	assert.Equal(t, "{\"code\":\"00000\",\"msg\":\"success\"", output.(string)[0:31])
}
