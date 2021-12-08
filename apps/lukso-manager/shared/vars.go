package shared

import (
	"net"
	"net/http"

	"github.com/boltdb/bolt"
)

var LuksoHomeDir = ""
var BinaryDir = ""
var NetworkDir = ""
var OutboundIP net.IP
var DataDir = "datadirs"
var SettingsDB *bolt.DB

var LUKSO_GITHUB = "https://api.github.com/repos/lukso-network/"

func GetDataDir(network string, client string) string {
	return NetworkDir + network + "/" + DataDir + "/" + client
}

func GetNetworkDir(network string) string {
	return NetworkDir + network
}

func Contains(s []string, e string) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func HandleError(err error, w http.ResponseWriter) {
	if err != nil {
		w.Header().Set("Content-Type", "application/json")
		w.Write([]byte(err.Error()))
		return
	}
}
