package bitget

type ApiFeeDetail struct {
	FeesValue string `json:"totalFee"`
	FeeToken  string `json:"feeCoin"`
}
