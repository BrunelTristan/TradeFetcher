package bitget

import (
	"tradeFetcher/internal/common"
	"tradeFetcher/internal/externalTools"
	bitgetModel "tradeFetcher/model/bitget"
)

type FutureTaxTransactionsGetter struct {
	apiQuery     common.IQuery[bitgetModel.ApiQueryParameters]
	routeBuilder externalTools.IApiRouteBuilder
}

func NewFutureTaxTransactionsGetter(
	aQuery common.IQuery[bitgetModel.ApiQueryParameters],
	rBuilder externalTools.IApiRouteBuilder,
) common.IQuery[bitgetModel.FutureTaxTransactionsQueryParameters] {
	return &FutureTaxTransactionsGetter{
		apiQuery:     aQuery,
		routeBuilder: rBuilder,
	}
}

func (g *FutureTaxTransactionsGetter) Get(parameters *bitgetModel.FutureTaxTransactionsQueryParameters) (interface{}, error) {
	route := g.routeBuilder.BuildRoute(
		[]string{bitgetModel.TAX_ROOT_ROUTE, bitgetModel.TAX_GET_FUTURE_TRANSACTION_SUB_ROUTE},
		map[string]string{},
	)

	return g.apiQuery.Get(&bitgetModel.ApiQueryParameters{
		Route: route,
	})
}
