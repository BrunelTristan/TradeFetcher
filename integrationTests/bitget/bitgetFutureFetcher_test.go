package bitget_IT

import (
	"github.com/stretchr/testify/assert"
	"go.uber.org/mock/gomock"
	"testing"
	"tradeFetcher/internal/bitget"
	"tradeFetcher/internal/generatedMocks"
	"tradeFetcher/internal/json"
	bitgetModel "tradeFetcher/model/bitget"
)

func TestBitgetFutureFetcherFetchLastTradesWithRealData(t *testing.T) {
	mockController := gomock.NewController(t)

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.FutureTransactionsQueryParameters](mockController)
	jsonConverter := json.NewJsonConverter[bitgetModel.ApiFutureTransactions]()
	tradeConverter := bitget.NewFutureTransactionToTradeConverter()

	externalGetterMock.
		EXPECT().
		Get(gomock.Any()).
		Times(1).
		Return("{\"code\":\"00000\",\"msg\":\"success\",\"requestTime\":1699267238892,\"data\":{\"fillList\":[{\"tradeId\":\"3687452565\",\"symbol\":\"TRXUSDT\",\"orderId\":\"655446847\",\"price\":\"0.26895\",\"baseVolume\":\"112\",\"feeDetail\":[{\"deduction\":\"no\",\"feeCoin\":\"USDT\",\"totalDeductionFee\":\"0\",\"totalFee\":\"-0.01807344\"}],\"side\":\"buy\",\"quoteVolume\":\"30.1224\",\"profit\":\"0.01232\",\"enterPointSource\":\"api\",\"tradeSide\":\"close\",\"posMode\":\"hedge_mode\",\"tradeScope\":\"taker\",\"cTime\":\"1736019712144\"},{\"tradeId\":\"xxxx\",\"symbol\":\"ETHUSDT\",\"orderId\":\"xxxx\",\"price\":\"1801.33\",\"baseVolume\":\"0.02\",\"feeDetail\":[{\"deduction\":\"no\",\"feeCoin\":\"USDT\",\"totalDeductionFee\":\"0\",\"totalFee\":\"-0.02161596\"}],\"side\":\"sell\",\"quoteVolume\":\"36.0266\",\"profit\":\"0.0252\",\"enterPointSource\":\"ios\",\"tradeSide\":\"sell_single\",\"posMode\":\"one_way_mode\",\"tradeScope\":\"taker\",\"cTime\":\"1698730804882\"}],\"endId\":\"123456789\"}}", nil)

	bitgetFetcher := bitget.NewBitgetFutureFetcher(externalGetterMock, jsonConverter, tradeConverter)

	assert.NotNil(t, bitgetFetcher)

	trades, err := bitgetFetcher.FetchLastTrades()

	assert.Nil(t, err)
	assert.NotEmpty(t, trades)
	assert.Len(t, trades, 2)

	if 2 == len(trades) {
		assert.Equal(t, "TRXUSDT", trades[0].Pair)
		assert.Equal(t, 0.26895, trades[0].Price)
		assert.Equal(t, float64(112), trades[0].Quantity)
		assert.Equal(t, 0.01807344, trades[0].Fees)
		assert.Equal(t, int64(1736019712), trades[0].ExecutedTimestamp)
		assert.False(t, trades[0].Open)
		assert.True(t, trades[0].Long)
		assert.Equal(t, "ETHUSDT", trades[1].Pair)
		assert.Equal(t, 1801.33, trades[1].Price)
		assert.Equal(t, 0.02, trades[1].Quantity)
		assert.Equal(t, 0.02161596, trades[1].Fees)
		assert.Equal(t, int64(1698730804), trades[1].ExecutedTimestamp)
		assert.True(t, trades[1].Open)
		assert.False(t, trades[1].Long)
	}
}
