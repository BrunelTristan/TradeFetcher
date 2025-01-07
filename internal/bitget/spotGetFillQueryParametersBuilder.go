package bitget

import (
	"tradeFetcher/internal/common"
	bitgetModel "tradeFetcher/model/bitget"
)

type SpotGetFillQueryParametersBuilder struct {
	symbol string
}

func NewSpotGetFillQueryParametersBuilder(asset string) common.IQueryParametersBuilder[bitgetModel.SpotGetFillQueryParameters] {
	return &SpotGetFillQueryParametersBuilder{
		symbol: asset,
	}
}

func (b *SpotGetFillQueryParametersBuilder) BuildQueryParameters() (*bitgetModel.SpotGetFillQueryParameters, error) {
	return &bitgetModel.SpotGetFillQueryParameters{Symbol: b.symbol}, nil
}
