#!/usr/bin/env bash


NETWORK="l15"
PLATFORM="unknown";
ARCHITECTURE=$(uname -m);

ORCHESTRATOR_TAG="";
PANDORA_TAG="";
VANGUARD_TAG="";

if [[ "$OSTYPE" = "linux-gnu" ]]; then
  PLATFORM="Linux";
elif [[ "$OSTYPE" = "darwin"* ]]; then
  PLATFORM="Darwin"
elif [[ "$OSTYPE" = "cygwin" ]]; then
  PLATFORM="Linux"
elif [[ "$OSTYPE" = "freebsd" ]]; then
  PLATFORM="Linux"
fi

if [[ "$PLATFORM" = "unknown" ]]; then
  echo "Platform not supported.";
  exit;
fi

if [[ $PLATFORM == "Linux" ]]; then
  sudo apt-get update;
  sudo apt-get install curl \
  wget \
  unzip -y;
fi

download() {
  URL="$1";
  LOCATION="$2";
  if [[ $PLATFORM == "Linux" ]]; then
    sudo wget -O $LOCATION $URL;
  fi

  if [[ $PLATFORM == "Darwin" ]]; then
    sudo curl -o $LOCATION -Lk $URL;
  fi
}

download_network_config() {
  NETWORK=$1;
  CDN="https://storage.googleapis.com/l15-cdn/networks/$NETWORK";
  sudo mkdir -p /opt/lukso/networks/$NETWORK/config;
  TARGET=/opt/lukso/networks/$NETWORK/config;
  download $CDN/network-config.yaml $TARGET/network-config.yaml;
  download $CDN/pandora-genesis.json $TARGET/pandora-genesis.json;
  download $CDN/vanguard-genesis.ssz $TARGET/vanguard-genesis.ssz;
  download $CDN/vanguard-config.yaml $TARGET/vanguard-config.yaml;
  download $CDN/pandora-nodes.json $TARGET/pandora-nodes.json;
}

sudo mkdir \
/opt/lukso \
/opt/lukso/tmp \
/opt/lukso/binaries \
/opt/lukso/networks \
/opt/lukso/networks/"$NETWORK" \
/opt/lukso/networks/"$NETWORK"/config;

download https://raw.githubusercontent.com/lukso-network/lukso-orchestrator/feature/l15-setup/scripts/lukso /opt/lukso/lukso;

sudo chmod +x /opt/lukso/lukso;
sudo ln -sfn /opt/lukso/lukso /usr/local/bin/lukso;

download_network_config l15;

sudo rm -rf /opt/lukso/tmp;

sudo lukso bind-binaries \
--orchestrator v0.1.0-beta.1 \
--pandora v0.1.0-beta.1 \
--vanguard v0.1.0-beta.1 \
--validator v0.1.0-beta.1 \
--eth2stats 1;

echo "Ready! type lukso to start the node!";
