package integrationTest

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tradeFetcher/internal/composition"
	"tradeFetcher/model/configuration"
)

func TestMain(t *testing.T) {
	conf := &configuration.CmdLineConfiguration{ConfigFilePath: "/src/integrationTests/files/globalConfig.json"}

	root := composition.NewCompositionRoot(conf)

	root.Build()

	fetcher := root.ComposeFetcher()
	processors := root.ComposeProcessUnit()

	trades, err := fetcher.FetchLastTrades()
	assert.Nil(t, err)

	assert.LessOrEqual(t, 1, len(processors))
	for _, processor := range processors {
		err = processor.ProcessTrades(trades)
		assert.Nil(t, err)
	}
}
