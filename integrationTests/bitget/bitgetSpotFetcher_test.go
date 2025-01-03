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

func TestBitgetSpotFetcherFetchLastTradesWithRealData(t *testing.T) {
	mockController := gomock.NewController(t)

	assetList := []string{"XRPUSDT"}

	externalGetterMock := generatedMocks.NewMockIQuery[bitgetModel.SpotGetFillQueryParameters](mockController)
	jsonConverter := json.NewJsonConverter[bitgetModel.ApiSpotGetFills]()

	externalGetterMock.
		EXPECT().
		Get(gomock.Any()).
		Times(1).
		Return("{\"code\":\"00000\",\"msg\":\"success\",\"requestTime\":1735824191441,\"data\":[{\"userId\":\"123456789\",\"symbol\":\"XRPUSDT\",\"orderId\":\"1111111111111111111\",\"tradeId\":\"2222222222222222221\",\"orderType\":\"market\",\"side\":\"sell\",\"priceAvg\":\"2.3723\",\"size\":\"63.7035\",\"amount\":\"151.12381305\",\"feeDetail\":{\"deduction\":\"yes\",\"feeCoin\":\"BGB\",\"totalDeductionFee\":\"0\",\"totalFee\":\"0.0204152398581561\"},\"tradeScope\":\"taker\",\"cTime\":\"1735776414573\",\"uTime\":\"1735776414605\"},{\"userId\":\"123456789\",\"symbol\":\"XRPUSDT\",\"orderId\":\"1111111111111111112\",\"tradeId\":\"2222222222222222222\",\"orderType\":\"market\",\"side\":\"buy\",\"priceAvg\":\"2.3191\",\"size\":\"63.7035\",\"amount\":\"147.73478685\",\"feeDetail\":{\"deduction\":\"yes\",\"feeCoin\":\"BGB\",\"totalDeductionFee\":\"0\",\"totalFee\":\"0.0198668397175996\"},\"tradeScope\":\"taker\",\"cTime\":\"1735761608295\",\"uTime\":\"1735761608327\"},{\"userId\":\"123456789\",\"symbol\":\"XRPUSDT\",\"orderId\":\"1111111111111111113\",\"tradeId\":\"2222222222222222223\",\"orderType\":\"market\",\"side\":\"sell\",\"priceAvg\":\"2.3081\",\"size\":\"64.9507\",\"amount\":\"149.91271067\",\"feeDetail\":{\"deduction\":\"yes\",\"feeCoin\":\"BGB\",\"totalDeductionFee\":\"0\",\"totalFee\":\"0.0204938770567328\"},\"tradeScope\":\"taker\",\"cTime\":\"1735756412100\",\"uTime\":\"1735756412126\"},{\"userId\":\"123456789\",\"symbol\":\"XRPUSDT\",\"orderId\":\"1111111111111111114\",\"tradeId\":\"2222222222222222224\",\"orderType\":\"market\",\"side\":\"buy\",\"priceAvg\":\"2.2621\",\"size\":\"64.9507\",\"amount\":\"146.92497847\",\"feeDetail\":{\"deduction\":\"yes\",\"feeCoin\":\"BGB\",\"totalDeductionFee\":\"0\",\"totalFee\":\"0.0201993440068741\"},\"tradeScope\":\"taker\",\"cTime\":\"1735754408723\",\"uTime\":\"1735754408738\"},{\"userId\":\"123456789\",\"symbol\":\"XRPUSDT\",\"orderId\":\"1111111111111111115\",\"tradeId\":\"2222222222222222225\",\"orderType\":\"market\",\"side\":\"sell\",\"priceAvg\":\"2.2033\",\"size\":\"238.4214\",\"amount\":\"525.31387062\",\"feeDetail\":{\"deduction\":\"yes\",\"feeCoin\":\"BGB\",\"totalDeductionFee\":\"0\",\"totalFee\":\"0.0720842361056604\"},\"tradeScope\":\"taker\",\"cTime\":\"1735746364634\",\"uTime\":\"1735746364661\"},{\"userId\":\"123456789\",\"symbol\":\"XRPUSDT\",\"orderId\":\"1111111111111111116\",\"tradeId\":\"2222222222222222226\",\"orderType\":\"market\",\"side\":\"buy\",\"priceAvg\":\"2.0191\",\"size\":\"97.0234\",\"amount\":\"195.89994694\",\"feeDetail\":{\"deduction\":\"yes\",\"feeCoin\":\"BGB\",\"totalDeductionFee\":\"0\",\"totalFee\":\"0.0254209176888889\"},\"tradeScope\":\"taker\",\"cTime\":\"1735569211016\",\"uTime\":\"1735569211032\"},{\"userId\":\"123456789\",\"symbol\":\"XRPUSDT\",\"orderId\":\"1111111111111111117\",\"tradeId\":\"2222222222222222227\",\"orderType\":\"market\",\"side\":\"buy\",\"priceAvg\":\"2.1412\",\"size\":\"45.7453\",\"amount\":\"97.94983636\",\"feeDetail\":{\"deduction\":\"yes\",\"feeCoin\":\"BGB\",\"totalDeductionFee\":\"0\",\"totalFee\":\"0.0106096739764681\"},\"tradeScope\":\"taker\",\"cTime\":\"1735234336014\",\"uTime\":\"1735234336030\"}]}", nil)

	bitgetFetcher := bitget.NewBitgetSpotFetcher(assetList, externalGetterMock, jsonConverter)

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
		assert.False(t, trades[0].Open)
		assert.Equal(t, "XRPUSDT", trades[1].Pair)
		assert.Equal(t, 2.3191, trades[1].Price)
		assert.Equal(t, 63.7035, trades[1].Quantity)
		assert.Equal(t, 0.0198668397175996, trades[1].Fees)
		assert.Equal(t, int64(1735761608), trades[1].ExecutedTimestamp)
		assert.True(t, trades[1].Open)
		assert.Equal(t, "XRPUSDT", trades[2].Pair)
		assert.Equal(t, 2.3081, trades[2].Price)
		assert.Equal(t, 64.9507, trades[2].Quantity)
		assert.Equal(t, 0.0204938770567328, trades[2].Fees)
		assert.Equal(t, int64(1735756412), trades[2].ExecutedTimestamp)
		assert.False(t, trades[2].Open)
		assert.Equal(t, "XRPUSDT", trades[3].Pair)
		assert.Equal(t, 2.2621, trades[3].Price)
		assert.Equal(t, 64.9507, trades[3].Quantity)
		assert.Equal(t, 0.0201993440068741, trades[3].Fees)
		assert.Equal(t, int64(1735754408), trades[3].ExecutedTimestamp)
		assert.True(t, trades[3].Open)
		assert.Equal(t, "XRPUSDT", trades[4].Pair)
		assert.Equal(t, 2.2033, trades[4].Price)
		assert.Equal(t, 238.4214, trades[4].Quantity)
		assert.Equal(t, 0.0720842361056604, trades[4].Fees)
		assert.Equal(t, int64(1735746364), trades[4].ExecutedTimestamp)
		assert.False(t, trades[4].Open)
		assert.Equal(t, "XRPUSDT", trades[5].Pair)
		assert.Equal(t, 2.0191, trades[5].Price)
		assert.Equal(t, 97.0234, trades[5].Quantity)
		assert.Equal(t, 0.0254209176888889, trades[5].Fees)
		assert.Equal(t, int64(1735569211), trades[5].ExecutedTimestamp)
		assert.True(t, trades[5].Open)
		assert.Equal(t, "XRPUSDT", trades[6].Pair)
		assert.Equal(t, 2.1412, trades[6].Price)
		assert.Equal(t, 45.7453, trades[6].Quantity)
		assert.Equal(t, 0.0106096739764681, trades[6].Fees)
		assert.Equal(t, int64(1735234336), trades[6].ExecutedTimestamp)
		assert.True(t, trades[6].Open)
	}
}
