package main

import (
	"flag"
	"fmt"
	"tradeFetcher/internal/composition"
	"tradeFetcher/model/configuration"
)

const (
	major = 2
	minor = 2
	patch = 0
)

func readFlags() (shouldDisplayVersion bool, conf *configuration.CmdLineConfiguration) {
	conf = &configuration.CmdLineConfiguration{}
	showVersion := flag.Bool("v", false, "display version")
	flag.StringVar(&conf.ConfigFilePath, "cfg", "", "file path to configuration file")

	flag.Parse()

	shouldDisplayVersion = *showVersion

	return
}

func run(shouldDisplayVersion bool, conf *configuration.CmdLineConfiguration) {
	if shouldDisplayVersion {
		displayVersion()
	} else {
		launch(conf)
	}
}

func displayVersion() {
	fmt.Printf("V%d.%d.%d\n", major, minor, patch)
}

func launch(conf *configuration.CmdLineConfiguration) {
	root := composition.NewCompositionRoot(conf)

	root.Build()

	orchestrator := root.ComposeOrchestration()

	if orchestrator != nil {
		orchestrator.Orchestrate()
	} else {
		flag.PrintDefaults()
	}
}

func main() {
	shouldDisplayVersion, conf := readFlags()

	run(shouldDisplayVersion, conf)
}
