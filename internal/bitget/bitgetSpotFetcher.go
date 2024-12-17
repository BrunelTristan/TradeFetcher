package bitget

import (
	"fmt"
	"strconv"
	"tradeFetcher/internal/externalTools"
	"tradeFetcher/internal/fetcher"
	"tradeFetcher/internal/json"
	bitgetModel "tradeFetcher/model/bitget"
	customError "tradeFetcher/model/error"
	"tradeFetcher/model/trading"
)

type BitgetSpotFetcher struct {
	tradeGetter       externalTools.IGetter
	spotFillConverter json.IJsonConverter[bitgetModel.ApiSpotGetFills]
}

func NewBitgetSpotFetcher(tGetter externalTools.IGetter, converter json.IJsonConverter[bitgetModel.ApiSpotGetFills]) fetcher.IFetcher {
	return &BitgetSpotFetcher{
		tradeGetter:       tGetter,
		spotFillConverter: converter,
	}
}

func (f BitgetSpotFetcher) FetchLastTrades() ([]trading.Trade, error) {
	jsonGet, err := f.tradeGetter.Get(nil)

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
		// TODO convert in a dedicated converter struct
		trades[index].Pair = trade.Symbol

		err = f.convertFloat64FromString(trade.Price, "Price", &trades[index].Price)
		if err != nil {
			return nil, err
		}

		err = f.convertFloat64FromString(trade.Size, "Quantity", &trades[index].Quantity)
		if err != nil {
			return nil, err
		}
	}

	return trades, nil
}

func (f BitgetSpotFetcher) convertFloat64FromString(input string, fieldName string, output *float64) error {
	val, err := strconv.ParseFloat(input, 64)
	if err != nil {
		return &customError.BitgetError{
			Code:    9999,
			Message: fmt.Sprintf("%s conversion to float error on : %s", fieldName, input),
		}
	}

	*output = val

	return nil
}
