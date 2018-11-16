package main

import (
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/mesg-foundation/core/x/xsignal"
	mesg "github.com/mesg-foundation/go-service"

	"github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/libs/log"
	"github.com/tendermint/tendermint/node"
	"github.com/tendermint/tendermint/p2p"
	"github.com/tendermint/tendermint/privval"
	"github.com/tendermint/tendermint/types"
)

// Return codes for the examples
const (
	CodeTypeOK            uint32 = 0
	CodeTypeEncodingError uint32 = 1
	CodeTypeBadNonce      uint32 = 2
	CodeTypeUnauthorized  uint32 = 3
)

type input struct {
	Number int `json:"number"`
}

type output struct {
	Code int `json:"code"`
}

func handler(execution *mesg.Execution) (string, mesg.Data) {
	var input input
	if err := execution.Data(&input); err != nil {
		return "success", &output{1}
	}

	resp, err := http.Get("localhost:26657/broadcast_tx_commit?tx=0x00")
	if err != nil {
		return "success", &output{2}
	}

	if resp.StatusCode != http.StatusOK {
		return "success", &output{3}
	}

	return "success", &output{0}
}

func start() error {
	if os.Getenv("MESG") != "no" {
		service, err := mesg.New()
		if err != nil {
			return err
		}
		err = service.Listen(mesg.Task("broadcast", handler))
		if err != nil {
			return err
		}
	}

	return nil
}

var (
	cfg          = config.DefaultConfig()
	logger       = log.NewTMLogger(log.NewSyncWriter(os.Stdout))
	nodeProvider = node.DefaultNewNode
)

func tendermintInit() error {
	// private validator
	privValFile := cfg.PrivValidatorFile()
	var pv *privval.FilePV
	if common.FileExists(privValFile) {
		pv = privval.LoadFilePV(privValFile)
		logger.Info("Found private validator", "path", privValFile)
	} else {
		pv = privval.GenFilePV(privValFile)
		pv.Save()
		logger.Info("Generated private validator", "path", privValFile)
	}

	nodeKeyFile := cfg.NodeKeyFile()
	if common.FileExists(nodeKeyFile) {
		logger.Info("Found node key", "path", nodeKeyFile)
	} else {
		if _, err := p2p.LoadOrGenNodeKey(nodeKeyFile); err != nil {
			return err
		}
		logger.Info("Generated node key", "path", nodeKeyFile)
	}

	// genesis file
	genFile := cfg.GenesisFile()
	if common.FileExists(genFile) {
		logger.Info("Found genesis file", "path", genFile)
	} else {
		genDoc := types.GenesisDoc{
			ChainID:         fmt.Sprintf("test-chain-%v", common.RandStr(6)),
			GenesisTime:     time.Now(),
			ConsensusParams: types.DefaultConsensusParams(),
		}
		genDoc.Validators = []types.GenesisValidator{{
			PubKey: pv.GetPubKey(),
			Power:  10,
		}}

		if err := genDoc.SaveAs(genFile); err != nil {
			return err
		}
		logger.Info("Generated genesis file", "path", genFile)
	}

	return nil
}

func tendermintNode() error {
	n, err := nodeProvider(cfg, logger)
	if err != nil {
		return fmt.Errorf("Failed to create node: %v", err)
	}

	// Stop upon receiving SIGTERM or CTRL-C
	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		for sig := range c {
			logger.Error(fmt.Sprintf("captured %v, exiting...", sig))
			if n.IsRunning() {
				n.Stop()
			}
			os.Exit(1)
		}
	}()

	if err := n.Start(); err != nil {
		return fmt.Errorf("Failed to start node: %v", err)
	}
	logger.Info("Started node", "nodeInfo", n.Switch().NodeInfo())

	return nil
}

func main() {
	os.MkdirAll("/tendermint/config")
	cfg.BaseConfig.RootDir = "/tendermint"
	cfg.BaseConfig.ProxyApp = "mesg"

	if err := tendermintInit(); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	if err := tendermintNode(); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}

	if err := start(); err != nil {
		fmt.Print(err)
		os.Exit(1)
	}
	<-xsignal.WaitForInterrupt()
}
