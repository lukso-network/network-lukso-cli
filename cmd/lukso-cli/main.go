package main

import (
	"fmt"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"google.golang.org/grpc"
	"os"
	runtimeDebug "runtime/debug"
	"sync"
	"time"
)

// We want to spin also 3 libraries at once, and secretly rule them by cli. It matches for me somehow

// This binary will also support only some of the possible networks.
// Make a pull request to attach your network.
// We are also very open to any improvements. Please make some issue or hackmd proposal to make it better.

const (
	ubuntu  = "linux"
	macos   = "darwin"
	windows = "windows"

	DefaultELNetworkID = 1337222
	DefaultELHTTPPort  = 8598
	DefaultELWSPort    = 8599
	DefaultELP2PPort   = 30398

	DefaultCLGRPCPort        = 4098
	DefaultCLGRPCGatewayPort = 3598
	DefaultCLP2PTCPPort      = 13098
	DefaultCLP2PUDPPort      = 13098

	PidFilename     = "pid"
	ArchNodeCmdName = "arch"
)

var (
	appName                   = "lukso-cli"
	ELTag                     string
	validatorTag              string
	CLTag                     string
	CLStatsClientTag          string
	log                       = logrus.WithField("prefix", appName)
	systemOs                  string
	systemArch                string
	ELRuntimeFlags            []string
	validatorRuntimeFlags     []string
	CLRuntimeFlags            []string
	CLStatsClientRuntimeFlags []string
	archiveNode               bool
)

func init() {
	allFlags := make([]cli.Flag, 0)
	allFlags = append(allFlags, ELFlags...)
	allFlags = append(allFlags, validatorFlags...)
	allFlags = append(allFlags, CLFlags...)
	allFlags = append(allFlags, CLStatsClientFlags...)
	allFlags = append(allFlags, appFlags...)
	appFlags = allFlags
}

func main() {
	setupOperatingSystem()

	beforeFunc := func(ctx *cli.Context) error {
		formatter := new(prefixed.TextFormatter)
		formatter.TimestampFormat = "2006-01-02 15:04:05"
		formatter.FullTimestamp = true
		// If persistent log files are written - we disable the log messages coloring because
		// the colors are ANSI codes and seen as gibberish in the log files.
		formatter.DisableColors = ctx.String(LogFileName.Name) != ""
		logrus.SetFormatter(formatter)

		// EL related parsing
		ELTag = ctx.String(ELTagFlag)
		ELRuntimeFlags = prepareELFlags(ctx)

		// Validator related parsing
		validatorTag = ctx.String(validatorTagFlag)
		validatorRuntimeFlags = prepareValidatorFlags(ctx)

		// CL Stats Client related parsing
		CLStatsClientTag = ctx.String(CLStatsClientFlag)
		CLStatsClientRuntimeFlags = prepareCLStatsClientFlags(ctx)

		// CL related parsing
		CLTag = ctx.String(CLTagFlag)
		CLRuntimeFlags = prepareCLFlags(ctx)

		return nil
	}

	archBeforeFunc := func(ctx *cli.Context) (err error) {
		if ctx.Command.Name == ArchNodeCmdName {
			archiveNode = true
		}

		return
	}

	app := &cli.App{
		Name: "lukso-cli",
		Commands: []*cli.Command{
			{
				Name:  "start",
				Usage: "starts ecosystem components",
				Subcommands: []*cli.Command{
					{
						Name:   "all",
						Usage:  "starts all components",
						Before: archBeforeFunc,
						Action: downloadAndRunBinaries,
					},
					{
						Name:   ArchNodeCmdName,
						Usage:  "starts bootnode setup only (archive node)",
						Before: archBeforeFunc,
						Action: downloadAndRunBinaries,
					},
				},
			},
			{
				Name:   "stop",
				Usage:  "Stops all merge ecosystem components",
				Before: archBeforeFunc,
				Action: stopAllBinaries,
			},
		},
		Flags:  appFlags,
		Before: beforeFunc,
	}

	defer func() {
		if x := recover(); x != nil {
			log.Errorf("Runtime panic: %v\n%v", x, string(runtimeDebug.Stack()))
			panic(x)
		}
	}()

	err := app.Run(os.Args)
	if nil != err {
		log.Fatal(err)
	}
}

