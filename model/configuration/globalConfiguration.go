package configuration

import (
	"tradeFetcher/model/bitget"
)

type GlobalConfiguration struct {
	TradeHistoryFilePath              string
	OrchestrationPeriodicityInSeconds int64
	BitgetAccount                     *bitget.AccountConfiguration
	BitgetSpotAssets                  []string
}
