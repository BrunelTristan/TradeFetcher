package bitget

type ApiFutureTransaction struct {
	Symbol     string
	Side       string
	Price      string
	Size       string `json:"baseVolume"`
	LastUpdate string `json:"cTime"`
	FeeDetail  []*ApiFeeDetail
}

type ApiFutureTransactionsList struct {
	FillList []*ApiFutureTransaction
}

type ApiFutureTransactions struct {
	ApiResponse
	Data *ApiFutureTransactionsList
}
