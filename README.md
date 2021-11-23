# Node-GUI

Very basic Node GUI. Runs on port `8111`.

## Install dependencies & setup
Prepared directories and settings.db file:

```bash
mkdir -p ~/.lukso && touch ~/.lukso/settings.db 
```
install npm packages using yarn:

```bash
yarn install 
```

Install gow - the [Go Watcher](https://github.com/mitranim/gow):

```bash
go install github.com/mitranim/gow@latest
```
Don't forget to export paths for Go in order to have gow work: 

```bash
export GOPATH=~/go
export PATH=$PATH:$GOPATH/bin
```

## How to run

To start the web application inside the root of the project run:

```bash
yarn start
```
Open [http://localhost:4200](http://localhost:4200)

For running the Go proxy service between webapp and the node run the following:

```bash
cd apps/lukso-manager && gow run .
```

## Unit and E2E

Coming soon.

## How to install

Coming soon.
