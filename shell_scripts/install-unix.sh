#!/usr/bin/env bash


NETWORK="l15-dev"
PLATFORM="unknown";
NETWORK_VERSION="3"

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

if [[ $PLATFORM == "linux" ]]; then
  sudo apt-get update;
  sudo apt-get install curl \
  wget \
  unzip -y;
fi

if [[ ! -d "/usr/local/bin" ]]; then
  sudo mkdir -p /usr/local/bin;
fi

download() {
  URL="$1";
  LOCATION="$2";
  if [[ $PLATFORM == "linux" ]]; then
    sudo wget -O $LOCATION $URL;
  fi

  if [[ $PLATFORM == "darwin" ]]; then
    sudo curl -o $LOCATION -Lk $URL;
  fi
}

download_network_config() {
  NETWORK=$1
  NETWORK_NAME="$(cut -d'-' -f1 <<<"$NETWORK")"
  NETWORK_MODE="$(cut -d'-' -f2 <<<"$NETWORK")"

  CDN="https://raw.githubusercontent.com/lukso-network/network-configs/l16-dev/${NETWORK_NAME}/${NETWORK_MODE}/${NETWORK_VERSION}"
  sudo mkdir -p /opt/lukso/networks/$NETWORK/config
  TARGET=/opt/lukso/networks/$NETWORK/config
  download $CDN/genesis.json?ignoreCache=1 $TARGET/geth-genesis.json
  download $CDN/genesis.ssz?ignoreCache=1 $TARGET/beacon-genesis.ssz
  download $CDN/config.yaml?ignoreCache=1 $TARGET/beacon-config.yaml
  # TODO: CHANGE THIS WHEN DEPLOYING REAL NETWORK
  # download $CDN/pandora-nodes.json?ignoreCache=1 $TARGET/geth-nodes.json
}

sudo mkdir \
/opt/lukso \
/opt/lukso/tmp \
/opt/lukso/binaries \
/opt/lukso/networks ;

# TODO: CHANGE THIS LOCATION LATER. IT IS FOR TEST PURPOSE ONLY.
download https://raw.githubusercontent.com/lukso-network/network-lukso-cli/feature/lukso-cli-with-kintsugi/shell_scripts/lukso /opt/lukso/lukso;

sudo chmod +x /opt/lukso/lukso;
sudo ln -sfn /opt/lukso/lukso /usr/local/bin/lukso;

download_network_config l15-prod;
download_network_config l15-staging;
download_network_config l15-dev;

sudo rm -rf /opt/lukso/tmp;

sudo lukso bind-binaries \
--geth v1.0.0 \
--beacon v1.0.0 \
--validator v1.0.0 \
--deposit v1.0.0 \
--eth2stats v1.0.0;

echo "Ready! type lukso to start the node!";
