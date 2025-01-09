package bitget_IT

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"tradeFetcher/internal/bitget"
	"tradeFetcher/internal/generatedMocks"
	"tradeFetcher/internal/json"
	bitgetModel "tradeFetcher/model/bitget"
	"tradeFetcher/model/trading"
)

func TestTradesFetcherFetchLastTradesWithRealSpotData(t *testing.T) {
	mockController := gomock.NewController(t)

	paramBuilderMock := generatedMocks.NewMockIQueryParametersBuilder[bitgetModel.SpotGetFillQueryParameters](mockController)
	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.SpotGetFillQueryParameters](mockController)
	jsonConverter := json.NewJsonConverter[bitgetModel.ApiSpotGetFills]()
	tradeConverter := bitget.NewSpotFillToTradeConverter()
	getterToStruct := bitget.NewApiQueryToStructDecorator[bitgetModel.SpotGetFillQueryParameters, bitgetModel.ApiSpotGetFills](
		externalGetterMock,
		paramBuilderMock,
		jsonConverter,
	)

	builtParameters := &bitgetModel.SpotGetFillQueryParameters{Symbol: "XRPUSDT"}

	paramBuilderMock.
		EXPECT().
		BuildQueryParameters().
		Times(1).
		Return(builtParameters, nil)

	externalGetterMock.
		EXPECT().
		Get(gomock.Eq(builtParameters)).
		Times(1).
		Return("{\"code\":\"00000\",\"msg\":\"success\",\"requestTime\":1735824191441,\"data\":[{\"userId\":\"123456789\",\"symbol\":\"XRPUSDT\",\"orderId\":\"1111111111111111111\",\"tradeId\":\"2222222222222222221\",\"orderType\":\"market\",\"side\":\"sell\",\"priceAvg\":\"2.3723\",\"size\":\"63.7035\",\"amount\":\"151.12381305\",\"feeDetail\":{\"deduction\":\"yes\",\"feeCoin\":\"BGB\",\"totalDeductionFee\":\"0\",\"totalFee\":\"0.0204152398581561\"},\"tradeScope\":\"taker\",\"cTime\":\"1735776414573\",\"uTime\":\"1735776414605\"},{\"userId\":\"123456789\",\"symbol\":\"XRPUSDT\",\"orderId\":\"1111111111111111112\",\"tradeId\":\"2222222222222222222\",\"orderType\":\"market\",\"side\":\"buy\",\"priceAvg\":\"2.3191\",\"size\":\"63.7035\",\"amount\":\"147.73478685\",\"feeDetail\":{\"deduction\":\"yes\",\"feeCoin\":\"BGB\",\"totalDeductionFee\":\"0\",\"totalFee\":\"0.0198668397175996\"},\"tradeScope\":\"taker\",\"cTime\":\"1735761608295\",\"uTime\":\"1735761608327\"},{\"userId\":\"123456789\",\"symbol\":\"XRPUSDT\",\"orderId\":\"1111111111111111113\",\"tradeId\":\"2222222222222222223\",\"orderType\":\"market\",\"side\":\"sell\",\"priceAvg\":\"2.3081\",\"size\":\"64.9507\",\"amount\":\"149.91271067\",\"feeDetail\":{\"deduction\":\"yes\",\"feeCoin\":\"BGB\",\"totalDeductionFee\":\"0\",\"totalFee\":\"0.0204938770567328\"},\"tradeScope\":\"taker\",\"cTime\":\"1735756412100\",\"uTime\":\"1735756412126\"},{\"userId\":\"123456789\",\"symbol\":\"XRPUSDT\",\"orderId\":\"1111111111111111114\",\"tradeId\":\"2222222222222222224\",\"orderType\":\"market\",\"side\":\"buy\",\"priceAvg\":\"2.2621\",\"size\":\"64.9507\",\"amount\":\"146.92497847\",\"feeDetail\":{\"deduction\":\"yes\",\"feeCoin\":\"BGB\",\"totalDeductionFee\":\"0\",\"totalFee\":\"0.0201993440068741\"},\"tradeScope\":\"taker\",\"cTime\":\"1735754408723\",\"uTime\":\"1735754408738\"},{\"userId\":\"123456789\",\"symbol\":\"XRPUSDT\",\"orderId\":\"1111111111111111115\",\"tradeId\":\"2222222222222222225\",\"orderType\":\"market\",\"side\":\"sell\",\"priceAvg\":\"2.2033\",\"size\":\"238.4214\",\"amount\":\"525.31387062\",\"feeDetail\":{\"deduction\":\"yes\",\"feeCoin\":\"BGB\",\"totalDeductionFee\":\"0\",\"totalFee\":\"0.0720842361056604\"},\"tradeScope\":\"taker\",\"cTime\":\"1735746364634\",\"uTime\":\"1735746364661\"},{\"userId\":\"123456789\",\"symbol\":\"XRPUSDT\",\"orderId\":\"1111111111111111116\",\"tradeId\":\"2222222222222222226\",\"orderType\":\"market\",\"side\":\"buy\",\"priceAvg\":\"2.0191\",\"size\":\"97.0234\",\"amount\":\"195.89994694\",\"feeDetail\":{\"deduction\":\"yes\",\"feeCoin\":\"BGB\",\"totalDeductionFee\":\"0\",\"totalFee\":\"0.0254209176888889\"},\"tradeScope\":\"taker\",\"cTime\":\"1735569211016\",\"uTime\":\"1735569211032\"},{\"userId\":\"123456789\",\"symbol\":\"XRPUSDT\",\"orderId\":\"1111111111111111117\",\"tradeId\":\"2222222222222222227\",\"orderType\":\"market\",\"side\":\"buy\",\"priceAvg\":\"2.1412\",\"size\":\"45.7453\",\"amount\":\"97.94983636\",\"feeDetail\":{\"deduction\":\"yes\",\"feeCoin\":\"BGB\",\"totalDeductionFee\":\"0\",\"totalFee\":\"0.0106096739764681\"},\"tradeScope\":\"taker\",\"cTime\":\"1735234336014\",\"uTime\":\"1735234336030\"}]}", nil)

	bitgetFetcher := bitget.NewTradesFetcher[bitgetModel.SpotGetFillQueryParameters, bitgetModel.ApiSpotFill](getterToStruct, tradeConverter)

	assert.NotNil(t, bitgetFetcher)

	trades, err := bitgetFetcher.FetchLastTrades()

	assert.Nil(t, err)
	assert.NotEmpty(t, trades)
	assert.Len(t, trades, 7)

	if 7 == len(trades) {
		assert.Equal(t, "XRPUSDT", trades[0].Pair)
		assert.Equal(t, 2.3723, trades[0].Price)
		assert.Equal(t, 63.7035, trades[0].Quantity)
		assert.Equal(t, 0.0204152398581561, trades[0].Fees)
		assert.Equal(t, int64(1735776414), trades[0].ExecutedTimestamp)
		assert.Equal(t, trading.CLOSE, trades[0].TransactionType)
		assert.True(t, trades[0].Long)
		assert.Equal(t, "XRPUSDT", trades[1].Pair)
		assert.Equal(t, 2.3191, trades[1].Price)
		assert.Equal(t, 63.7035, trades[1].Quantity)
		assert.Equal(t, 0.0198668397175996, trades[1].Fees)
		assert.Equal(t, int64(1735761608), trades[1].ExecutedTimestamp)
		assert.Equal(t, trading.OPENING, trades[1].TransactionType)
		assert.True(t, trades[1].Long)
		assert.Equal(t, "XRPUSDT", trades[2].Pair)
		assert.Equal(t, 2.3081, trades[2].Price)
		assert.Equal(t, 64.9507, trades[2].Quantity)
		assert.Equal(t, 0.0204938770567328, trades[2].Fees)
		assert.Equal(t, int64(1735756412), trades[2].ExecutedTimestamp)
		assert.Equal(t, trading.CLOSE, trades[2].TransactionType)
		assert.True(t, trades[2].Long)
		assert.Equal(t, "XRPUSDT", trades[3].Pair)
		assert.Equal(t, 2.2621, trades[3].Price)
		assert.Equal(t, 64.9507, trades[3].Quantity)
		assert.Equal(t, 0.0201993440068741, trades[3].Fees)
		assert.Equal(t, int64(1735754408), trades[3].ExecutedTimestamp)
		assert.Equal(t, trading.OPENING, trades[3].TransactionType)
		assert.True(t, trades[3].Long)
		assert.Equal(t, "XRPUSDT", trades[4].Pair)
		assert.Equal(t, 2.2033, trades[4].Price)
		assert.Equal(t, 238.4214, trades[4].Quantity)
		assert.Equal(t, 0.0720842361056604, trades[4].Fees)
		assert.Equal(t, int64(1735746364), trades[4].ExecutedTimestamp)
		assert.Equal(t, trading.CLOSE, trades[4].TransactionType)
		assert.True(t, trades[4].Long)
		assert.Equal(t, "XRPUSDT", trades[5].Pair)
		assert.Equal(t, 2.0191, trades[5].Price)
		assert.Equal(t, 97.0234, trades[5].Quantity)
		assert.Equal(t, 0.0254209176888889, trades[5].Fees)
		assert.Equal(t, int64(1735569211), trades[5].ExecutedTimestamp)
		assert.Equal(t, trading.OPENING, trades[5].TransactionType)
		assert.True(t, trades[5].Long)
		assert.Equal(t, "XRPUSDT", trades[6].Pair)
		assert.Equal(t, 2.1412, trades[6].Price)
		assert.Equal(t, 45.7453, trades[6].Quantity)
		assert.Equal(t, 0.0106096739764681, trades[6].Fees)
		assert.Equal(t, int64(1735234336), trades[6].ExecutedTimestamp)
		assert.Equal(t, trading.OPENING, trades[6].TransactionType)
		assert.True(t, trades[6].Long)
	}
}

