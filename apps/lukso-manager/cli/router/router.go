package router

import "lukso/apps/lukso-manager/runner"

func Handle(cmd string, arg string) {
	switch cmd {
	case "start":
		println("Starting")
		switch arg {
		case "all":
			runner.StartClients()
		}
	}
}
