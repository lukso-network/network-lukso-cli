#!/usr/bin/env bash


NETWORK="l16-dev"
PLATFORM="unknown";
NETWORK_VERSION="17"

# for Apple M1s
if [ "$(uname -s)" == "Darwin" ] && [ "$(uname -m)" == "arm64" ]
then
ARCHITECTURE="amd64"
else
ARCHITECTURE=$(uname -m)
ARCHITECTURE=${ARCHITECTURE/x86_64/amd64}
ARCHITECTURE=${ARCHITECTURE/aarch64/arm64}
fi
readonly os_arch_suffix="$(uname -s | tr '[:upper:]' '[:lower:]')-$ARCHITECTURE"

PLATFORM=""
case "$OSTYPE" in
darwin*) PLATFORM="darwin" ;;
linux*) PLATFORM="linux" ;;
msys*) PLATFORM="windows" ;;
cygwin*) PLATFORM="windows" ;;
*) exit 1 ;;
esac
readonly PLATFORM

if [ "$PLATFORM" == "windows" ]; then
    ARCHITECTURE="amd64.exe"
elif [[ "$os_arch_suffix" == *"arm64"* ]]; then
    ARCHITECTURE="arm64"
fi

if [[ "$ARCHITECTURE" == "armv7l" ]]; then
    color "31" "32-bit ARM is not supported. Please install a 64-bit operating system."
    exit 1
fi

download() {
  URL="$1";
  LOCATION="$2";
  if [[ $PLATFORM == "linux" ]]; then
    wget -O $LOCATION $URL;
  fi

  if [[ $PLATFORM == "darwin" ]]; then
    curl -o $LOCATION -Lk $URL;
  fi
}

download_network_config() {
  NETWORK=$1
  NETWORK_NAME="$(cut -d'-' -f1 <<<"$NETWORK")"
  NETWORK_MODE="$(cut -d'-' -f2 <<<"$NETWORK")"

  CDN="https://raw.githubusercontent.com/lukso-network/network-configs/l16-dev/${NETWORK_NAME}/${NETWORK_MODE}/${NETWORK_VERSION}"
  echo $CDN
  mkdir -p ./configs
  TARGET=./configs
  download $CDN/genesis.json?ignoreCache=1 $TARGET/genesis.json
  download $CDN/genesis.ssz?ignoreCache=1 $TARGET/genesis.ssz
  download $CDN/config.yaml?ignoreCache=1 $TARGET/config.yaml
}

update_env_variables() {

  PUB_IP_ADDRESS=$(curl ident.me)

  if [[ $PLATFORM == "linux" ]]; then
    sed -i "s/ETH_STATS_NAME=.*/ETH_STATS_NAME=$(uname -n)/g" .env
    sed -i "s/ETH_2_STATS_NAME=.*/ETH_2_STATS_NAME=$(uname -n)/g" .env
    sed -i "s/PRYSM_HOST_IP=.*/PRYSM_HOST_IP=${PUB_IP_ADDRESS}/g" .env
    sed -i "s/GETH_EXTERNAL_IP=.*/GETH_EXTERNAL_IP=extip:${PUB_IP_ADDRESS}/g" .env

  fi

  if [[ $PLATFORM == "darwin" ]]; then
    sed -i "" "s/ETH_STATS_NAME=.*/ETH_STATS_NAME=$(uname -n)/g" .env
    sed -i "" "s/ETH_2_STATS_NAME=.*/ETH_2_STATS_NAME=$(uname -n)/g" .env
    sed -i "" "s/PRYSM_HOST_IP=.*/PRYSM_HOST_IP=${PUB_IP_ADDRESS}/g" .env
    sed -i "" "s/GETH_EXTERNAL_IP=.*/GETH_EXTERNAL_IP=extip:${PUB_IP_ADDRESS}/g" .env
  fi
}

mkdir -p ./bin

# download latest configs
download_network_config l16-dev;

# download eth2 validator tool and give exec permission
download https://github.com/lukso-network/network-validator-tools/releases/download/v1.0.0/network-validator-tools-v1.0.0-${PLATFORM}-${ARCHITECTURE} ./bin/eth2-val-tools
chmod +x ./bin/eth2-val-tools

# download makefile docker-compose and .env file
download https://raw.githubusercontent.com/lukso-network/network-lukso-cli/feature/lukso-cli-with-kintsugi/shell_scripts/Makefile ./Makefile
download https://raw.githubusercontent.com/lukso-network/network-config-gen/l16-dev/validator-activation/cloud-docker-compose-setup/validator/docker-compose.yml?token=GHSAT0AAAAAABQQS5FPENWIOUTGAMKMR7CQYQ6AEGA ./docker-compose.yml;
download https://raw.githubusercontent.com/lukso-network/network-config-gen/l16-dev/validator-activation/cloud-docker-compose-setup/validator/.env?token=GHSAT0AAAAAABQQS5FPUTOBR4HUFBFXA5EGYQ6AE7A ./.env

update_env_variables

echo ""
echo "#################### Please Read Carefully ####################"
echo "Update \"GETH_ETHERBASE\" with your address \".env\" file"
echo "use \"make help\" to check available options"
echo "###############################################################"