func TestTradesFetcherFetchLastTradesWithRealFutureData(t *testing.T) {
	mockController := gomock.NewController(t)

	paramBuilderMock := generatedMocks.NewMockIQueryParametersBuilder[bitgetModel.FutureTransactionsQueryParameters](mockController)
	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.FutureTransactionsQueryParameters](mockController)
	jsonConverter := json.NewJsonConverter[bitgetModel.ApiFutureTransactions]()
	tradeConverter := bitget.NewFutureTransactionToTradeConverter()

	getterToStruct := bitget.NewApiQueryToStructDecorator[bitgetModel.FutureTransactionsQueryParameters, bitgetModel.ApiFutureTransactions](
		externalGetterMock,
		paramBuilderMock,
		jsonConverter,
	)

	paramBuilderMock.
		EXPECT().
		BuildQueryParameters().
		Times(1).
		Return(nil, nil)

	externalGetterMock.
		EXPECT().
		Get(gomock.Any()).
		Times(1).
		Return("{\"code\":\"00000\",\"msg\":\"success\",\"requestTime\":1699267238892,\"data\":{\"fillList\":[{\"tradeId\":\"3687452565\",\"symbol\":\"TRXUSDT\",\"orderId\":\"655446847\",\"price\":\"0.26895\",\"baseVolume\":\"112\",\"feeDetail\":[{\"deduction\":\"no\",\"feeCoin\":\"USDT\",\"totalDeductionFee\":\"0\",\"totalFee\":\"-0.01807344\"}],\"side\":\"buy\",\"quoteVolume\":\"30.1224\",\"profit\":\"0.01232\",\"enterPointSource\":\"api\",\"tradeSide\":\"close\",\"posMode\":\"hedge_mode\",\"tradeScope\":\"taker\",\"cTime\":\"1736019712144\"},{\"tradeId\":\"xxxx\",\"symbol\":\"ETHUSDT\",\"orderId\":\"xxxx\",\"price\":\"1801.33\",\"baseVolume\":\"0.02\",\"feeDetail\":[{\"deduction\":\"no\",\"feeCoin\":\"USDT\",\"totalDeductionFee\":\"0\",\"totalFee\":\"-0.02161596\"}],\"side\":\"sell\",\"quoteVolume\":\"36.0266\",\"profit\":\"0.0252\",\"enterPointSource\":\"ios\",\"tradeSide\":\"sell_single\",\"posMode\":\"one_way_mode\",\"tradeScope\":\"taker\",\"cTime\":\"1698730804882\"}],\"endId\":\"123456789\"}}", nil)

	bitgetFetcher := bitget.NewTradesFetcher[bitgetModel.FutureTransactionsQueryParameters, bitgetModel.ApiFutureTransaction](getterToStruct, tradeConverter)

	assert.NotNil(t, bitgetFetcher)

	trades, err := bitgetFetcher.FetchLastTrades()

	assert.Nil(t, err)
	assert.NotEmpty(t, trades)
	assert.Len(t, trades, 2)

	if 2 == len(trades) {
		assert.Equal(t, "TRXUSDT", trades[0].Pair)
		assert.Equal(t, 0.26895, trades[0].Price)
		assert.Equal(t, 112.0, trades[0].Quantity)
		assert.Equal(t, 0.01807344, trades[0].Fees)
		assert.Equal(t, int64(1736019712), trades[0].ExecutedTimestamp)
		assert.Equal(t, trading.CLOSE, trades[0].TransactionType)
		assert.True(t, trades[0].Long)
		assert.Equal(t, "ETHUSDT", trades[1].Pair)
		assert.Equal(t, 1801.33, trades[1].Price)
		assert.Equal(t, 0.02, trades[1].Quantity)
		assert.Equal(t, 0.02161596, trades[1].Fees)
		assert.Equal(t, int64(1698730804), trades[1].ExecutedTimestamp)
		assert.Equal(t, trading.OPENING, trades[1].TransactionType)
		assert.False(t, trades[1].Long)
	}
}

