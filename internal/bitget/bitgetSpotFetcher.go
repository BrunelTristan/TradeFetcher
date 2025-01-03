package bitget

import (
	"fmt"
	"strconv"
	"tradeFetcher/internal/common"
	"tradeFetcher/internal/fetcher"
	"tradeFetcher/internal/json"
	bitgetModel "tradeFetcher/model/bitget"
	customError "tradeFetcher/model/error"
	"tradeFetcher/model/trading"
)

type BitgetSpotFetcher struct {
	assetsList        []string
	tradeGetter       common.IQuery[bitgetModel.SpotGetFillQueryParameters]
	spotFillConverter json.IJsonConverter[bitgetModel.ApiSpotGetFills]
}

func NewBitgetSpotFetcher(
	assets []string,
	tGetter common.IQuery[bitgetModel.SpotGetFillQueryParameters],
	converter json.IJsonConverter[bitgetModel.ApiSpotGetFills]) fetcher.IFetcher {
	return &BitgetSpotFetcher{
		assetsList:        assets,
		tradeGetter:       tGetter,
		spotFillConverter: converter,
	}
}

func (f BitgetSpotFetcher) FetchLastTrades() ([]trading.Trade, error) {
	trades := make([]trading.Trade, 0)

	for _, asset := range f.assetsList {
		// TODO use go routine for multi threading
		list, err := f.fetchLastTradesForAsset(asset)

		if err != nil {
			return nil, err
		}

		trades = append(trades, list...)
	}

	return trades, nil
}

func (f BitgetSpotFetcher) fetchLastTradesForAsset(asset string) ([]trading.Trade, error) {
	jsonGet, err := f.tradeGetter.Get(&bitgetModel.SpotGetFillQueryParameters{Symbol: asset})

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
		trades[index].Long = true

		err = f.convertFloat64FromString(trade.Price, "Price", &trades[index].Price)
		if err != nil {
			return nil, err
		}

		err = f.convertFloat64FromString(trade.Size, "Quantity", &trades[index].Quantity)
		if err != nil {
			return nil, err
		}

		err = f.convertInt64FromString(trade.LastUpdate, "Timestamp", &trades[index].ExecutedTimestamp)
		if err != nil {
			return nil, err
		}

		trades[index].ExecutedTimestamp /= 1000

		err = f.convertFloat64FromString(trade.FeeDetail.FeesValue, "Fees", &trades[index].Fees)
		if err != nil {
			return nil, err
		}

		if trade.Side == bitgetModel.BUY_KEYWORD {
			trades[index].Open = true
		} else if trade.Side == bitgetModel.SELL_KEYWORD {
			trades[index].Open = false
		} else {
			return nil, &customError.BitgetError{
				Code:    9999,
				Message: fmt.Sprintf("Side conversion to Open/Close error on : %s", trade.Side),
			}
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

func (f BitgetSpotFetcher) convertInt64FromString(input string, fieldName string, output *int64) error {
	val, err := strconv.ParseInt(input, 10, 64)
	if err != nil {
		return &customError.BitgetError{
			Code:    9999,
			Message: fmt.Sprintf("%s conversion to int error on : %s", fieldName, input),
		}
	}

	*output = val

	return nil
}
