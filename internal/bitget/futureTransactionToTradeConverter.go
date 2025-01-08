package bitget

import (
	"fmt"
	"strconv"
	"strings"
	"tradeFetcher/internal/converter"
	bitgetModel "tradeFetcher/model/bitget"
	customError "tradeFetcher/model/error"
	"tradeFetcher/model/trading"
)

type FutureTransactionToTradeConverter struct {
}

func NewFutureTransactionToTradeConverter() converter.IStructConverter[bitgetModel.ApiFutureTransaction, trading.Trade] {
	return &FutureTransactionToTradeConverter{}
}

func (c *FutureTransactionToTradeConverter) Convert(parameters *bitgetModel.ApiFutureTransaction) (*trading.Trade, error) {
	if parameters == nil {
		return nil, nil
	}

	trade := &trading.Trade{Pair: parameters.Symbol}

	floatVal, err := strconv.ParseFloat(parameters.Price, 64)
	if err != nil {
		return nil, c.buildConvertionError(
			"Price",
			"Price",
			parameters.Price,
			" is not a float 64")
	}
	trade.Price = floatVal

	floatVal, err = strconv.ParseFloat(parameters.Size, 64)
	if err != nil {
		return nil, c.buildConvertionError(
			"Size",
			"Quantity",
			parameters.Size,
			" is not a float 64")
	}
	trade.Quantity = floatVal

	floatVal, err = strconv.ParseFloat(parameters.FeeDetail[0].FeesValue, 64)
	if err != nil {
		return nil, c.buildConvertionError(
			"FeeDetail.FeesValue",
			"Fees",
			parameters.FeeDetail[0].FeesValue,
			" is not a float 64")
	}
	trade.Fees = -floatVal

	intVal, err := strconv.ParseInt(parameters.LastUpdate, 10, 64)
	if err != nil {
		return nil, c.buildConvertionError(
			"LastUpdate",
			"ExecutedTimestamp",
			parameters.LastUpdate,
			" is not a int 64")
	}
	trade.ExecutedTimestamp = intVal / 1000

	if parameters.Side == bitgetModel.BUY_KEYWORD {
		trade.Long = true
	} else if parameters.Side == bitgetModel.SELL_KEYWORD {
		trade.Long = false
	} else {
		return nil, c.buildConvertionError(
			"Side",
			"Long",
			parameters.Side,
			fmt.Sprintf(" is not %s or %s", bitgetModel.BUY_KEYWORD, bitgetModel.SELL_KEYWORD))
	}

	if strings.Contains(parameters.TradeSide, bitgetModel.OPEN_KEYWORD) {
		trade.TransactionType = trading.OPENING
	} else if strings.Contains(parameters.TradeSide, bitgetModel.CLOSE_KEYWORD) {
		trade.TransactionType = trading.CLOSE
	} else if strings.Contains(parameters.TradeSide, bitgetModel.BUY_KEYWORD) {
		if trade.Long {
			trade.TransactionType = trading.OPENING
		} else {
			trade.TransactionType = trading.CLOSE
		}
	} else if strings.Contains(parameters.TradeSide, bitgetModel.SELL_KEYWORD) {
		if !trade.Long {
			trade.TransactionType = trading.OPENING
		} else {
			trade.TransactionType = trading.CLOSE
		}
	} else {
		return nil, c.buildConvertionError(
			"TradeSide",
			"Open",
			parameters.TradeSide,
			fmt.Sprintf(" does not contain %s, %s, %s or %s",
				bitgetModel.BUY_KEYWORD,
				bitgetModel.SELL_KEYWORD,
				bitgetModel.OPEN_KEYWORD,
				bitgetModel.CLOSE_KEYWORD,
			))
	}

	return trade, nil
}

func (c *FutureTransactionToTradeConverter) buildConvertionError(
	inputField string,
	outputField string,
	value string,
	message string,
) error {
	return &customError.ConversionError{
		InputField:   inputField,
		OutputField:  outputField,
		InputStruct:  "bitgetModel.ApiFutureTransaction",
		OutputStruct: "trading.Trade",
		Message:      fmt.Sprintf("%s %s", value, message),
	}
}
