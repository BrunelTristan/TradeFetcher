package composition

import (
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
	singletons           map[string]interface{}
	startupConfiguration *configModel.CmdLineConfiguration
	configLoader         configuration.IConfigurationLoader[configModel.GlobalConfiguration]
	globalConfig         *configModel.GlobalConfiguration
}

func NewCompositionRoot(conf *configModel.CmdLineConfiguration) *CompositionRoot {
	return &CompositionRoot{
		singletons:           make(map[string]interface{}),
		startupConfiguration: conf,
	}
}

func (c *CompositionRoot) Build() {
	c.configLoader = configuration.NewConfigurationLoaderFromJsonFile[configModel.GlobalConfiguration](c.startupConfiguration.ConfigFilePath)
	c.globalConfig, _ = c.configLoader.Load()

	if c.globalConfig == nil {
		return
	}

	c.singletons["IQuery[bitgetModel.ApiQueryParameters]"] = bitget.NewBitgetApiQuery(
		c.globalConfig.BitgetAccount,
		bitget.NewBitgetApiSignatureBuilder(
			c.globalConfig.BitgetAccount,
			security.NewSha256Crypter(),
			externalTools.NewBase64Encoder()),
	)

	c.singletons["IApiRouteBuilder"] = externalTools.NewApiRouteBuilder()
}

func (c *CompositionRoot) ComposeFetcher() fetcher.IFetcher {
	if c.globalConfig == nil {
		return nil
	}

	return fetcher.NewSortByDateFetcherDecorator(
		fetcher.NewBitgetFetcher(
			bitget.NewBitgetSpotFetcher(
				c.globalConfig.BitgetSpotAssets,
				bitget.NewBitgetSpotFillsGetter(
					c.singletons["IQuery[bitgetModel.ApiQueryParameters]"].(*bitget.BitgetApiQuery),
					c.singletons["IApiRouteBuilder"].(*externalTools.ApiRouteBuilder),
				),
				json.NewJsonConverter[bitgetModel.ApiSpotGetFills](),
			),
			bitget.NewBitgetFutureFetcher(
				bitget.NewBitgetFutureTransactionsGetter(
					c.singletons["IQuery[bitgetModel.ApiQueryParameters]"].(*bitget.BitgetApiQuery),
					c.singletons["IApiRouteBuilder"].(*externalTools.ApiRouteBuilder),
				),
				json.NewJsonConverter[bitgetModel.ApiFutureTransactions](),
			),
		),
	)
}

func (c *CompositionRoot) ComposeProcessUnit() processUnit.IProcessUnit {
	return processUnit.NewTradeDisplayer(
		formatter.NewCsvTradeFormatter())
}
