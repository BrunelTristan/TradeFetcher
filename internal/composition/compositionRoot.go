package composition

import (
	"reflect"
	"tradeFetcher/internal/bitget"
	"tradeFetcher/internal/configuration"
	"tradeFetcher/internal/externalTools"
	"tradeFetcher/internal/fetcher"
	"tradeFetcher/internal/formatter"
	"tradeFetcher/internal/json"
	"tradeFetcher/internal/processUnit"
	"tradeFetcher/internal/security"
	bitgetModel "tradeFetcher/model/bitget"
	configModel "tradeFetcher/model/configuration"
)

type CompositionRoot struct {
	singletons           map[reflect.Type]interface{}
	startupConfiguration *configModel.CmdLineConfiguration
	configLoader         configuration.IConfigurationLoader[configModel.GlobalConfiguration]
	globalConfig         *configModel.GlobalConfiguration
}

func NewCompositionRoot(conf *configModel.CmdLineConfiguration) *CompositionRoot {
	return &CompositionRoot{
		singletons:           make(map[reflect.Type]interface{}),
		startupConfiguration: conf,
	}
}

func (c *CompositionRoot) Build() {
	// TODO load from config
	c.configLoader = configuration.NewConfigurationLoaderFromJsonFile[configModel.GlobalConfiguration](c.startupConfiguration.ConfigFilePath)
	c.globalConfig, _ = c.configLoader.Load()
}

func (c *CompositionRoot) ComposeFetcher() fetcher.IFetcher {
	if c.globalConfig == nil {
		return nil
	}

	return fetcher.NewBitgetFetcher(
		bitget.NewBitgetSpotFetcher(
			[]string{"XRPUSDT"},
			bitget.NewBitgetSpotFillsGetter(
				bitget.NewBitgetApiQuery(
					c.globalConfig.BitgetAccount,
					bitget.NewBitgetApiSignatureBuilder(
						c.globalConfig.BitgetAccount,
						security.NewSha256Crypter(),
						externalTools.NewBase64Encoder()),
				),
				externalTools.NewApiRouteBuilder(),
			),
			json.NewJsonConverter[bitgetModel.ApiSpotGetFills](),
		),
	)
}

func (c *CompositionRoot) ComposeProcessUnit() processUnit.IProcessUnit {
	return processUnit.NewTradeDisplayer(
		formatter.NewCsvTradeFormatter())
}
