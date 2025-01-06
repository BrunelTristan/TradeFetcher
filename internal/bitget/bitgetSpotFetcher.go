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

type BitgetSpotFetcher struct {
	tradeGetter    common.IQuery[bitgetModel.SpotGetFillQueryParameters]
	tradeConverter converter.IStructConverter[bitgetModel.ApiSpotFill, trading.Trade]
}

func NewBitgetSpotFetcher(
	tGetter common.IQuery[bitgetModel.SpotGetFillQueryParameters],
	tConverter converter.IStructConverter[bitgetModel.ApiSpotFill, trading.Trade],
) fetcher.IFetcher {
	return &BitgetSpotFetcher{
		tradeGetter:    tGetter,
		tradeConverter: tConverter,
	}
}

func (f BitgetSpotFetcher) FetchLastTrades() ([]trading.Trade, error) {
	getResponse, err := f.tradeGetter.Get(nil)
	if err != nil {
		return nil, err
	}

	apiResponse := getResponse.(*bitgetModel.ApiSpotGetFills)
	code, err := strconv.Atoi(apiResponse.Code)
	if err != nil || code != 0 {
		return nil, &customError.BitgetError{
			Code:    code,
			Message: apiResponse.Message,
		}
	}

	trades := make([]trading.Trade, len(apiResponse.Data))

	for index, trade := range apiResponse.Data {
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
