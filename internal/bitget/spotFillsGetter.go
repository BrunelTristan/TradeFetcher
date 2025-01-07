package bitget

import (
	"tradeFetcher/internal/common"
	"tradeFetcher/internal/externalTools"
	bitgetModel "tradeFetcher/model/bitget"
	customError "tradeFetcher/model/error"
)

type SpotFillsGetter struct {
	apiQuery     common.IQuery[bitgetModel.ApiQueryParameters]
	routeBuilder externalTools.IApiRouteBuilder
}

func NewSpotFillsGetter(
	aQuery common.IQuery[bitgetModel.ApiQueryParameters],
	rBuilder externalTools.IApiRouteBuilder,
) common.IQuery[bitgetModel.SpotGetFillQueryParameters] {
	return &SpotFillsGetter{
		apiQuery:     aQuery,
		routeBuilder: rBuilder,
	}
}

func (g *SpotFillsGetter) Get(parameters *bitgetModel.SpotGetFillQueryParameters) (interface{}, error) {
	if parameters == nil {
		return nil, &customError.RestApiError{HttpCode: 999}
	}

	route := g.routeBuilder.BuildRoute(
		[]string{bitgetModel.SPOT_ROOT_ROUTE, bitgetModel.SPOT_GET_FILLS_SUB_ROUTE},
		map[string]string{"symbol": parameters.Symbol},
	)

	return g.apiQuery.Get(&bitgetModel.ApiQueryParameters{
		Route: route,
	})
}
