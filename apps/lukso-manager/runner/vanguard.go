package runner

import (
	"fmt"
	"lukso/apps/lukso-manager/shared"
	"os/exec"
	"strings"
)

func startVanguard(version string, network string, config *NetworkConfig, timestamp string) (cmd *exec.Cmd, err error) {
	client := "vanguard"
	bootnodes := strings.Split(config.VANGUARDBOOTNODES, ",")

	args := []string{
		"--accept-terms-of-use",
		"--chain-id=" + fmt.Sprint(config.CHAINID),
		"--network-id=" + fmt.Sprint(config.NETWORKID),
		"--datadir=" + shared.GetDataDir(network, client),
		"--genesis-state=" + shared.NetworkDir + network + "/config/vanguard-genesis.ssz",
		"--chain-config-file=" + shared.NetworkDir + network + "/config/vanguard-config.yaml",
		"--bootstrap-node=" + bootnodes[0],
		"--bootstrap-node=" + bootnodes[1],
		"--bootstrap-node=" + bootnodes[2],
		"--http-web3provider=http://127.0.0.1:8545",
		"--deposit-contract=0x000000000000000000000000000000000000cafe",
		"--contract-deployment-block=0",
		"--rpc-host=127.0.0.1",
		"--verbosity=debug",
		"--min-sync-peers=1",
		"--p2p-max-peers=50",
		"--orc-http-provider=http://127.0.0.1:7877",
		"--rpc-port=4000",
		"--p2p-udp-port=12000",
		"--p2p-tcp-port=13000",
		"--grpc-gateway-port=3500",
		"--update-head-timely",
		"--lukso-network",
		"--p2p-host-ip=" + shared.OutboundIP.String(),
		"--log-file=" + shared.NetworkDir + network + "/logs/" + fmt.Sprint(config.GENESISTIME) + "/" + client + "-" + version + "-" + timestamp + ".log",
	}

	cmd, errBinary := StartBinary(client, version, args)
	if errBinary != nil {
		return
	}

	return cmd, nil
}
