package runner

func startPandora(version string, network string) {
	client := "pandora"
	args := []string{
		"--datadir=/home/rryter/.lukso/networks/" + network + "/datadirs/pandora",
		"--networkid=231",
		"--port=30405",
		"--http",
		"--http.addr=127.0.0.1",
		"--http.port=8545",
		"--bootnodes=enode://f9c894b27de56e0a96d849e5a9878cf1baa8c9fbcb7f7fd60e2694398afbb00569e530242470d46bf53749ca2329bc13fddf99b63236ef7463ea9a369a29acfa@34.90.108.99:30405,enode://bd425e22328db3c6979342c3001bb0fbec9559718e079552a9846126a6c313a34e9da6857e2b1f1d0219fbd08b43daa658b9ebcd7b51cff16d97e41cd4c3dd2a@35.204.144.37:30405,enode://3390c97e8252fff4a9262a9f11092b10b985319c7e5e8f1549caabc2f08de0fe87c7f74f6daf17f6867011d190413a47c27b879e2093df306516227db2a874f5@34.90.102.27:30405",
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
		"--ethstats=l15-rryter-gui:6Tcpc53R5V763Aur9LgD@dev.stats.pandora.l15.lukso.network",
	}

	StartBinary(client, version, args)
}
