package runner

import (
	"lukso-cli/config"
	"lukso-cli/runner/pandora"
)

func Start(arg string, LuksoSettings *config.LuksoValues) {
	switch arg {
	case "pandora":
		pandora.Start(LuksoSettings)
	}
}

func Action(cmd string, arg string, LuksoSettings *config.LuksoValues) {

	switch cmd {
	case "start":
		Start(arg, LuksoSettings)
	}

}
