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

func TestRunForFetching(t *testing.T) {
	conf := configuration.CmdLineConfiguration{}

	run(false, &conf)
}
