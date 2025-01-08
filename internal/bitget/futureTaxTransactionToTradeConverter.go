package bitget

import (
	"fmt"
	"strconv"
	"tradeFetcher/internal/converter"
	bitgetModel "tradeFetcher/model/bitget"
	customError "tradeFetcher/model/error"
	"tradeFetcher/model/trading"
)

type FutureTaxTransactionToTradeConverter struct {
}

func NewFutureTaxTransactionToTradeConverter() converter.IStructConverter[bitgetModel.ApiFutureTaxTransaction, trading.Trade] {
	return &FutureTaxTransactionToTradeConverter{}
}

func (c *FutureTaxTransactionToTradeConverter) Convert(parameters *bitgetModel.ApiFutureTaxTransaction) (*trading.Trade, error) {
	if parameters == nil {
		return nil, nil
	}

	trade := &trading.Trade{Pair: parameters.Symbol}

	floatVal, err := strconv.ParseFloat(parameters.Fee, 64)
	if err != nil {
		return nil, c.buildConvertionError(
			"Fee",
			"Fees",
			parameters.Fee,
			" is not a float 64")
	}
	trade.Fees = -floatVal

	intVal, err := strconv.ParseInt(parameters.Timestamp, 10, 64)
	if err != nil {
		return nil, c.buildConvertionError(
			"Timestamp",
			"ExecutedTimestamp",
			parameters.Timestamp,
			" is not a int 64")
	}
	trade.ExecutedTimestamp = intVal / 1000

	// TODO test TaxType
	trade.TransactionType = trading.FUNDING

	return trade, nil
}

func (c *FutureTaxTransactionToTradeConverter) buildConvertionError(
	inputField string,
	outputField string,
	value string,
	message string,
) error {
	return &customError.ConversionError{
		InputField:   inputField,
		OutputField:  outputField,
		InputStruct:  "bitgetModel.ApiFutureTaxTransaction",
		OutputStruct: "trading.Trade",
		Message:      fmt.Sprintf("%s %s", value, message),
	}
}
