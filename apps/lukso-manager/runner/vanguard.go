package runner

func startVanguard(version string, network string) {
	client := "vanguard"
	args := []string{
		"--accept-terms-of-use",
		"--chain-id=231",
		"--network-id=231",
		"--datadir=/home/rryter/.lukso/networks/" + network + "/datadirs/vanguard",
		"--genesis-state=/opt/lukso/networks/l15-dev/config/vanguard-genesis.ssz",
		"--chain-config-file=/opt/lukso/networks/l15-dev/config/vanguard-config.yaml",
		"--bootstrap-node=enr:-LK4QEuvsLAqYHxOywH2z90bQATmBhW4jATtTDQM3UmwvaPPZrB_Hw27RfdCpWL5LGjeJMpQzGH733AggZJHICWqKUIBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpCGlkzsY6VTF___________gmlkgnY0gmlwhCJaZhuJc2VjcDI1NmsxoQPLl7XpO1pdZh7KjzgnmdlrwHXwpv4kQ5Jcs3pr4lIe94N0Y3CCMsiDdWRwgi7g",
		"--bootstrap-node=enr:-LK4QM9F8hglsGH8IDWebcZw23gH_UtD3vRUoG6z-58veXxkG1nBpiQDiYUvCN2ITFFph0_e-J_76NnjaCRCxw1psMcBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpCGlkzsY6VTF___________gmlkgnY0gmlwhCPMkCWJc2VjcDI1NmsxoQONZbpRUPmtDBCSySZXl_6wnlCc6hltHHI9v--swVK61IN0Y3CCMsiDdWRwgi7g",
		"--bootstrap-node=enr:-LK4QCAEYjF2-gOzDfn_X9-Ns_M-EyapmVYZCRpyyW2_PuahHTJsvDX59Jr9UJWn3feVfc1bN97Lf9Uj5MzxseRjpQwBh2F0dG5ldHOIAAAAAAAAAACEZXRoMpCGlkzsY6VTF___________gmlkgnY0gmlwhCJabGOJc2VjcDI1NmsxoQL80SkifzxBiR278Eol3_7T-34JEW6swCXF0TF53VDz9YN0Y3CCMsiDdWRwgi7g",
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
		"--p2p-host-ip=46.127.26.82",
	}

	StartBinary(client, version, args)
}
