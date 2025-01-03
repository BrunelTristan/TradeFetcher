package bitget

type ApiFeeDetail struct {
	FeesValue string `json:"totalFee"`
	FeeToken  string `json:"feeCoin"`
}

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
