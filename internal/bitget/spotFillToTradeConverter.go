package bitget

import (
	"fmt"
	"strconv"
	"tradeFetcher/internal/converter"
	bitgetModel "tradeFetcher/model/bitget"
	customError "tradeFetcher/model/error"
	"tradeFetcher/model/trading"
)

type SpotFillToTradeConverter struct {
}

func NewSpotFillToTradeConverter() converter.IStructConverter[bitgetModel.ApiSpotFill, trading.Trade] {
	return &SpotFillToTradeConverter{}
}

func (c *SpotFillToTradeConverter) Convert(parameters *bitgetModel.ApiSpotFill) (*trading.Trade, error) {
	if parameters == nil {
		return nil, nil
	}

	trade := &trading.Trade{Pair: parameters.Symbol, Long: true}

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

	floatVal, err = strconv.ParseFloat(parameters.FeeDetail.FeesValue, 64)
	if err != nil {
		return nil, c.buildConvertionError(
			"FeeDetail.FeesValue",
			"Fees",
			parameters.FeeDetail.FeesValue,
			" is not a float 64")
	}
	trade.Fees = floatVal

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
		trade.Open = true
	} else if parameters.Side == bitgetModel.SELL_KEYWORD {
		trade.Open = false
	} else {
		return nil, c.buildConvertionError(
			"Side",
			"Open",
			parameters.Side,
			fmt.Sprintf(" is not %s or %s", bitgetModel.BUY_KEYWORD, bitgetModel.SELL_KEYWORD))
	}

	return trade, nil
}

func (c *SpotFillToTradeConverter) buildConvertionError(
	inputField string,
	outputField string,
	value string,
	message string,
) error {
	return &customError.ConversionError{
		InputField:   inputField,
		OutputField:  outputField,
		InputStruct:  "bitgetModel.ApiSpotFill",
		OutputStruct: "trading.Trade",
		Message:      fmt.Sprintf("%s %s", value, message),
	}
}
