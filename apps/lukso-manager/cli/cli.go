package cli

import (
	"log"
	"os"

	"github.com/urfave/cli/v2"

	"lukso/apps/lukso-manager/runner"
)

var Cmd string
var Arg string

var API bool
var GUI bool

func Init() {

	app := cli.NewApp()
	app.Name = "LUKSO CLI"
	app.Usage = "Tool for managing LUKSO node"
	app.UsageText = "lukso <command> [argument] [--flags]"
	app.Flags = getLuksoFlags()
	app.EnableBashCompletion = true

	app.Commands = []*cli.Command{
		getStartCommand(),
		getStopCommand(),
		getVersionCommand(),
	}

	app.After = func(c *cli.Context) error {
		LoadFlags(c)
		return nil
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

	runner.HandleCli(Cmd, Arg)

}
