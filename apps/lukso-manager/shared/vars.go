package shared

import "github.com/boltdb/bolt"

var LuksoHomeDir = ""
var BinaryDir = ""
var NetworkDir = ""
var OutboundIP = ""
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
