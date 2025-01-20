package main

import (
	"syscall"
	"testing"
	"time"
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

	go func() {
		run(false, &conf)

		time.Sleep(10 * time.Millisecond)

		_ = syscall.Kill(syscall.Getpid(), syscall.SIGINT)
	}()

	time.Sleep(100 * time.Millisecond)
}
