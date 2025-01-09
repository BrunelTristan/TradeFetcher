package bitget

import (
	"strconv"
	"tradeFetcher/internal/common"
	"tradeFetcher/internal/converter"
	"tradeFetcher/internal/fetcher"
	bitgetModel "tradeFetcher/model/bitget"
	customError "tradeFetcher/model/error"
	"tradeFetcher/model/trading"
)

type TradesFetcher[P any, O any] struct {
	tradeGetter    common.IQuery[P]
	tradeConverter converter.IStructConverter[O, trading.Trade]
}

func NewTradesFetcher[P any, O any](
	tGetter common.IQuery[P],
	tConverter converter.IStructConverter[O, trading.Trade],
) fetcher.IFetcher {
	return &TradesFetcher[P, O]{
		tradeGetter:    tGetter,
		tradeConverter: tConverter,
	}
}

func (f TradesFetcher[P, O]) FetchLastTrades() ([]*trading.Trade, error) {
	getResponse, err := f.tradeGetter.Get(nil)
	if err != nil {
		return nil, err
	}

	apiResponse := getResponse.(bitgetModel.IApiResponse[O])
	code, err := strconv.Atoi(apiResponse.GetCode())
	if err != nil || code != 0 {
		return nil, &customError.BitgetError{
			Code:    code,
			Message: apiResponse.GetMessage(),
		}
	}

	tradeList := apiResponse.GetList()
	trades := []*trading.Trade{}

	for _, trade := range tradeList {
		convertedTrade, err := f.tradeConverter.Convert(trade)
		if err != nil {
			return nil, &customError.BitgetError{
				Code:    9999,
				Message: err.Error(),
			}
		} else if convertedTrade != nil {
			trades = append(trades, convertedTrade)
		}
	}

	return trades, nil
}
