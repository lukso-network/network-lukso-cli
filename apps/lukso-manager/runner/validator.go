package runner

import (
	"lukso-manager/shared"
	"os/exec"
)

func startValidator(version string, network string) (cmd *exec.Cmd, err error) {
	client := "lukso-validator"
	networkDir := shared.GetNetworkDir(network)

	args := []string{
		"--datadir=" + shared.GetDataDir(network, "validator"),
		"--accept-terms-of-use",
		"--beacon-rpc-provider=127.0.0.1:4000",
		"--chain-config-file=" + networkDir + "/config/vanguard-config.yaml",
		"--verbosity=info",
		"--pandora-http-provider=http://127.0.0.1:8545",
		"--wallet-dir=" + networkDir + "/vanguard_wallet",
		"--wallet-password-file=" + networkDir + "/passwords/keys",
		"--rpc",
		"--log-file=" + networkDir + "/logs/validator",
		"--lukso-network",
	}

	cmd, errBinary := StartBinary(client, version, args)
	if errBinary != nil {
		return
	}

	return cmd, nil
}
