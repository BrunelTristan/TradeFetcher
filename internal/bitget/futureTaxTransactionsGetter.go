package bitget

import (
	"strconv"
	"time"
	"tradeFetcher/internal/common"
	"tradeFetcher/internal/externalTools"
	bitgetModel "tradeFetcher/model/bitget"
)

const TWENTY_DAYS_IN_SECONDS = 20 * 24 * 3600

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
	currentTimestamp := time.Now().Unix()
	route := g.routeBuilder.BuildRoute(
		[]string{bitgetModel.TAX_ROOT_ROUTE, bitgetModel.TAX_GET_FUTURE_TRANSACTION_SUB_ROUTE},
		map[string]string{
			bitgetModel.END_TIME:   strconv.FormatInt(currentTimestamp, 10) + "000",
			bitgetModel.START_TIME: strconv.FormatInt(currentTimestamp-TWENTY_DAYS_IN_SECONDS, 10) + "000",
		},
	)

	return g.apiQuery.Get(&bitgetModel.ApiQueryParameters{
		Route: route,
	})
}
