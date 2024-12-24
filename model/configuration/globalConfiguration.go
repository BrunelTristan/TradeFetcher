package configuration

import (
	"tradeFetcher/model/bitget"
)

type GlobalConfiguration struct {
	BitgetAccount    *bitget.AccountConfiguration
	BitgetSpotAssets []string
}
