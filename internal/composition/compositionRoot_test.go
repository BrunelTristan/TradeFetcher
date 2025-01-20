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

func TestComposeOrchestrationWithoutValidConfig(t *testing.T) {
	conf := &configuration.CmdLineConfiguration{}
	root := NewCompositionRoot(conf)

	root.Build()

	orchestrator := root.ComposeOrchestration()

	assert.Nil(t, orchestrator)
}

func TestComposeOrchestrationWithValidConfig(t *testing.T) {
	conf := &configuration.CmdLineConfiguration{ConfigFilePath: "/src/integrationTests/files/globalConfig.json"}
	root := NewCompositionRoot(conf)

	root.Build()

	orchestrator := root.ComposeOrchestration()

	assert.NotNil(t, orchestrator)
}
