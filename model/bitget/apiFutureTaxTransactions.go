package bitget

type ApiFutureTaxTransaction struct {
	Symbol  string
	TaxType string `json:"futureTaxType"`
	//Amout     string
	Fee       string
	Timestamp string `json:"ts"`
}

type ApiFutureTaxTransactions struct {
	ApiResponse
	Data []*ApiFutureTaxTransaction
}

func (t *ApiFutureTaxTransactions) GetCode() string {
	return t.ApiResponse.Code
}

func (t *ApiFutureTaxTransactions) GetMessage() string {
	return t.ApiResponse.Message
}

func (t *ApiFutureTaxTransactions) GetList() []*ApiFutureTaxTransaction {
	return t.Data
}
