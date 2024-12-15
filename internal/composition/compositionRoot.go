package composition

import (
	"reflect"
	"tradeFetcher/internal/fetcher"
	"tradeFetcher/internal/formatter"
	"tradeFetcher/internal/processUnit"
	"tradeFetcher/model/configuration"
)

type CompositionRoot struct {
	singletons           map[reflect.Type]interface{}
	startupConfiguration *configuration.CmdLineConfiguration
}

func NewCompositionRoot(conf *configuration.CmdLineConfiguration) *CompositionRoot {
	return &CompositionRoot{
		singletons:           make(map[reflect.Type]interface{}),
		startupConfiguration: conf,
	}
}

func (c CompositionRoot) Build() {
}

func (c CompositionRoot) ComposeFetcher() fetcher.IFetcher {
	return fetcher.NewBitgetFetcher(fetcher.NewFakeFetcher())
}

func (c CompositionRoot) ComposeProcessUnit() processUnit.IProcessUnit {
	return processUnit.NewTradeDisplayer(
		formatter.NewCsvTradeFormatter())
}
