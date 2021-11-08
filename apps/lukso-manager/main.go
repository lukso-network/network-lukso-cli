package main

import (
	"lukso/downloader"
	"lukso/metrics"
	"lukso/runner"
	"lukso/shared"
	"lukso/validator"
	"lukso/webserver"
	"os"

	"github.com/gorilla/mux"
)

func init() {
	userHomeDir, errHome := os.UserHomeDir()
	if errHome != nil {
		panic("Can not get the UserHomeDir")
	}

	shared.UserHomeDir = userHomeDir
	shared.BinaryDir = userHomeDir + "/.lukso/downloads/"
	shared.NetworkDir = userHomeDir + "/.lukso/networks/"
}

func main() {
	app := webserver.App{
		Router: mux.NewRouter(),
	}

	app.Router.Methods("GET").Path("/health").HandlerFunc(metrics.VanguardMetrics)

	app.Router.Methods("GET").Path("/vanguard/metrics").HandlerFunc(metrics.VanguardMetrics)
	app.Router.Methods("GET").Path("/validator/metrics").HandlerFunc(metrics.ValidatorMetrics)
	app.Router.Methods("GET").Path("/pandora/debug/metrics").HandlerFunc(metrics.PandoraMetrics)
	app.Router.Methods("GET").Path("/downloaded-versions").HandlerFunc(downloader.GetDownloadedVersions)
	app.Router.Methods("GET").Path("/available-versions").HandlerFunc(downloader.GetAvailableVersions)

	app.Router.Methods("POST").Path("/update-client").HandlerFunc(downloader.DownloadClient)
	app.Router.Methods("POST").Path("/start-clients").HandlerFunc(runner.StartClients)
	app.Router.Methods("POST").Path("/stop-clients").HandlerFunc(runner.StopClients)
	app.Router.Methods("POST").Path("/launchpad/generate-keys").HandlerFunc(validator.GenerateValidatorKeys)

	app.Start()
}
