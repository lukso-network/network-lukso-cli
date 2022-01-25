package runner

import (
	"fmt"
	"lukso-cli/config"
	"lukso-cli/runner/orchestrator"
	"lukso-cli/runner/pandora"
	"lukso-cli/runner/validator"
	"lukso-cli/runner/vanguard"
)

func Start(arg string, LuksoSettings *config.LuksoValues) {

	switch arg {
	case "all":
		orchestrator.Start(LuksoSettings)
		pandora.Start(LuksoSettings)
		vanguard.Start(LuksoSettings)
		if LuksoSettings.Validate {
			validator.Start(LuksoSettings)
		}

	case "orchestrator":
		orchestrator.Start(LuksoSettings)

	case "pandora":
		pandora.Start(LuksoSettings)

	case "vanguard":
		vanguard.Start(LuksoSettings)

	case "validator":
		validator.Start(LuksoSettings)
	}
}

func Stop(arg string, LuksoSettings *config.LuksoValues) error {
	return nil
}

func Reset(arg string, LuksoSettings *config.LuksoValues) error {
	return nil
}

func Version() {
	fmt.Println("v0.0.1")
}

func Action(cmd string, arg string, LuksoSettings *config.LuksoValues) {
	switch cmd {
	case "start":
		Start(arg, LuksoSettings)
	case "stop":
		Stop(arg, LuksoSettings)
	case "reset":
		Reset(arg, LuksoSettings)
	case "version":
		Version()
	}

}