func downloadAndRunBinaries(ctx *cli.Context) (err error) {
	// Get os, then download all binaries into datadir matching desired system
	// After successful download run binary with desired arguments spin and connect them
	// Orchestrator can be run from-memory
	err = downloadGenesis(ctx)
	if nil != err {
		return
	}

	err = downloadEL(ctx)
	if nil != err {
		return
	}

	err = downloadConfig(ctx)
	if nil != err {
		return
	}

	err = downloadCL(ctx)
	if nil != err {
		return
	}

	err = startEL(ctx)
	if nil != err {
		return
	}

	time.Sleep(time.Second * 5)

	err = startCL(ctx)
	if nil != err {
		return
	}

	time.Sleep(time.Second * 5)

	if !archiveNode {
		err = downloadValidator(ctx)
		if nil != err {
			return
		}

		err = startValidator(ctx)
		if err != nil {
			return
		}
	}

	err = startCLStatsClient(ctx)

	return
}

func stopAllBinaries(ctx *cli.Context) (err error) {
	err = stopCLStatsClient(ctx)
	if err != nil {
		return
	}

	if !archiveNode {
		err = stopValidator(ctx)
		if err != nil {
			return
		}
	}

	err = stopCL(ctx)
	if err != nil {
		return
	}

	err = stopEL(ctx)
	if err != nil {
		return
	}

	return
}

func downloadEL(ctx *cli.Context) (err error) {
	log.WithField("dependencyTag", ELTag).Info("Downloading Execution Layer client")
	elDataDir := ctx.String(ELDatadirFlag)
	err = clientDependencies[ELDependencyName].Download(ELTag, elDataDir)

	return
}

func downloadGenesis(ctx *cli.Context) (err error) {
	log.WithField("dependencyTag", ELTag).Info("Downloading Execution Layer genesis")
	elDataDir := ctx.String(ELDatadirFlag)
	err = clientDependencies[ELGenesisDependencyName].Download(ELTag, elDataDir)

	if nil != err {
		return
	}

	log.WithField("dependencyTag", CLTag).Info("Downloading Consensus Layer genesis")
	CLDataDir := ctx.String(CLDatadirFlag)
	err = clientDependencies[CLGenesisDependencyName].Download(CLTag, CLDataDir)

	return
}

func downloadConfig(ctx *cli.Context) (err error) {
	log.WithField("dependencyTag", CLTag).Info("Downloading Consensus Layer config")
	CLDataDir := ctx.String(CLDatadirFlag)
	err = clientDependencies[CLConfigDependencyName].Download(CLTag, CLDataDir)

	return
}

func downloadCL(ctx *cli.Context) (err error) {
	log.WithField("dependencyTag", CLTag).Info("Downloading Consensus Layer")
	CLDataDir := ctx.String(CLDatadirFlag)
	err = clientDependencies[CLDependencyName].Download(CLTag, CLDataDir)

	return
}

func downloadValidator(ctx *cli.Context) (err error) {
	log.WithField("dependencyTag", validatorTag).Info("Downloading validator")
	validatorDataDir := ctx.String(validatorDatadirFlag)
	err = clientDependencies[validatorDependencyName].Download(validatorTag, validatorDataDir)

	return
}

func downloadCLStatsClient(ctx *cli.Context) (err error) {
	log.WithField("dependencyTag", CLStatsClientTag).Info("Downloading CL Stats Client")
	dataDir := ctx.String(clStatsClientDatadirFlag)
	err = clientDependencies[CLStatsClientDependencyName].Download(CLStatsClientTag, dataDir)

	return
}

