# LUKSO CLI
>⚠️ This page may change. Not everything is ready yet.


## Repository struct 
In `./shell_scripts` are currently used scripts that will be replaced by proper binary.

* `install-unix.sh` installer for Linux/Darwin
* `lukso` script for Linux/Darwin
* `install-win.ps1` installer for Windows
* `lukso-win.ps1` script for Windows

## System requirements

### Minimum specifications

These specifications must be met in order to successfully run the Vanguard, Pandora, and Orchestrator clients.

- Operating System: 64-bit Linux, Mac OS X 10.14+
- Processor: Intel Core i5–760 or AMD FX-8100 or better
- Memory: 8GB RAM
- Storage: 20GB available space SSD
- Internet: Broadband connection

### Recommended specifications

These hardware specifications are recommended, but not required to run the Vanguard, Pandora, and Orchestrator clients.

- Processor: Intel Core i7–4770 or AMD FX-8310 or better
- Memory: 16GB RAM
- Storage: 100GB available space SSD
- Internet: Broadband connection

>⚠️ Currently we do not support Apples new M1 chips and Windows yet. 

## Installation ( Linux/MacOS )

`curl https://install.l15.lukso.network | bash`

This shell script will:
1. Create directory under `/opt/lukso`
2. Download binary executables and config files required for node startup.
3. Place them in `/opt/lukso`
4. Create symbolic link in `/usr/local/bin`.


