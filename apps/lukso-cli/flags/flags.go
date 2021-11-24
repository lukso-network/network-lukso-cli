package flags

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func InitFlags() {

	app := cli.NewApp()
	app.Name = "LUKSO CLI"
	app.Usage = "Tool for managing LUKSO node"
	app.UsageText = "lukso <command> [argument] [--flags]"
	app.Flags = getLuksoFlags()

	app.Commands = []cli.Command{
		getStartCommand(),
		getStopCommand(),
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
