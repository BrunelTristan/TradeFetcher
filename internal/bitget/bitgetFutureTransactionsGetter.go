package bitget

import (
	"tradeFetcher/internal/common"
	"tradeFetcher/internal/externalTools"
	bitgetModel "tradeFetcher/model/bitget"
)

type BitgetFutureTransactionsGetter struct {
	apiQuery     common.IQuery[bitgetModel.ApiQueryParameters]
	routeBuilder externalTools.IApiRouteBuilder
}

func NewBitgetFutureTransactionsGetter(
	aQuery common.IQuery[bitgetModel.ApiQueryParameters],
	rBuilder externalTools.IApiRouteBuilder,
) common.IQuery[bitgetModel.FutureTransactionsQueryParameters] {
	return &BitgetFutureTransactionsGetter{
		apiQuery:     aQuery,
		routeBuilder: rBuilder,
	}
}

func (g *BitgetFutureTransactionsGetter) Get(parameters *bitgetModel.FutureTransactionsQueryParameters) (interface{}, error) {
	route := g.routeBuilder.BuildRoute(
		[]string{bitgetModel.FUTURE_ROOT_ROUTE, bitgetModel.FUTURE_GET_TRANSACTION_SUB_ROUTE},
		map[string]string{"productType": "USDT-FUTURES"},
	)

	return g.apiQuery.Get(&bitgetModel.ApiQueryParameters{
		Route: route,
	})
}
