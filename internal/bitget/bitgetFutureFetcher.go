package bitget

// TODO mutualise with spotGetFills

import (
	"fmt"
	"strconv"
	"strings"
	"tradeFetcher/internal/common"
	"tradeFetcher/internal/fetcher"
	"tradeFetcher/internal/json"
	bitgetModel "tradeFetcher/model/bitget"
	customError "tradeFetcher/model/error"
	"tradeFetcher/model/trading"
)

type BitgetFutureFetcher struct {
	tradeGetter           common.IQuery[bitgetModel.FutureTransactionsQueryParameters]
	transactionsConverter json.IJsonConverter[bitgetModel.ApiFutureTransactions]
}

func NewBitgetFutureFetcher(
	tGetter common.IQuery[bitgetModel.FutureTransactionsQueryParameters],
	tConverter json.IJsonConverter[bitgetModel.ApiFutureTransactions],
) fetcher.IFetcher {
	return &BitgetFutureFetcher{
		tradeGetter:           tGetter,
		transactionsConverter: tConverter,
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

		// TODO manage multi fees
		err = f.convertFloat64FromString(trade.FeeDetail[0].FeesValue, "Fees", &trades[index].Fees)
		if err != nil {
			return nil, err
		}

		if trade.Side == bitgetModel.BUY_KEYWORD {
			trades[index].Long = true
		} else if trade.Side == bitgetModel.SELL_KEYWORD {
			trades[index].Long = false
		} else {
			return nil, &customError.BitgetError{
				Code:    9999,
				Message: fmt.Sprintf("Side conversion error on : %s", trade.Side),
			}
		}

		if strings.Contains(trade.TradeSide, bitgetModel.OPEN_KEYWORD) {
			trades[index].Open = true
		} else if strings.Contains(trade.TradeSide, bitgetModel.CLOSE_KEYWORD) {
			trades[index].Open = false
		} else if strings.Contains(trade.TradeSide, bitgetModel.BUY_KEYWORD) {
			trades[index].Open = trades[index].Long
		} else if strings.Contains(trade.TradeSide, bitgetModel.SELL_KEYWORD) {
			trades[index].Open = !trades[index].Long
		} else {
			return nil, &customError.BitgetError{
				Code:    9999,
				Message: fmt.Sprintf("Open/Close conversion error on : %s", trade.TradeSide),
			}
		}
	}

	return trades, nil
}

func (f BitgetFutureFetcher) convertFloat64FromString(input string, fieldName string, output *float64) error {
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

func (f BitgetFutureFetcher) convertInt64FromString(input string, fieldName string, output *int64) error {
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
