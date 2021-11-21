package runner

import (
	"fmt"
	"io"
	"log"
	"lukso-manager/settings"
	"lukso-manager/shared"
	"os"
	"os/exec"
	"strings"
)

func startPandora(version string, network string, settings settings.Settings) (cmd *exec.Cmd, err error) {
	client := "pandora"
	dataDir := shared.GetDataDir(network, client)
	networkDir := shared.GetNetworkDir(network)
	if settings.HostName == "" {
		settings.HostName, _ = os.Hostname()
	}

	hostname := "l15-" + settings.HostName

	config, err := ReadConfig(network)
	if err != nil {
		return
	}

	statsPrefix := ""
	if !(network == "l15-prod") {
		statsPrefix = strings.Split(network, "-")[1] + "."
	}

	args := []string{
		"--datadir=" + dataDir,
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
		"--miner.etherbase=" + settings.Coinbase,
		"--miner.gaslimit=80000000",
		"--syncmode=full",
		"--verbosity=4",
		"--nat=extip:" + shared.OutboundIP.String(),
		"--metrics",
		"--metrics.expensive",
		"--pprof",
		"--pprof.addr=127.0.0.1",
		"--ethstats=" + hostname + ":6Tcpc53R5V763Aur9LgD@" + statsPrefix + "stats.pandora.l15.lukso.network",
	}

	command := exec.Command("bash", "-c", shared.BinaryDir+client+"/"+version+"/"+client+" --datadir "+dataDir+" init "+networkDir+"/config/pandora-genesis.json &>/dev/null")
	if startError := command.Start(); startError != nil {
		log.Fatal(startError)
		return
	}

	command.Wait()

	in, err := os.Open(networkDir + "/config/pandora-nodes.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer in.Close()

	out, err := os.Create(dataDir + "/geth/static-nodes.json")
	if err != nil {
		fmt.Println(err)
		return
	}
	defer out.Close()

	io.Copy(in, out)
	out.Close()

	cmd, errBinary := StartBinary(client, version, args)
	if errBinary != nil {
		return
	}

	return cmd, nil
}
