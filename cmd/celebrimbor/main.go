package main

import (
	"fmt"
	"github.com/ethereum/go-ethereum/common/math"
	joonix "github.com/joonix/log"
	"github.com/lukso-network/lukso-orchestrator/orchestrator/node"
	"github.com/lukso-network/lukso-orchestrator/orchestrator/vanguardchain/client"
	"github.com/lukso-network/lukso-orchestrator/shared/cmd"
	"github.com/lukso-network/lukso-orchestrator/shared/journald"
	"github.com/lukso-network/lukso-orchestrator/shared/logutil"
	"github.com/sirupsen/logrus"
	"github.com/urfave/cli/v2"
	prefixed "github.com/x-cray/logrus-prefixed-formatter"
	"os"
	"runtime"
	runtimeDebug "runtime/debug"
	"sync"
	"time"
)

// We want to spin also 3 libraries at once, and secretly rule them by cli. It matches for me somehow

// This binary will also support only some of the possible networks.
// Make a pull request to attach your network.
// We are also very open to any improvements. Please make some issue or hackmd proposal to make it better.
// Join our silesiacoin discord https://discord.gg/RX6GsNr to ask some questions

const (
	ubuntu  = "linux"
	macos   = "darwin"
	windows = "windows"
)

var (
	appName               = "celebrimbor"
	elTag                 string
	validatorTag          string
	CLTag                 string
	log                   = logrus.WithField("prefix", appName)
	systemOs              string
	ELRuntimeFlags        []string
	validatorRuntimeFlags []string
	CLRuntimeFlags        []string
)

func init() {
	allFlags := make([]cli.Flag, 0)
	allFlags = append(allFlags, ELFlags...)
	allFlags = append(allFlags, validatorFlags...)
	allFlags = append(allFlags, CLFlags...)
	allFlags = append(allFlags, appFlags...)

	appFlags = cmd.WrapFlags(allFlags)
}

