package main

import (
	"encoding/binary"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	mesg "github.com/mesg-foundation/go-service"

	"github.com/tendermint/tendermint/abci/types"
	"github.com/tendermint/tendermint/config"
	"github.com/tendermint/tendermint/libs/common"
	"github.com/tendermint/tendermint/libs/log"
	"github.com/tendermint/tendermint/node"
	"github.com/tendermint/tendermint/p2p"
	"github.com/tendermint/tendermint/privval"
	"github.com/tendermint/tendermint/proxy"
	xxxtypes "github.com/tendermint/tendermint/types"
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
	nodeProvider = DefaultNewNode
)

// DefaultNewNode returns a Tendermint node with default settings for the
// PrivValidator, ClientCreator, GenesisDoc, and DBProvider.
// It implements NodeProvider.
func DefaultNewNode(cfg *config.Config, logger log.Logger) (*node.Node, error) {
	// Generate node PrivKey
	nodeKey, err := p2p.LoadOrGenNodeKey(cfg.NodeKeyFile())
	if err != nil {
		return nil, err
	}
	return node.NewNode(cfg,
		privval.LoadOrGenFilePV(cfg.PrivValidatorFile()),
		nodeKey,
		DefaultClientCreator(cfg.ProxyApp, cfg.ABCI, cfg.DBDir()),
		node.DefaultGenesisDocProviderFunc(cfg),
		node.DefaultDBProvider,
		node.DefaultMetricsProvider(cfg.Instrumentation),
		logger,
	)
}

func DefaultClientCreator(addr, transport, dbDir string) proxy.ClientCreator {
	return proxy.NewLocalClientCreator(NewCounterApplication(false))
}

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
		genDoc := xxxtypes.GenesisDoc{
			ChainID:         fmt.Sprintf("test-chain-%v", common.RandStr(6)),
			GenesisTime:     time.Now(),
			ConsensusParams: xxxtypes.DefaultConsensusParams(),
		}
		genDoc.Validators = []xxxtypes.GenesisValidator{{
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
	os.MkdirAll("/tendermint/config", 0777)
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
	select {}
}

type CounterApplication struct {
	types.BaseApplication

	hashCount int
	txCount   int
	serial    bool
}

func NewCounterApplication(serial bool) *CounterApplication {
	return &CounterApplication{serial: serial}
}

func (app *CounterApplication) Info(req types.RequestInfo) types.ResponseInfo {
	return types.ResponseInfo{Data: fmt.Sprintf("{\"hashes\":%v,\"txs\":%v}", app.hashCount, app.txCount)}
}

func (app *CounterApplication) SetOption(req types.RequestSetOption) types.ResponseSetOption {
	key, value := req.Key, req.Value
	if key == "serial" && value == "on" {
		app.serial = true
	} else {
		/*
			TODO Panic and have the ABCI server pass an exception.
			The client can call SetOptionSync() and get an `error`.
			return types.ResponseSetOption{
				Error: fmt.Sprintf("Unknown key (%s) or value (%s)", key, value),
			}
		*/
		return types.ResponseSetOption{}
	}

	return types.ResponseSetOption{}
}

func (app *CounterApplication) DeliverTx(tx []byte) types.ResponseDeliverTx {
	if app.serial {
		if len(tx) > 8 {
			return types.ResponseDeliverTx{
				Code: CodeTypeEncodingError,
				Log:  fmt.Sprintf("Max tx size is 8 bytes, got %d", len(tx))}
		}
		tx8 := make([]byte, 8)
		copy(tx8[len(tx8)-len(tx):], tx)
		txValue := binary.BigEndian.Uint64(tx8)
		if txValue != uint64(app.txCount) {
			return types.ResponseDeliverTx{
				Code: CodeTypeBadNonce,
				Log:  fmt.Sprintf("Invalid nonce. Expected %v, got %v", app.txCount, txValue)}
		}
	}
	app.txCount++
	return types.ResponseDeliverTx{Code: CodeTypeOK}
}

func (app *CounterApplication) CheckTx(tx []byte) types.ResponseCheckTx {
	if app.serial {
		if len(tx) > 8 {
			return types.ResponseCheckTx{
				Code: CodeTypeEncodingError,
				Log:  fmt.Sprintf("Max tx size is 8 bytes, got %d", len(tx))}
		}
		tx8 := make([]byte, 8)
		copy(tx8[len(tx8)-len(tx):], tx)
		txValue := binary.BigEndian.Uint64(tx8)
		if txValue < uint64(app.txCount) {
			return types.ResponseCheckTx{
				Code: CodeTypeBadNonce,
				Log:  fmt.Sprintf("Invalid nonce. Expected >= %v, got %v", app.txCount, txValue)}
		}
	}
	return types.ResponseCheckTx{Code: CodeTypeOK}
}

func (app *CounterApplication) Commit() (resp types.ResponseCommit) {
	app.hashCount++
	if app.txCount == 0 {
		return types.ResponseCommit{}
	}
	hash := make([]byte, 8)
	binary.BigEndian.PutUint64(hash, uint64(app.txCount))
	return types.ResponseCommit{Data: hash}
}

func (app *CounterApplication) Query(reqQuery types.RequestQuery) types.ResponseQuery {
	switch reqQuery.Path {
	case "hash":
		return types.ResponseQuery{Value: []byte(fmt.Sprintf("%v", app.hashCount))}
	case "tx":
		return types.ResponseQuery{Value: []byte(fmt.Sprintf("%v", app.txCount))}
	default:
		return types.ResponseQuery{Log: fmt.Sprintf("Invalid query path. Expected hash or tx, got %v", reqQuery.Path)}
	}
}
