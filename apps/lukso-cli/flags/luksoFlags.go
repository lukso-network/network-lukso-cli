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
			Value:       "",
			Usage:       "Picks which setup to use",
			Destination: &FlagValues.Network,
		},
		cli.BoolFlag{
			Name: "l15-prod",
			Usage: "Alias for --network l15-prod",
		},
		cli.BoolFlag{
			Name: "l15-staging",
			Usage: "Alias for --network l15-staging",
		},
		cli.BoolFlag{
			Name: "l15-dev",
			Usage: "Alias for --network l15-dev",
		},
	}
	return luksoFlags
}