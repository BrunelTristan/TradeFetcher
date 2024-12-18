package main

import (
	"flag"
	"fmt"
	"tradeFetcher/internal/composition"
	"tradeFetcher/model/configuration"
)

const (
	major = 0
	minor = 0
	patch = 0
)

func readFlags() (shouldDisplayVersion bool, conf *configuration.CmdLineConfiguration) {
	conf = &configuration.CmdLineConfiguration{}
	showVersion := flag.Bool("v", false, "display version")

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

	fetcher := root.ComposeFetcher()
	processor := root.ComposeProcessUnit()

	trades := fetcher.FetchLastTrades()
	processor.ProcessTrades(trades)
}

func main() {
	shouldDisplayVersion, conf := readFlags()

	run(shouldDisplayVersion, conf)
}
