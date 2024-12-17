package bitget

type ApiSpotFill struct {
	Symbol string
	Side   string
	Price  string
	Size   string
}

type ApiSpotGetFills struct {
	ApiResponse
	Data []*ApiSpotFill
}
