package main

import (
	"fmt"
	"log"
	"lukso/apps/lukso-manager/cli"
	"lukso/apps/lukso-manager/downloader"
	"lukso/apps/lukso-manager/metrics"
	"lukso/apps/lukso-manager/runner"
	"lukso/apps/lukso-manager/settings"
	"lukso/apps/lukso-manager/setup"
	"lukso/apps/lukso-manager/shared"
	"lukso/apps/lukso-manager/validator"
	"lukso/apps/lukso-manager/webserver"
	"os"

	"github.com/boltdb/bolt"
	externalip "github.com/glendc/go-external-ip"
	"github.com/gorilla/mux"
)

func init() {
	userHomeDir, errHome := os.UserHomeDir()
	if errHome != nil {
		panic("Can not get the UserHomeDir")
	}

	shared.LuksoHomeDir = userHomeDir + "/.lukso"
	shared.BinaryDir = shared.LuksoHomeDir + "/binaries/"
	shared.NetworkDir = shared.LuksoHomeDir + "/networks/"

	os.MkdirAll(shared.LuksoHomeDir, 0775)
	os.MkdirAll(shared.BinaryDir, 0775)
	os.MkdirAll(shared.NetworkDir, 0775)

	db, err := bolt.Open(shared.LuksoHomeDir+"/lukso-manager.db", 0640, nil)
	if err != nil {
		log.Fatal(err)
	}
	shared.SettingsDB = db

	consensus := externalip.DefaultConsensus(nil, nil)
	consensus.UseIPProtocol(4)

	// Get your IP,
	// which is never <nil> when err is <nil>.
	ip, err := consensus.ExternalIP()
	if err == nil {
		fmt.Println(ip.String()) // print IPv4/IPv6 in string format
	}
	shared.OutboundIP = ip

}

func main() {

	cli.Init()

	if false {

		app := webserver.App{
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

		app.Start()
	}
}
