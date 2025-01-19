package composition

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"tradeFetcher/model/configuration"
)

func TestBuild(t *testing.T) {
	conf := &configuration.CmdLineConfiguration{}
	root := NewCompositionRoot(conf)

	root.Build()
}

func TestComposeFetcherWithoutValidConfig(t *testing.T) {
	conf := &configuration.CmdLineConfiguration{ConfigFilePath: "nothing"}
	root := NewCompositionRoot(conf)

	root.Build()

	fetcher := root.ComposeFetcher()

	assert.Nil(t, fetcher)
}

func TestComposeFetcherWithValidConfig(t *testing.T) {
	conf := &configuration.CmdLineConfiguration{ConfigFilePath: "/src/integrationTests/files/globalConfig.json"}
	root := NewCompositionRoot(conf)

	root.Build()

	fetcher := root.ComposeFetcher()

	assert.NotNil(t, fetcher)
}

func TestComposeProcessUnitWithoutValidConfig(t *testing.T) {
	conf := &configuration.CmdLineConfiguration{}
	root := NewCompositionRoot(conf)

	root.Build()

	units := root.ComposeProcessUnit()

	assert.Nil(t, units)
}

func TestComposeProcessUnitWithValidConfig(t *testing.T) {
	conf := &configuration.CmdLineConfiguration{ConfigFilePath: "/src/integrationTests/files/globalConfig.json"}
	root := NewCompositionRoot(conf)

	root.Build()

	units := root.ComposeProcessUnit()

	assert.NotNil(t, units)
	assert.LessOrEqual(t, 2, len(units))
}
