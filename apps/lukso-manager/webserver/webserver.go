package webserver

import (
	"log"
	"lukso/apps/lukso-manager/downloader"
	"lukso/apps/lukso-manager/metrics"
	"lukso/apps/lukso-manager/runner"
	"lukso/apps/lukso-manager/settings"
	"lukso/apps/lukso-manager/setup"
	"lukso/apps/lukso-manager/validator"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/rs/cors"
)

type App struct {
	Router *mux.Router
}

func (app *App) Start(addr string) {
	handler := cors.Default().Handler(app.Router)

	log.Fatal(http.ListenAndServe(addr, handler))
}

func StartAPIServer() {
	app := App{
		Router: mux.NewRouter(),
	}

	app.Router.Methods("GET").Path("/health").HandlerFunc(metrics.Health)

	app.Router.Methods("GET").Path("/vanguard/metrics").HandlerFunc(metrics.VanguardMetrics)
	app.Router.Methods("GET").Path("/validator/metrics").HandlerFunc(metrics.ValidatorMetrics)
	app.Router.Methods("GET").Path("/pandora/debug/metrics").HandlerFunc(metrics.PandoraMetrics)
	app.Router.Methods("GET").Path("/pandora/peers-over-time").HandlerFunc(metrics.GetPandoraPeersOverTime)
	app.Router.Methods("GET").Path("/vanguard/peers-over-time").HandlerFunc(metrics.GetVanguardPeersOverTime)
	app.Router.Methods("GET").Path("/downloaded-versions").HandlerFunc(downloader.GetDownloadedVersions)
	app.Router.Methods("GET").Path("/available-versions").HandlerFunc(downloader.GetAvailableVersions)
	app.Router.Methods("GET").Path("/deposit-data").HandlerFunc(validator.GetDepositData)

	app.Router.Methods("POST").Path("/initial-setup").HandlerFunc(setup.Setup)
	app.Router.Methods("POST").Path("/update-client").HandlerFunc(downloader.DownloadClient)
	app.Router.Methods("POST").Path("/start-clients").HandlerFunc(runner.StartClients)
	app.Router.Methods("POST").Path("/stop-clients").HandlerFunc(runner.StopClients)
	app.Router.Methods("POST").Path("/launchpad/generate-keys").HandlerFunc(validator.GenerateValidatorKeys)
	app.Router.Methods("POST").Path("/launchpad/import-keys").HandlerFunc(validator.ImportValidatorKeys)
	app.Router.Methods("POST").Path("/launchpad/reset-validator").HandlerFunc(validator.ResetValidator)
	app.Router.Methods("POST").Path("/settings").HandlerFunc(settings.SaveSettingsEndpoint)
	app.Router.Methods("GET").Path("/settings").HandlerFunc(settings.GetSettingsEndpoint)

	go func() { app.Start(":3000") }()
}

func StartGUIServer() {
	app := App{
		Router: mux.NewRouter(),
	}

	app.Router.Handle("/", http.FileServer(http.Dir("../../../dist/apps/lukso-gui")))

	go func() { app.Start(":4000") }()
}
