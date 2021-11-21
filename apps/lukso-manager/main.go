package main

import (
	"log"
	"lukso-manager/downloader"
	"lukso-manager/metrics"
	"lukso-manager/runner"
	"lukso-manager/settings"
	"lukso-manager/shared"
	"lukso-manager/validator"
	"lukso-manager/webserver"
	"net"
	"os"

	"github.com/boltdb/bolt"
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

	db, err := bolt.Open(shared.LuksoHomeDir+"/settings.db", 0640, nil)
	if err != nil {
		log.Fatal(err)
	}
	shared.SettingsDB = db
	shared.OutboundIP = getOutboundIP().String()

}

func getOutboundIP() net.IP {
	conn, err := net.Dial("udp", "8.8.8.8:80")
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	localAddr := conn.LocalAddr().(*net.UDPAddr)

	return localAddr.IP
}

func main() {
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

	app.Router.Methods("POST").Path("/update-client").HandlerFunc(downloader.DownloadClient)
	app.Router.Methods("POST").Path("/start-clients").HandlerFunc(runner.StartClients)
	app.Router.Methods("POST").Path("/stop-clients").HandlerFunc(runner.StopClients)
	app.Router.Methods("POST").Path("/launchpad/generate-keys").HandlerFunc(validator.GenerateValidatorKeys)
	app.Router.Methods("POST").Path("/launchpad/import-keys").HandlerFunc(validator.ImportValidatorKeys)
	app.Router.Methods("POST").Path("/settings").HandlerFunc(settings.SaveSettingsEndpoint)
	app.Router.Methods("GET").Path("/settings").HandlerFunc(settings.GetSettingsEndpoint)

	app.Start()
}