func main() {
	app := cli.App{}
	app.Name = appName
	app.Usage = "Spins all merge ecosystem components"
	app.Flags = appFlags
	app.Action = downloadAndRunBinaries

	app.Before = func(ctx *cli.Context) error {
		format := ctx.String(cmd.LogFormat.Name)
		switch format {
		case "text":
			formatter := new(prefixed.TextFormatter)
			formatter.TimestampFormat = "2006-01-02 15:04:05"
			formatter.FullTimestamp = true
			// If persistent log files are written - we disable the log messages coloring because
			// the colors are ANSI codes and seen as gibberish in the log files.
			formatter.DisableColors = ctx.String(cmd.LogFileName.Name) != ""
			logrus.SetFormatter(formatter)
		case "fluentd":
			f := joonix.NewFormatter()
			if err := joonix.DisableTimestampFormat(f); err != nil {
				panic(err)
			}
			logrus.SetFormatter(f)
		case "json":
			logrus.SetFormatter(&logrus.JSONFormatter{})
		case "journald":
			if err := journald.Enable(); err != nil {
				return err
			}
		default:
			return fmt.Errorf("unknown log format %s", format)
		}

		logFileName := ctx.String(cmd.LogFileName.Name)
		if logFileName != "" {
			if err := logutil.ConfigurePersistentLogging(logFileName); err != nil {
				log.WithError(err).Error("Failed to configuring logging to disk.")
			}
		}

		runtime.GOMAXPROCS(runtime.NumCPU())

		setupOperatingSystem()

		// el related parsing
		elTag = ctx.String(ELTagFlag)
		ELRuntimeFlags = prepareELFlags(ctx)

		// Validator related parsing
		validatorTag = ctx.String(validatorTagFlag)
		validatorRuntimeFlags = prepareValidatorFlags(ctx)

		// CL related parsing
		CLTag = ctx.String(CLTagFlag)
		CLRuntimeFlags = prepareCLFlags(ctx)

		return nil
	}

	defer func() {
		if x := recover(); x != nil {
			log.Errorf("Runtime panic: %v\n%v", x, string(runtimeDebug.Stack()))
			panic(x)
		}
	}()

	err := app.Run(os.Args)

	if nil != err {
		log.Error(err.Error())
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

	err = downloadValidator(ctx)

	if nil != err {
		return
	}

	err = downloadCL(ctx)

	if nil != err {
		return
	}

	err = downloadConfig(ctx)

	if nil != err {
		return
	}

	err = startEL(ctx)

	if nil != err {
		return
	}

	time.Sleep(time.Second * 6)

	err = startCL(ctx)

	if nil != err {
		return
	}

	time.Sleep(time.Second * 3)

	err = startValidator(ctx)

	if nil != err {
		return
	}

	// TODO: remove this
	return startOrchestrator(ctx)
}

func downloadEL(ctx *cli.Context) (err error) {
	log.WithField("dependencyTag", elTag).Info("I am downloading el")
	elDataDir := ctx.String(ELDatadirFlag)
	err = clientDependencies[ELDependencyName].Download(elTag, elDataDir)

	return
}

func downloadGenesis(ctx *cli.Context) (err error) {
	log.WithField("dependencyTag", elTag).Info("I am downloading el genesis")
	elDataDir := ctx.String(ELDatadirFlag)
	err = clientDependencies[ELGenesisDependencyName].Download(elTag, elDataDir)

	if nil != err {
		return
	}

	log.WithField("dependencyTag", CLTag).Info("I am downloading CL genesis")
	CLDataDir := ctx.String(CLDatadirFlag)
	err = clientDependencies[CLGenesisDependencyName].Download(CLTag, CLDataDir)

	return
}

func downloadConfig(ctx *cli.Context) (err error) {
	log.WithField("dependencyTag", CLTag).Info("I am downloading CL config")
	CLDataDir := ctx.String(CLDatadirFlag)
	err = clientDependencies[CLConfigDependencyName].Download(CLTag, CLDataDir)

	return
}

func downloadCL(ctx *cli.Context) (err error) {
	log.WithField("dependencyTag", CLTag).Info("I am downloading CL")
	CLDataDir := ctx.String(CLDatadirFlag)
	err = clientDependencies[CLDependencyName].Download(CLTag, CLDataDir)

	return
}

func downloadValidator(ctx *cli.Context) (err error) {
	log.WithField("dependencyTag", validatorTag).Info("I am downloading validator")
	validatorDataDir := ctx.String(CLDatadirFlag)
	err = clientDependencies[validatorDependencyName].Download(CLTag, validatorDataDir)

	return
}

func startEL(ctx *cli.Context) (err error) {
	log.WithField("dependencyTag", elTag).Info("I am running genesis.json init")
	elDataDir := ctx.String(ELDatadirFlag)
	elGenesisArguments := []string{
		"init",
		clientDependencies[ELGenesisDependencyName].ResolveBinaryPath(elTag, elDataDir),
		"--datadir",
		elDataDir,
	}

	err = clientDependencies[ELDependencyName].Run(
		elTag,
		elDataDir,
		elGenesisArguments,
		ctx.Bool(ELOutputFlag),
	)

	if nil != err {
		return
	}

	time.Sleep(time.Second * 3)

	log.WithField("dependencyTag", elTag).Info("I am running execution engine")
	err = clientDependencies[ELDependencyName].Run(
		elTag,
		elDataDir,
		ELRuntimeFlags,
		ctx.Bool(ELOutputFlag),
	)

	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(1)

	go func() {
		for {
			_, currentErr := os.Stat(DefaultELRPCEndpoint)
			if nil == currentErr {
				log.Info("el ipc is up")
				waitGroup.Done()

				return
			}

			if os.IsNotExist(currentErr) {
				time.Sleep(time.Millisecond * 50)
				log.Info("el ipc is dead")

				continue
			}

			panic(err)
		}
	}()

	waitGroup.Wait()

	return
}

func startCL(ctx *cli.Context) (err error) {
	log.WithField("dependencyTag", CLTag).Info("I am running CL")
	CLDataDir := ctx.String(CLDatadirFlag)
	err = clientDependencies[CLDependencyName].Run(
		CLTag,
		CLDataDir,
		CLRuntimeFlags,
		ctx.Bool(CLOutputFlag),
	)

	if nil != err {
		return
	}

	clGrpcEndpoint := ctx.String(CLGrpcEndpoint)

	if "" == clGrpcEndpoint {
		clGrpcEndpoint = DefaultELRPCEndpoint
	}

	for {
		log.Info("CL readiness check")
		time.Sleep(time.Millisecond * 250)
		vanClient, currentErr := client.Dial(
			ctx.Context,
			clGrpcEndpoint,
			time.Second,
			32,
			math.MaxInt32,
		)

		if nil != currentErr {
			log.WithField("cause", "CL not ready yet").Error(currentErr)
			continue
		}

		log.Info("CL is ready")
		vanClient.Close()
		break
	}

	return
}

func startValidator(ctx *cli.Context) (err error) {
	// First command should be to create wallet or prompt to do this by your own. This is one-time
	log.WithField("dependencyTag", validatorTag).Info("I am running CL")
	validatorDataDir := ctx.String(CLDatadirFlag)
	err = clientDependencies[validatorDependencyName].Run(
		validatorTag,
		validatorDataDir,
		validatorRuntimeFlags,
		false,
	)

	return
}

func startOrchestrator(ctx *cli.Context) (err error) {
	verbosity := ctx.String(cmd.VerbosityFlag.Name)
	level, err := logrus.ParseLevel(verbosity)
	if err != nil {
		return err
	}
	logrus.SetLevel(level)

	log.WithField("ELFlags", ELRuntimeFlags).
		WithField("CLFlags", CLRuntimeFlags).
		WithField("validatorFlags", validatorRuntimeFlags).Info("\n I will try to run setup with this additional flags \n")

	orchestrator, err := node.New(ctx)
	if err != nil {
		return err
	}
	orchestrator.Start()
	return nil
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	return !os.IsNotExist(err)
}
