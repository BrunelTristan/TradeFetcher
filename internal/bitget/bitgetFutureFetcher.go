package bitget

// TODO mutualise with spotGetFills

import (
	"strconv"
	"tradeFetcher/internal/common"
	"tradeFetcher/internal/converter"
	"tradeFetcher/internal/fetcher"
	"tradeFetcher/internal/json"
	bitgetModel "tradeFetcher/model/bitget"
	customError "tradeFetcher/model/error"
	"tradeFetcher/model/trading"
)

type BitgetFutureFetcher struct {
	tradeGetter           common.IQuery[bitgetModel.FutureTransactionsQueryParameters]
	transactionsConverter json.IJsonConverter[bitgetModel.ApiFutureTransactions]
	tradeConverter        converter.IStructConverter[bitgetModel.ApiFutureTransaction, trading.Trade]
}

func NewBitgetFutureFetcher(
	tGetter common.IQuery[bitgetModel.FutureTransactionsQueryParameters],
	jConverter json.IJsonConverter[bitgetModel.ApiFutureTransactions],
	tConverter converter.IStructConverter[bitgetModel.ApiFutureTransaction, trading.Trade],
) fetcher.IFetcher {
	return &BitgetFutureFetcher{
		tradeGetter:           tGetter,
		transactionsConverter: jConverter,
		tradeConverter:        tConverter,
	}
}

func (f BitgetFutureFetcher) FetchLastTrades() ([]trading.Trade, error) {
	jsonGet, err := f.tradeGetter.Get(&bitgetModel.FutureTransactionsQueryParameters{})

	if err != nil {
		return nil, err
	}

	apiResponse, err := f.transactionsConverter.Import(jsonGet.(string))

	if err != nil {
		return nil, err
	}

	code, err := strconv.Atoi(apiResponse.Code)

	if err != nil || code != 0 {
		return nil, &customError.BitgetError{
			Code:    code,
			Message: apiResponse.Message,
		}
	}

	trades := make([]trading.Trade, len(apiResponse.Data.FillList))

	for index, trade := range apiResponse.Data.FillList {
		convertedTrade, err := f.tradeConverter.Convert(trade)
		if err != nil {
			return nil, &customError.BitgetError{
				Code:    9999,
				Message: err.Error(),
			}
		}
		// TODO use pointer instead of raw data to avoid copy in every layers
		trades[index] = *convertedTrade
	}

	return trades, nil
}