## Installation ( Windows )
>🛠️ Work In Progress, available soon.  
> Requires [PowerShell](https://docs.microsoft.com/en-us/powershell/scripting/install/installing-powershell?view=powershell-7.1)

## Running
Enter `lukso start` to start an archive node  
Enter `lukso start --validate --coinbase <ETH1_Address>` to start a validator node (Read instructions on validating first) 

You may need to use `sudo` on `macos` devices.

## Config file
Enter `lukso config` in your shell to generate config file.

Example:
~~~yaml
COINBASE: "0x616e6f6e796d6f75730000000000000000000000"
WALLET_DIR: "/home/user/.lukso/l16-prod/beacon-wallet"
DATADIR: "/home/user/.lukso/l16-prod/datadirs"
LOGSDIR: "/home/user/.lukso/l16-prod/logs"
NODE_NAME: "l15-johnsmith123"
~~~
After that, you can use `--config /path/to/config.yaml` insted of other flags:  



## Available parameters
`lukso <command> [argument] [--flags]`

| command   | description            | argument |
|-----------|------------------------|----------------------|
| start     | Starts up all or specific client(s) | [geth, beacon, validator, eth2stats-client, **all**] |
| stop      | Stops all or specific client(s)     | [geth, beacon, validator, eth2stats-client, **all**] |
| reset     | Clears client(s) datadirs (this also removes chain-data) | [geth, beacon, validator, all, **none**]
| config    | Interactive tool for creating config file | |
| keygen    | Runs `network-validator-tools` to generate keystore and wallet | |
| gen-deposit-data    | Uses `network-validator-tools` to `deposit-keys.json`| |
| logs      | Show logs | [orchestrator, geth, beacon, validator, eth2stats-client, lukso-status] |
| bind-binaries      | sets client(s) to desired version | 
| version      | Shows the LUKSO script version | 
> In **bold** is a behaviour when argument is skipped (default)

### start

| name      | description            | Argument  |
|-----------|------------------------|---|
| --network | Picks which setup to use | Name of network from list: `mainnet, l16-prod, l16-staging, l16-dev`
| --l16-prod | Shorthand alias for `--network l16-prod` | <none\>
| --l16-staging | Shorthand alias for `--network l16-staging` | <none\>
| --l16-dev | Shorthand alias for `--network l16-dev` | <none\>
| --config | Path to config file     | Path ex. `config.yaml` |
| --validate | Starts validator      | <none\>
| --coinbase | Sets geth coinbase. This is public address for block mining rewards (default = first account created) (default: "0") | ETH1 addres ex. `0x144a9533B3d759d647597762d33a1cD6f9Bf118c`
| --node-name  | Name of node that's shown on geth stats and beacon stats | String ex. `johnsmith123` 
| --logsdir  | Sets the logs path | String ex. `/mnt/external/lukso-logs` 
| --datadir  | Sets datadir path | String ex. `/mnt/external/lukso-datadir`
| --home  | Sets path for datadir and logs in a single location (--datadir and --logs take priority) | String ex. `/var/lukso` 
| --network-version  | Picup the network version for which configs will be downloaded 
| --geth-verbosity  | Sets geth logging depth (note: geth uses integers for that flag, script will convert those to proper values) | String ex. `silent, error, warn, info, debug, trace` 
| --geth-bootnodes  | Sets geth bootnodes | Strings of bootnodes separated by commas: `enode://72caa...,enode://b4a11a...`
| --geth-http-port  | Sets geth RPC (over http) port | Number between 1023-65535
| --geth-metrics  | Enables geth metrics server | <none\>
| --geth-nodekey  | P2P node key file | Path to file (relative or absolute)
| --geth-rpcvhosts  | Sets geth rpc virtual hosts (use quotes if you want to set \* `'*'` otherwise shell will resolve it) | Comma-separated list of virtual hosts Ex. `localhost` or `*`
| --geth-external-ip  | Sets external IP for geth (overrides --external-ip if present) | String ex. `72.122.32.234`
| --geth-universal-profile-expose  | Exposes "net,eth,txpool,web3" API's on geth RPC | <none\>
| --geth-unsafe-expose  | Exposes ALL API's ("admin,net,eth,debug,miner,personal,txpool,web3") API's on geth RPC | <none\>
| --beacon-verbosity  | Sets beacon-client logging depth | String ex. `silent, error, warn, info, debug, trace`
| --beacon-bootnodes  | Sets beacon-client bootnodes | Strings of bootnodes separated by commas: `enr:-Ku4QAmY...,enr:-M23QLmY...`
| --beacon-p2p-priv-key  | The file containing the private key to use in communications with other peers. | Path to file (relative or absolute)
| --beacon-external-ip  | Sets external IP for beacon-client (overrides --external-ip if present) | IP ex. `72.122.32.234`
| --beacon-p2p-host-dns  | Sets host DNS beacon-client (overrides --external-ip AND --beacon-external-ip if present) | DNS name ex. `l16-nodes-1.nodes.l16.lukso.network`
| --beacon-rpc-host  | Sets beacon-client RPC listening interface | IP ex. `127.0.0.1`
| --beacon-monitoring-host  | Sets beacon-client monitoring listening interface | IP ex. `127.0.0.1`
| --validator-verbosity  | Sets validator logging depth | String ex. `silent, error, warn, info, debug, trace`
| --wallet-dir  | Sets directory of `lukso-validator` wallet  | Path to directory, relative or absolute
| --wallet-password-file  | Sets location of password file for validator (without it, it will always prompt for password)  | Path to a file, relative or absolute
| --cors-domain  | Sets CORS domain (note: if you want to set every origin you must type asterisk wrapped in quotes `'*'` otherwise shell may try to resolve it | CORS Domain ex. `localhost`, `*`
| --external-ip  | Sets external IP for geth and beacon-chain | String ex. `72.122.32.234`
| --allow-respin  | Deletes all datadirs IF network config changed (based on genesis time) | <none\>
| --beacon-http-web3provider  | An eth1 web3 provider string http endpoint or IPC socket path. (default: http://127.0.0.1:8545) | URL address, e.g. `http://127.0.0.1:8545`
| --beacon-rpc-host  | Host on which the RPC server should listen. (default: 127.0.0.1) | IPv4 address, e.g. `127.0.0.1`
| --beacon-rpc-port  | Port on which the RPC server should listen. (default: 4000) | Port, e.g. `4000`
| --beacon-udp-port  | beacon chain client UDP port. The port used by discv5. (default: 12000) | Port number, e.g. `12000`
| --beacon-tcp-port  | beacon chain client TCP port. The port used by libp2p. (default: 13000) | Port number, e.g. `13000`
| --beacon-grpc-gateway-port   | Vanguard gRPC gateway port. The port on which the gateway server runs on (default: 3500) | Gateway port, e.g. `3500`
| --validator-beacon-rpc-provider | Beacon node RPC provider endpoint. (default is: 127.0.0.1:4000) | IPv4 with port, e.g. `127.0.0.1:4000`
| --validator-geth-http-provider | A geth rpc endpoint. This is our geth client http endpoint. (default is: http://127.0.0.1:8545) | URL or IPC socket path, e.g. `http://127.0.0.1:8545`
| --eth2stats-beacon-addr | Beacon node endpoint address for eth2stats-client. (default: 127.0.0.1:4000) | IPv4 with port, e.g. `127.0.0.1:4000`
| --geth-port | Geth client TCP/UDP port exposed. Default is: 30405 | Port number, e.g. `30405`
| --geth-http-addr | Geth client http address exposed. Default is: 127.0.0.1 | IPv4 address, e.g. `127.0.0.1`
| --geth-http-port | Geth client http port exposed. Default is: 8545 | Port number, e.g. `8545`
| --geth-ws-addr  | Geth client websocket address exposed. Default is: 127.0.0.1 | IPv4 address, e.g. `127.0.0.1`
| --geth-ws-port | Geth client websocket port exposed. Default is: 8546 | Port number, e.g. `8546`
| --geth-http-miner-addr | Geth HTTP URL to notify of new work packages. Default is: http://127.0.0.1:7877 | HTTP address, e.g. `http://127.0.0.1:7877`
| --geth-ws-miner-addr | Geth Websocket URL to notify of new work packages. Default is: ws://127.0.0.1:7878 | WS address, e.g. `ws://127.0.0.1:7878`
| --geth-ethstats | Geth flag to activate ethstats listing on remote dashboard. If enabled you should see your node by your node name provided via --node-name flag or lukso config. (default:  disabled) | Token and address like `token123@stats.example.com`
| --beacon-ethstats | Beacon-chain flag fo activate eth2stats listing on remote dashboard. If enabled you should see your node by your node name provided via --node-name flag or lukso config. (default:  disabled) | Address with port, e.g. `192.168.0.1:9090`
| --beacon-min-sync-peers | The required number of valid beacon-chain peers to connect with before syncing. (default: 2) | Number of required peers, e.g. `1`
| --beacon-max-p2p-peers | The max number of beacon-chain p2p peers to maintain. (default: 50) | Peers count, e.g. `70`
| --beacon-ethstats-metrics | The metrics address for beacon-chain eth2stats-client service (default: http://127.0.0.1:8080/metrics) | HTTP address with port and `metrics` endpoint, e.g. `http://127.0.0.1:8080/metrics`
| --status-page | This flag is for lukso-status activation. With this service you can check your node status over web browser (default: disabled). Default web address is: http://127.0.0.1:8111 | <none\>

How to use flags with values? Provide a flag and value like: `lukso start --datadir /data/network-node`

### stop
| name      | description            | Argument  |
|-----------|------------------------|---|
| --force   | Adds force option to kill commands (may result in corruption of node data)     | <none\> |

### gen-deposit-data
| name      | description            | Argument  |
|-----------|------------------------|---|
| --keys-dir  | Sets directory of `lukso-deposit-cli` keys | Path to directory, relative or absolute

### keygen
| name      | description            | Argument  |
|-----------|------------------------|---|
| --wallet-dir  | Sets directory of `lukso-validator` wallet  | Path to directory, relative or absolute


### bind-binaries 
| name      | description            | Argument  |
|-----------|------------------------|---|
| --geth   | download and set `geth` to given tag  | Tag, ex. `v0.1.0-rc.1` |
| --beacon   | download and set `beacon-chain` to given tag  | Tag, ex. `v0.1.0-rc.1` |
| --validator   | download and set `validator` to given tag  | Tag, ex. `v0.1.0-rc.1` |
| --deposit   | download and set `lukso-deposit-cli` to given tag  | Tag, ex. `v0.1.0-rc.1` |
| --eth2stats   | download and set `eth2stats` to given tag  | Tag, ex. `v0.1.0-rc.1` |

