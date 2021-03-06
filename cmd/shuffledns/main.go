package main

import (
	"github.com/projectdiscovery/gologger"
	"github.com/0xrishabh/shuffledns/pkg/runner"
)

func main() {
	// Parse the command line flags and read config files
	options := runner.ParseOptions()

	runner, err := runner.New(options)
	if err != nil {
		gologger.Fatalf("Could not create runner: %s\n", err)
	}

	runner.RunEnumeration()
	runner.Close()
}
