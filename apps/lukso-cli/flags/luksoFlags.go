package flags
import (

	"github.com/urfave/cli"

	"lukso-cli/config"
)

var FlagValues config.LuksoValues
var luksoFlags []cli.Flag

func getLuksoFlags () []cli.Flag {
	luksoFlags = []cli.Flag{
		cli.StringFlag{
			Name:        "config",
			Value:       "",
			Usage:       "Loads config file",
			Destination: &FlagValues.Config,
		},
		cli.StringFlag{
			Name:        "network",
			Value:       "l15",
			Usage:       "Picks which setup to use",
			Destination: &FlagValues.Network,
		},
	}
	return luksoFlags
}