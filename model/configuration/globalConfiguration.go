package configuration

import (
	"tradeFetcher/model/bitget"
)

type GlobalConfiguration struct {
	TradeHistoryFilePath string
	BitgetAccount        *bitget.AccountConfiguration
	BitgetSpotAssets     []string
}
