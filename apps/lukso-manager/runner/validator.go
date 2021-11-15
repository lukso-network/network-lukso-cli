package runner

import "lukso/shared"

func startValidator(version string, network string) (err error) {
	client := "lukso-validator"

	args := []string{
		"--datadir=" + shared.NetworkDir + network + "/datadirs/validator",
		"--accept-terms-of-use",
		"--beacon-rpc-provider=127.0.0.1:4000",
		"--chain-config-file=" + shared.NetworkDir + network + "/config/vanguard-config.yaml",
		"--verbosity=info",
		"--pandora-http-provider=http://127.0.0.1:8545",
		"--wallet-dir=" + shared.NetworkDir + network + "/vanguard_wallet",
		"--wallet-password-file=" + shared.NetworkDir + network + "/passwords/keys",
		"--rpc",
		"--log-file=" + shared.NetworkDir + network + "/logs/validator",
		"--lukso-network",
	}

	StartBinary(client, version, args)
	return
}