func startEL(ctx *cli.Context) (err error) {
	log.WithField("dependencyTag", ELTag).Info("Running genesis.json init")
	elDataDir := ctx.String(ELDatadirFlag)
	elGenesisArguments := []string{
		"init",
		clientDependencies[ELGenesisDependencyName].ResolveBinaryPath(ELTag, elDataDir),
		"--datadir",
		elDataDir,
	}

	pid, err := clientDependencies[ELDependencyName].Run(
		ELTag,
		elDataDir,
		elGenesisArguments,
		ctx.Bool(ELOutputFlag),
	)
	if nil != err {
		return
	}

	log.WithField("dependencyTag", ELTag).WithField("pid", pid).Info("Running execution engine init command")

	time.Sleep(time.Second * 3)

	log.WithField("dependencyTag", ELTag).Info("Starting execution engine")
	pid, err = clientDependencies[ELDependencyName].Run(
		ELTag,
		elDataDir,
		ELRuntimeFlags,
		ctx.Bool(ELOutputFlag),
	)
	if nil != err {
		return
	}

	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(1)

	go func() {
		for {
			ipcEndpoint := fmt.Sprintf("%s/geth.ipc", elDataDir)
			_, currentErr := os.Stat(ipcEndpoint)
			if nil == currentErr {
				log.WithField("pid", pid).Info("Execution Layer up")
				waitGroup.Done()

				return
			}

			if os.IsNotExist(currentErr) {
				time.Sleep(time.Millisecond * 50)
				log.Infof("Execution Layer dead, %s", ipcEndpoint)

				continue
			}

			panic(err)
		}
	}()

	waitGroup.Wait()

	return
}

func stopEL(ctx *cli.Context) (err error) {
	log.Info("Stopping Execution")
	dataDir := ctx.String(ELDatadirFlag)

	err = clientDependencies[ELDependencyName].Stop(
		dataDir,
	)
	if nil != err {
		return
	}

	return
}

func startCL(ctx *cli.Context) (err error) {
	log.WithField("dependencyTag", CLTag).Info("Starting Consensus Layer")
	CLDataDir := ctx.String(CLDatadirFlag)
	pid, err := clientDependencies[CLDependencyName].Run(
		CLTag,
		CLDataDir,
		CLRuntimeFlags,
		ctx.Bool(CLOutputFlag),
	)
	if nil != err {
		return
	}

	log.Info("Consensus Layer readiness check")
	time.Sleep(time.Millisecond * 250)
	dialOptions := constructDialOptions(0, 100, time.Second)
	vanClient, err := grpc.DialContext(
		ctx.Context,
		CLGrpcEndpoint,
		dialOptions...,
	)

	if nil != err || nil == vanClient {
		log.WithField("cause", "Consensus Layer not ready yet").Error(err)

		return fmt.Errorf("consensus layer not ready yet: %s", err.Error())
	}

	log.WithField("pid", pid).Info("Consensus Layer is ready")

	return vanClient.Close()
}

func stopCL(ctx *cli.Context) (err error) {
	log.Info("Stopping Consensus")
	dataDir := ctx.String(CLDatadirFlag)

	err = clientDependencies[CLDependencyName].Stop(
		dataDir,
	)
	if nil != err {
		return
	}

	return
}

func startValidator(ctx *cli.Context) (err error) {
	// First command should be to create wallet or prompt to do this by your own. This is one-time
	log.WithField("dependencyTag", validatorTag).Info("Starting Consensus Layer Validator")
	validatorDataDir := ctx.String(validatorDatadirFlag)
	pid, err := clientDependencies[validatorDependencyName].Run(
		validatorTag,
		validatorDataDir,
		validatorRuntimeFlags,
		ctx.Bool(validatorOutputFlag),
	)
	if nil != err {
		return
	}

	log.WithField("pid", pid).Info("Consensus Layer Validator is ready")

	return
}

func stopValidator(ctx *cli.Context) (err error) {
	log.Info("Stopping Validator")
	dataDir := ctx.String(validatorDatadirFlag)

	err = clientDependencies[validatorDependencyName].Stop(
		dataDir,
	)
	if nil != err {
		return
	}

	return
}

func startCLStatsClient(ctx *cli.Context) (err error) {
	log.WithField("dependencyTag", CLStatsClientTag).Info("Starting CL Stats Client")
	dataDir := ctx.String(clStatsClientDatadirFlag)
	pid, err := clientDependencies[CLStatsClientDependencyName].Run(
		CLStatsClientTag,
		dataDir,
		CLStatsClientRuntimeFlags,
		ctx.Bool(clStatsOutputFlag),
	)
	if nil != err {
		return
	}

	log.WithField("pid", pid).Info("CL Stats Client is ready")

	return
}

func stopCLStatsClient(ctx *cli.Context) (err error) {
	log.Info("Stopping CL Stats Client")
	dataDir := ctx.String(clStatsClientDatadirFlag)

	err = clientDependencies[CLStatsClientDependencyName].Stop(
		dataDir,
	)
	if nil != err {
		return
	}

	return
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
