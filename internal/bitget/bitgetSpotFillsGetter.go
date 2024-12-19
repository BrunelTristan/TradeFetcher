package bitget

import (
	"tradeFetcher/internal/externalTools"
	bitgetModel "tradeFetcher/model/bitget"
	customError "tradeFetcher/model/error"
)

type BitgetSpotFillsGetter struct {
	apiCommand externalTools.ICommand[bitgetModel.ApiCommandParameters]
}

func NewBitgetSpotFillsGetter(
	aCommand externalTools.ICommand[bitgetModel.ApiCommandParameters],
) externalTools.ICommand[bitgetModel.SpotGetFillCommandParameters] {
	return &BitgetSpotFillsGetter{
		apiCommand: aCommand,
	}
}

func (g *BitgetSpotFillsGetter) Get(parameters *bitgetModel.SpotGetFillCommandParameters) (interface{}, error) {
	if parameters == nil {
		return nil, &customError.RestApiError{HttpCode: 999}
	}

	return g.apiCommand.Get(&bitgetModel.ApiCommandParameters{
		Route: "/api/v2/spot/trade/fills?symbol=ETHUSDT",
	})
}
