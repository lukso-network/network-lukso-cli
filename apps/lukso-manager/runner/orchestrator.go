package runner

import (
	"lukso-manager/shared"
	"os/exec"
)

func startOrchestrator(version string, network string) (cmd *exec.Cmd, err error) {
	client := "lukso-orchestrator"
	args := []string{
		"--datadir=" + shared.NetworkDir + network + "/" + shared.DataDir + "/orchestrator",
		"--vanguard-grpc-endpoint=127.0.0.1:4000",
		"--http",
		"--http.addr=127.0.0.1",
		"--http.port=7877",
		"--ws",
		"--ws.addr=127.0.0.1",
		"--ws.port=7878",
		"--pandora-rpc-endpoint=ws://127.0.0.1:8546",
		"--verbosity=debug",
	}

	cmd, errBinary := StartBinary(client, version, args)
	if errBinary != nil {
		return
	}

	return cmd, nil
}
