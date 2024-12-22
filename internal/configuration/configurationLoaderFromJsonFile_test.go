package configuration

import (
	"fmt"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

type fakeConfig struct {
	Number int
	Label  string
}

func TestNewConfigurationLoaderFromJsonFile(t *testing.T) {
	object := NewConfigurationLoaderFromJsonFile[fakeConfig]("")

	assert.NotNil(t, object)
}

func TestLoadWithUnexistingFile(t *testing.T) {
	loader := NewConfigurationLoaderFromJsonFile[fakeConfig]("/fake/path/that/doesnt/exist")

	conf, err := loader.Load()

	assert.Nil(t, conf)
	assert.NotNil(t, err)
}

func TestLoadWithNonJsonFile(t *testing.T) {
	filePath := "/tmp/testfiles/TestLoadWithNonJsonFile.json"
	_ = os.MkdirAll("/tmp/testfiles/", os.ModePerm)
	_ = os.WriteFile(filePath, []byte("NotJson:{"), os.ModePerm)

	loader := NewConfigurationLoaderFromJsonFile[fakeConfig](filePath)

	conf, err := loader.Load()

	assert.Nil(t, conf)
	assert.NotNil(t, err)
}

func TestLoadWithJsonFile(t *testing.T) {
	filePath := "/tmp/testfiles/TestLoadWithJsonFile.json"
	_ = os.MkdirAll("/tmp/testfiles/", os.ModePerm)
	_ = os.WriteFile(filePath, []byte("{\"Number\":452,\"Label\":\"Unexpected isn't it ?\"}"), os.ModePerm)

	loader := NewConfigurationLoaderFromJsonFile[fakeConfig](filePath)

	conf, err := loader.Load()

	assert.NotNil(t, conf)
	assert.Nil(t, err)

	fmt.Println(err)

	assert.Equal(t, 452, conf.Number)
	assert.Equal(t, "Unexpected isn't it ?", conf.Label)
}
