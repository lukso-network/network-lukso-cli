package main

import (
	"lukso/apps/lukso-manager/src/downloader"
	"lukso/apps/lukso-manager/src/metrics"
	"lukso/apps/lukso-manager/src/webserver"

	"github.com/gorilla/mux"
)

func main() {
	app := webserver.App{
		Router: mux.NewRouter(),
	}

	app.Router.
		Methods("GET").
		Path("/vanguard/metrics").
		HandlerFunc(metrics.VanguardMetrics)

	app.Router.
		Methods("GET").
		Path("/validator/metrics").
		HandlerFunc(metrics.ValidatorMetrics)

	app.Router.
		Methods("GET").
		Path("/pandora/debug/metrics").
		HandlerFunc(metrics.PandoraMetrics)

	app.Router.
		Methods("GET").
		Path("/downloaded-versions").
		HandlerFunc(downloader.GetDownloadedVersions)

	app.Router.
		Methods("GET").
		Path("/available-versions").
		HandlerFunc(downloader.GetAvailableVersions)

	app.Router.
		Methods("POST").
		Path("/update-client").
		HandlerFunc(downloader.DownloadClient)

	app.Start()
}
