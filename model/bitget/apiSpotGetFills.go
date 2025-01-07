package bitget

const BUY_KEYWORD = "buy"
const SELL_KEYWORD = "sell"

type ApiSpotFill struct {
	Symbol     string
	Side       string
	Price      string `json:"priceAvg"`
	Size       string
	LastUpdate string `json:"uTime"`
	FeeDetail  *ApiFeeDetail
}

type ApiSpotGetFills struct {
	ApiResponse
	Data []*ApiSpotFill
}

func (f *ApiSpotGetFills) GetCode() string {
	return f.ApiResponse.Code
}

func (f *ApiSpotGetFills) GetMessage() string {
	return f.ApiResponse.Message
}

func (f *ApiSpotGetFills) GetList() []*ApiSpotFill {
	return f.Data
}
