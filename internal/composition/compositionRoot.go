package composition

import (
	"reflect"
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
