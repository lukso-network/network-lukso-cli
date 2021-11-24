package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func initFlags() {

	app := cli.NewApp()
	app.Name = "LUKSO CLI"
	app.Usage = "Tool for managing LUKSO node"
	app.UsageText = "lukso <command> <argument> [--flags]"

	luksoFlags := []cli.Flag{
		cli.StringFlag{
			Name: "pandora",
		},
	}

	app.Commands = []cli.Command{
		{
			Name:      "start",
			Usage:     "Starts up the client",
			UsageText: "lukso",
			Flags:     luksoFlags,
		},
	}

	err := app.Run(os.Args)
	if err != nil {
		log.Fatal(err)
	}

}
