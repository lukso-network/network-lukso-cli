package main

import (
	"fmt"
	"log"
	"lukso/apps/lukso-manager/cli"
	"lukso/apps/lukso-manager/shared"
	"lukso/apps/lukso-manager/webserver"
	"os"

	"github.com/boltdb/bolt"
	externalip "github.com/glendc/go-external-ip"
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

	if cli.API || cli.GUI {
		webserver.StartAPIServer()
	}

	if cli.GUI {
		//webserver.StartGUIServer()
	}
}
