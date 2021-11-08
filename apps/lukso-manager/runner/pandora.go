package runner

import (
	"fmt"
	"io"
	"log"
	"lukso/shared"
	"os"
	"os/exec"
	"strings"
)

func startPandora(version string, network string) {
	client := "pandora"
	datadir := shared.NetworkDir + network + "/datadirs/" + client

	config, err := ReadConfig(network)
	if err != nil {
		fmt.Println("error reading config")
	}

	statsPrefix := ""
	if !(network == "l15-prod") {
		statsPrefix = strings.Split(network, "-")[1] + "."
	}

	args := []string{
		"--datadir=" + datadir,
		"--networkid=" + fmt.Sprint(config.NETWORKID),
		"--port=30405",
		"--http",
		"--http.addr=127.0.0.1",
		"--http.port=8545",
		"--bootnodes=" + config.PANDORABOOTNODES,
		"--ws",
		"--ws.addr=127.0.0.1",
		"--ws.port=8546",
		"--mine",
		"--miner.notify=ws://127.0.0.1:7878,http://127.0.0.1:7877",
		"--miner.etherbase=0x6Af9552d70F943378820edc3095A6bb0279051ff",
		"--miner.gaslimit=80000000",
		"--syncmode=full",
		"--verbosity=4",
		"--nat=extip:46.127.26.82",
		"--metrics",
		"--metrics.expensive",
		"--pprof",
		"--pprof.addr=127.0.0.1",
		"--ethstats=l15-rryter-gui:6Tcpc53R5V763Aur9LgD@" + statsPrefix + "stats.pandora.l15.lukso.network",
	}

	fmt.Println(shared.BinaryDir + client + "/" + version + "/" + client + " --datadir " + datadir + " init /opt/lukso/networks/" + network + "/config/pandora-genesis.json &>/dev/null")
	command := exec.Command("bash", "-c", shared.BinaryDir+client+"/"+version+"/"+client+" --datadir "+datadir+" init /opt/lukso/networks/"+network+"/config/pandora-genesis.json &>/dev/null")
	if startError := command.Start(); startError != nil {
		log.Println("ERROR STARTING " + client + "@" + version)
		log.Fatal(startError)
	}

	command.Wait()

	in, err := os.Open("/opt/lukso/networks/" + network + "/config/pandora-nodes.json")
	if err != nil {
		fmt.Println(err)
	}
	defer in.Close()

	out, err := os.Create(datadir + "/geth/static-nodes.json")
	if err != nil {
		fmt.Println(err)
	}
	defer out.Close()

	io.Copy(in, out)
	out.Close()

	StartBinary(client, version, args)
}
