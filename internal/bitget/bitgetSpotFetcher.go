package bitget

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

type BitgetSpotFetcher struct {
	asset             string
	tradeGetter       common.IQuery[bitgetModel.SpotGetFillQueryParameters]
	spotFillConverter json.IJsonConverter[bitgetModel.ApiSpotGetFills]
	tradeConverter    converter.IStructConverter[bitgetModel.ApiSpotFill, trading.Trade]
}

func NewBitgetSpotFetcher(
	assetToFetch string,
	tGetter common.IQuery[bitgetModel.SpotGetFillQueryParameters],
	jConverter json.IJsonConverter[bitgetModel.ApiSpotGetFills],
	tConverter converter.IStructConverter[bitgetModel.ApiSpotFill, trading.Trade],
) fetcher.IFetcher {
	return &BitgetSpotFetcher{
		asset:             assetToFetch,
		tradeGetter:       tGetter,
		spotFillConverter: jConverter,
		tradeConverter:    tConverter,
	}
}

func (f BitgetSpotFetcher) FetchLastTrades() ([]trading.Trade, error) {
	jsonGet, err := f.tradeGetter.Get(&bitgetModel.SpotGetFillQueryParameters{Symbol: f.asset})

	if err != nil {
		return nil, err
	}

	apiResponse, err := f.spotFillConverter.Import(jsonGet.(string))

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
