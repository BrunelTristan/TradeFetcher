package bitget

type ApiSpotFill struct {
	Symbol   string
	Side     string
	Price    string
	Size     string
	Fees     string
	FeeToken string
}

type ApiSpotGetFills struct {
	ApiResponse
	Data []*ApiSpotFill
}
