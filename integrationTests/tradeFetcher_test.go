package integrationTest

import (
	"testing"
	"tradeFetcher/internal/composition"
	"tradeFetcher/model/configuration"
)

func TestMain(t *testing.T) {
	conf := &configuration.CmdLineConfiguration{}

	root := composition.NewCompositionRoot(conf)

	root.Build()

	fetcher := root.ComposeFetcher()
	processor := root.ComposeProcessUnit()

	trades := fetcher.FetchLastTrades()
	processor.ProcessTrades(trades)
}
