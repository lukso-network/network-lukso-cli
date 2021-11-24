package main

import (
	"lukso-cli/config"
	"lukso-cli/flags"
)

func main() {
	flags.InitFlags()

	if flags.FlagValues.Config != "" {
		println("works")
		config.LoadConfig(flags.FlagValues.Config)
	}

	// var LuksoSettings config.LuksoValues

}
