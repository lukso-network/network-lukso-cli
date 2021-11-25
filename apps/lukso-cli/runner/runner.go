package runner

import (
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

func Action(cmd string, arg string, LuksoSettings *config.LuksoValues) {

	switch cmd {
	case "start":
		Start(arg, LuksoSettings)
	}

}
