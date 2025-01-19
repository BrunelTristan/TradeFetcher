package composition

import (
	"tradeFetcher/internal/bitget"
	"tradeFetcher/internal/common"
	"tradeFetcher/internal/configuration"
	"tradeFetcher/internal/converter"
	"tradeFetcher/internal/externalTools"
	"tradeFetcher/internal/fetcher"
	"tradeFetcher/internal/fileRetriever"
	"tradeFetcher/internal/formatter"
	"tradeFetcher/internal/json"
	"tradeFetcher/internal/processUnit"
	"tradeFetcher/internal/security"
	"tradeFetcher/internal/threading"
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
	c.singletons["IJsonConverter[bitgetModel.ApiFutureTaxTransactions]"] = json.NewJsonConverter[bitgetModel.ApiFutureTaxTransactions]()
	c.singletons["IStructConverter[bitgetModel.ApiSpotGetFills,trading.Trade]"] = bitget.NewSpotFillToTradeConverter()
	c.singletons["IStructConverter[bitgetModel.ApiFutureTransaction,trading.Trade]"] = bitget.NewFutureTransactionToTradeConverter()
	c.singletons["IStructConverter[bitgetModel.ApiFutureTaxTransaction,trading.Trade]"] = bitget.NewFutureTaxTransactionToTradeConverter()

	c.singletons["ILastProceedRetriever"] = fileRetriever.NewLastTradeFileRetriever(c.globalConfig.TradeHistoryFilePath)

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
	c.singletons["IQuery[bitgetModel.FutureTaxTransactionsQueryParameters]"] = bitget.NewFutureTaxTransactionsGetter(
		c.singletons["IQuery[bitgetModel.ApiQueryParameters]"].(common.IQuery[bitgetModel.ApiQueryParameters]),
		c.singletons["IApiRouteBuilder"].(externalTools.IApiRouteBuilder),
	)

	c.singletons["CsvTradeFormatter"] = formatter.NewCsvTradeFormatter()
}

func (c *CompositionRoot) composeFetcher() fetcher.IFetcher {
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

	bitgetFetcherList = append(
		bitgetFetcherList,
		bitget.NewTradesFetcher[bitgetModel.FutureTaxTransactionsQueryParameters, bitgetModel.ApiFutureTaxTransaction](
			bitget.NewApiQueryToStructDecorator[bitgetModel.FutureTaxTransactionsQueryParameters, bitgetModel.ApiFutureTaxTransactions](
				c.singletons["IQuery[bitgetModel.FutureTaxTransactionsQueryParameters]"].(common.IQuery[bitgetModel.FutureTaxTransactionsQueryParameters]),
				common.NewNilQueryParametersBuilder[bitgetModel.FutureTaxTransactionsQueryParameters](),
				c.singletons["IJsonConverter[bitgetModel.ApiFutureTaxTransactions]"].(json.IJsonConverter[bitgetModel.ApiFutureTaxTransactions]),
			),
			c.singletons["IStructConverter[bitgetModel.ApiFutureTaxTransaction,trading.Trade]"].(converter.IStructConverter[bitgetModel.ApiFutureTaxTransaction, trading.Trade]),
		),
	)

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

	return fetcher.NewFilterByDateFetcherDecorator(fetcher.NewSortByDateFetcherDecorator(
		fetcher.NewBitgetFetcher(bitgetFetcherList),
	),
		c.singletons["ILastProceedRetriever"].(processUnit.ILastProceedRetriever),
	)
}

func (c *CompositionRoot) composeProcessUnit() []processUnit.IProcessUnit {
	return []processUnit.IProcessUnit{
		processUnit.NewTradeDisplayer(c.singletons["CsvTradeFormatter"].(formatter.ITradeFormatter)),
		processUnit.NewTradeFileSaver(
			c.singletons["CsvTradeFormatter"].(formatter.ITradeFormatter),
			c.globalConfig.TradeHistoryFilePath,
		),
	}
}

func (c *CompositionRoot) ComposeOrchestration() threading.IThreadOrchestrator {
	if c.globalConfig == nil {
		return nil
	}

	return threading.NewPeriodicThreadOrchestrator(
		threading.NewFetcherProcessorsWorker(
			c.composeFetcher(),
			c.composeProcessUnit(),
		),
		1000*c.globalConfig.OrchestrationPeriodicityInSeconds,
	)
}
