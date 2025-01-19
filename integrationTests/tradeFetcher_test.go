package integrationTest

import (
	"github.com/stretchr/testify/assert"
	"testing"
	"time"
	"tradeFetcher/internal/composition"
	"tradeFetcher/model/configuration"
)

func TestMain(t *testing.T) {
	conf := &configuration.CmdLineConfiguration{ConfigFilePath: "/src/integrationTests/files/globalConfig.json"}

	root := composition.NewCompositionRoot(conf)

	root.Build()

	orchestrator := root.ComposeOrchestration()

	assert.NotNil(t, orchestrator)

	orchestrator.Orchestrate()

	time.Sleep(1 * time.Second)

	orchestrator.EndOrchestration()
}