func TestTradesFetcherFetchLastTradesWithRealFutureTaxData(t *testing.T) {
	mockController := gomock.NewController(t)

	paramBuilderMock := generatedMocks.NewMockIQueryParametersBuilder[bitgetModel.FutureTaxTransactionsQueryParameters](mockController)
	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.FutureTaxTransactionsQueryParameters](mockController)
	jsonConverter := json.NewJsonConverter[bitgetModel.ApiFutureTaxTransactions]()
	tradeConverter := bitget.NewFutureTaxTransactionToTradeConverter()

	getterToStruct := bitget.NewApiQueryToStructDecorator[bitgetModel.FutureTaxTransactionsQueryParameters, bitgetModel.ApiFutureTaxTransactions](
		externalGetterMock,
		paramBuilderMock,
		jsonConverter,
	)

	paramBuilderMock.
		EXPECT().
		BuildQueryParameters().
		Times(1).
		Return(nil, nil)

	externalGetterMock.
		EXPECT().
		Get(gomock.Any()).
		Times(1).
		Return("{\"code\":\"00000\",\"msg\":\"success\",\"requestTime\":1736428213984,\"data\":[{\"id\":\"1260617742759391277\",\"symbol\":\"AAVEUSDT\",\"marginCoin\":\"USDT\",\"futureTaxType\":\"open_long\",\"amount\":\"0\",\"fee\":\"-0.0185376\",\"ts\":\"1736280611695\"},{\"id\":\"1260675616772616210\",\"symbol\":\"AAVEUSDT\",\"marginCoin\":\"USDT\",\"futureTaxType\":\"contract_main_settle_fee\",\"amount\":\"-0.0030672\",\"fee\":\"0\",\"ts\":\"1736294409935\"},{\"id\":\"1260752787134378007\",\"symbol\":\"AAVEUSDT\",\"marginCoin\":\"USDT\",\"futureTaxType\":\"close_long\",\"amount\":\"-0.896\",\"fee\":\"-0.018\",\"ts\":\"1736312808783\"},{\"id\":\"1260906543729766420\",\"symbol\":\"TRXUSDT\",\"marginCoin\":\"USDT\",\"futureTaxType\":\"open_short\",\"amount\":\"0\",\"fee\":\"-0.059945412\",\"ts\":\"1736349467212\"},{\"id\":\"1260917227205062710\",\"symbol\":\"TRXUSDT\",\"marginCoin\":\"USDT\",\"futureTaxType\":\"contract_main_settle_fee\",\"amount\":\"0.0052043524\",\"fee\":\"0\",\"ts\":\"1736352014351\"},{\"id\":\"1260935952998232066\",\"symbol\":\"TRXUSDT\",\"marginCoin\":\"USDT\",\"futureTaxType\":\"close_short\",\"amount\":\"0.31126\",\"fee\":\"-0.011741928\",\"ts\":\"1736356478928\"}]}", nil)

	bitgetFetcher := bitget.NewTradesFetcher[bitgetModel.FutureTaxTransactionsQueryParameters, bitgetModel.ApiFutureTaxTransaction](getterToStruct, tradeConverter)

	assert.NotNil(t, bitgetFetcher)

	trades, err := bitgetFetcher.FetchLastTrades()

	assert.Nil(t, err)
	assert.NotEmpty(t, trades)
	assert.Len(t, trades, 2)

	if 2 == len(trades) {
		assert.Equal(t, "AAVEUSDT", trades[0].Pair)
		assert.Equal(t, 0.0, trades[0].Price)
		assert.Equal(t, 0.0, trades[0].Quantity)
		assert.Equal(t, 0.0030672, trades[0].Fees)
		assert.Equal(t, int64(1736294409), trades[0].ExecutedTimestamp)
		assert.Equal(t, trading.FUNDING, trades[0].TransactionType)
		assert.Equal(t, "TRXUSDT", trades[1].Pair)
		assert.Equal(t, 0.0, trades[1].Price)
		assert.Equal(t, 0.0, trades[1].Quantity)
		assert.Equal(t, -0.0052043524, trades[1].Fees)
		assert.Equal(t, int64(1736352014), trades[1].ExecutedTimestamp)
		assert.Equal(t, trading.FUNDING, trades[1].TransactionType)
	}
}
