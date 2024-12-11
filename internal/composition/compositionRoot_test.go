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

func TestComposeFetcher(t *testing.T) {
	conf := &configuration.CmdLineConfiguration{}
	root := NewCompositionRoot(conf)

	root.Build()

	fetcher := root.ComposeFetcher()

	assert.NotNil(t, fetcher)
}
