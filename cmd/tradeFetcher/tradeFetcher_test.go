package main

import (
	"testing"
	"tradeFetcher/model/configuration"
)

func TestMain(t *testing.T) {
	main()
}

func TestRunForVersion(t *testing.T) {
	run(true, nil)
}

func TestRunForFetchingWithoutConfigFile(t *testing.T) {
	conf := configuration.CmdLineConfiguration{}

	run(false, &conf)
}

func TestRunForFetching(t *testing.T) {
	conf := configuration.CmdLineConfiguration{ConfigFilePath: "/src/integrationTests/files/globalConfig.json"}

	run(false, &conf)
}
