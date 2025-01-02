package configuration

import (
	"encoding/json"
	"os"
)

type ConfigurationLoaderFromJsonFile[T any] struct {
	filePath string
}

func NewConfigurationLoaderFromJsonFile[T any](
	path string,
) IConfigurationLoader[T] {
	return &ConfigurationLoaderFromJsonFile[T]{
		filePath: path,
	}
}

func (l *ConfigurationLoaderFromJsonFile[T]) Load() (*T, error) {
	file, err := os.Open(l.filePath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	decoder := json.NewDecoder(file)
	configuration := new(T)
	err = decoder.Decode(configuration)
	if err != nil {
		return nil, err
	}

	return configuration, nil
}
