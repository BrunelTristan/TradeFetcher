package composition

import (
	"tradeFetcher/internal/bitget"
	"tradeFetcher/internal/common"
	"tradeFetcher/internal/configuration"
	"tradeFetcher/internal/converter"
	"tradeFetcher/internal/externalTools"
	"tradeFetcher/internal/fetcher"
	"tradeFetcher/internal/formatter"
	"tradeFetcher/internal/json"
	"tradeFetcher/internal/processUnit"
	"tradeFetcher/internal/security"
	bitgetModel "tradeFetcher/model/bitget"
	configModel "tradeFetcher/model/configuration"
	"tradeFetcher/model/trading"
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

	c.singletons["IApiRouteBuilder"] = externalTools.NewApiRouteBuilder()
	c.singletons["IJsonConverter[bitgetModel.ApiSpotGetFills]"] = json.NewJsonConverter[bitgetModel.ApiSpotGetFills]()
	c.singletons["IJsonConverter[bitgetModel.ApiFutureTransactions]"] = json.NewJsonConverter[bitgetModel.ApiFutureTransactions]()
	c.singletons["IStructConverter[bitgetModel.ApiSpotGetFills,trading.Trade]"] = bitget.NewSpotFillToTradeConverter()
	c.singletons["IStructConverter[bitgetModel.ApiFutureTransaction,trading.Trade]"] = bitget.NewFutureTransactionToTradeConverter()

	c.singletons["IQuery[bitgetModel.ApiQueryParameters]"] = bitget.NewApiQuery(
		c.globalConfig.BitgetAccount,
		bitget.NewApiSignatureBuilder(
			c.globalConfig.BitgetAccount,
			security.NewSha256Crypter(),
			externalTools.NewBase64Encoder()),
	)

	c.singletons["IQuery[bitgetModel.SpotGetFillQueryParameters]"] = bitget.NewSpotFillsGetter(
		c.singletons["IQuery[bitgetModel.ApiQueryParameters]"].(common.IQuery[bitgetModel.ApiQueryParameters]),
		c.singletons["IApiRouteBuilder"].(externalTools.IApiRouteBuilder),
	)
	c.singletons["IQuery[bitgetModel.FutureTransactionsQueryParameters]"] = bitget.NewFutureTransactionsGetter(
		c.singletons["IQuery[bitgetModel.ApiQueryParameters]"].(common.IQuery[bitgetModel.ApiQueryParameters]),
		c.singletons["IApiRouteBuilder"].(externalTools.IApiRouteBuilder),
	)
}

func (c *CompositionRoot) ComposeFetcher() fetcher.IFetcher {
	if c.globalConfig == nil {
		return nil
	}

	bitgetFetcherList := []fetcher.IFetcher{
		bitget.NewTradesFetcher[bitgetModel.FutureTransactionsQueryParameters, bitgetModel.ApiFutureTransaction](
			bitget.NewApiQueryToStructDecorator[bitgetModel.FutureTransactionsQueryParameters, bitgetModel.ApiFutureTransactions](
				c.singletons["IQuery[bitgetModel.FutureTransactionsQueryParameters]"].(common.IQuery[bitgetModel.FutureTransactionsQueryParameters]),
				common.NewNilQueryParametersBuilder[bitgetModel.FutureTransactionsQueryParameters](),
				c.singletons["IJsonConverter[bitgetModel.ApiFutureTransactions]"].(json.IJsonConverter[bitgetModel.ApiFutureTransactions]),
			),
			c.singletons["IStructConverter[bitgetModel.ApiFutureTransaction,trading.Trade]"].(converter.IStructConverter[bitgetModel.ApiFutureTransaction, trading.Trade]),
		),
	}

	for _, asset := range c.globalConfig.BitgetSpotAssets {
		bitgetFetcherList = append(
			bitgetFetcherList,
			bitget.NewTradesFetcher[bitgetModel.SpotGetFillQueryParameters, bitgetModel.ApiSpotFill](
				bitget.NewApiQueryToStructDecorator[bitgetModel.SpotGetFillQueryParameters, bitgetModel.ApiSpotGetFills](
					c.singletons["IQuery[bitgetModel.SpotGetFillQueryParameters]"].(common.IQuery[bitgetModel.SpotGetFillQueryParameters]),
					bitget.NewSpotGetFillQueryParametersBuilder(asset),
					c.singletons["IJsonConverter[bitgetModel.ApiSpotGetFills]"].(json.IJsonConverter[bitgetModel.ApiSpotGetFills]),
				),
				c.singletons["IStructConverter[bitgetModel.ApiSpotGetFills,trading.Trade]"].(converter.IStructConverter[bitgetModel.ApiSpotFill, trading.Trade]),
			),
		)
	}

	return fetcher.NewSortByDateFetcherDecorator(
		fetcher.NewBitgetFetcher(bitgetFetcherList),
	)
}

func (c *CompositionRoot) ComposeProcessUnit() processUnit.IProcessUnit {
	return processUnit.NewTradeDisplayer(
		formatter.NewCsvTradeFormatter())
}
