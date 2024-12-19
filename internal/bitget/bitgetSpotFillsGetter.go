package bitget

import (
	"tradeFetcher/internal/externalTools"
	bitgetModel "tradeFetcher/model/bitget"
	customError "tradeFetcher/model/error"
)

type BitgetSpotFillsGetter struct {
	apiCommand   externalTools.ICommand[bitgetModel.ApiCommandParameters]
	routeBuilder externalTools.IApiRouteBuilder
}

func NewBitgetSpotFillsGetter(
	aCommand externalTools.ICommand[bitgetModel.ApiCommandParameters],
	rBuilder externalTools.IApiRouteBuilder,
) externalTools.ICommand[bitgetModel.SpotGetFillCommandParameters] {
	return &BitgetSpotFillsGetter{
		apiCommand:   aCommand,
		routeBuilder: rBuilder,
	}
}

func (g *BitgetSpotFillsGetter) Get(parameters *bitgetModel.SpotGetFillCommandParameters) (interface{}, error) {
	if parameters == nil {
		return nil, &customError.RestApiError{HttpCode: 999}
	}

	route := g.routeBuilder.BuildRoute(
		[]string{bitgetModel.SPOT_ROOT_ROUTE, bitgetModel.SPOT_GET_FILLS_SUB_ROUTE},
		map[string]string{"symbol": parameters.Symbol},
	)

	return g.apiCommand.Get(&bitgetModel.ApiCommandParameters{
		Route: route,
	})
}
