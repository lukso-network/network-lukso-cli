package runner

import (
	"fmt"
	"io"
	"log"
	"lukso/apps/lukso-manager/settings"
	"lukso/apps/lukso-manager/shared"
	"os"
	"os/exec"
	"strings"
)

func startPandora(
	version string,
	network string,
	settings settings.Settings,
	config *NetworkConfig,
	timestamp string,
) (cmd *exec.Cmd, err error) {
	client := "pandora"
	dataDir := shared.GetDataDir(network, client)
	networkDir := shared.GetNetworkDir(network)

	if settings.HostName == "" {
		settings.HostName, _ = os.Hostname()
	}

	hostname := "l15-" + settings.HostName

	statsPrefix := ""
	if !(network == "l15-prod") {
		statsPrefix = strings.Split(network, "-")[1] + "."
	}

	err = os.MkdirAll(dataDir, 0775)
	if err != nil {
		log.Fatal(err)
	}

	err = os.MkdirAll(dataDir+"/geth", 0775)
	if err != nil {
		log.Fatal(err)
	}

	pandoraVerbosity := "4"
	switch settings.Pandora.Verbosity {
	case "silent":
		pandoraVerbosity = "0"
	case "error":
		pandoraVerbosity = "1"
	case "warn":
		pandoraVerbosity = "2"
	case "info":
		pandoraVerbosity = "3"
	case "debug":
		pandoraVerbosity = "4"
	case "detail", "trace":
		pandoraVerbosity = "5"
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
		"--miner.notify=ws://127.0.0.1:7878,http://127.0.0.1:7877",
		"--miner.gaslimit=80000000",
		"--syncmode=full",
		"--verbosity=" + pandoraVerbosity,
		"--nat=extip:" + shared.OutboundIP.String(),
		"--metrics",
		"--metrics.expensive",
		"--pprof",
		"--pprof.addr=127.0.0.1",
		"--ethstats=" + hostname + ":6Tcpc53R5V763Aur9LgD@" + statsPrefix + "stats.pandora.l15.lukso.network",
		// "2> " + networkDir + "/logs/pandora-" + version + "-" + timestamp + ".log",
	}

	if settings.ValidatorEnabled {
		args = append(args, "--mine")
		args = append(args, "--miner.etherbase="+settings.Coinbase)
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
		log.Fatal(errBinary)
		return
	}

	return cmd, nil
}

func stopPandora() error {
	if err := CommandsByClient.pandora.Process.Kill(); err != nil {
		return err
	}
	return nil
}
