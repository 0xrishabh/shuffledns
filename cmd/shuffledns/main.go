package main

import (
	"github.com/projectdiscovery/gologger"
<<<<<<< HEAD
	"github.com/0xrishabh/shuffledns/pkg/runner"
=======
	"github.com/projectdiscovery/0xrishabh/pkg/runner"
>>>>>>> a56d138517f75a00780cc7702ae84ce8555500ff
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
