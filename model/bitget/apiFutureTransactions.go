package bitget

const OPEN_KEYWORD = "open"
const CLOSE_KEYWORD = "close"

type ApiFutureTransaction struct {
	Symbol     string
	Side       string
	Price      string
	Size       string `json:"baseVolume"`
	LastUpdate string `json:"cTime"`
	TradeSide  string
	FeeDetail  []*ApiFeeDetail
}

type ApiFutureTransactionsList struct {
	FillList []*ApiFutureTransaction
}

type ApiFutureTransactions struct {
	ApiResponse
	Data *ApiFutureTransactionsList
}

func (t *ApiFutureTransactions) GetCode() string {
	return t.ApiResponse.Code
}

func (t *ApiFutureTransactions) GetMessage() string {
	return t.ApiResponse.Message
}

func (t *ApiFutureTransactions) GetList() []*ApiFutureTransaction {
	return t.Data.FillList
}
