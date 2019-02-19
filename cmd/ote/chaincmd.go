package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/xyths/ot-engine/cmd/utils"
	"github.com/xyths/ot-engine/process"
	"gopkg.in/urfave/cli.v2"
	"log"
	"math/big"
)

var (
	deCommand = &cli.Command{
		Action:  download,
		Name:    "download",
		Aliases: []string{"d"},
		Usage:   "Download event log from blockchain",
		Flags: []cli.Flag{
			utils.FromFlag,
			utils.ToFlag,
		},
	}
	serveCommand = &cli.Command{
		Action:  serve,
		Name:    "serve",
		Aliases: []string{"s"},
		Usage:   "Serve as a HTTP server",
		Flags: []cli.Flag{
		},
	}
	listenCommand = &cli.Command{
		Action:  listen,
		Name:    "listen",
		Aliases: []string{"l"},
		Usage:   "Listen to blockchain and process the event log",
		Flags: []cli.Flag{
		},
	}
)

func download(ctx *cli.Context) (err error) {
	cfg := LoadConfig(ctx)
	fmt.Printf("server: %s, contract: %s\n", cfg.Node.Server, cfg.Node.Contract)

	from := ctx.Int64(utils.FromFlag.Name)
	to := ctx.Int64(utils.ToFlag.Name)

	client, err := ethclient.Dial(cfg.Node.Server)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(cfg.Node.Contract)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{address},
	}
	if from != 0 {
		query.FromBlock = big.NewInt(from)
	} else {
		if cfg.Node.FromFlag {
			query.FromBlock = cfg.Node.FromBlock
		}
	}
	if to != 0 {
		query.ToBlock = big.NewInt(to)
	} else {
		if cfg.Node.ToFlag {
			query.ToBlock = cfg.Node.ToBlock
		}
	}

	logs, err := client.FilterLogs(context.Background(), query)

	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		log.Fatal(err)
	}

	for i, vLog := range logs {
		fmt.Printf("TxHash[%d]: %s\n", i, vLog.TxHash.Hex())
		err := process.ParseOTLog(vLog, db)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	return err
}

func serve(ctx *cli.Context) error {
	engine := makeConfigEngine(ctx)
	engine.Setup()
	engine.Serve()
	return nil
}

func listen(ctx *cli.Context) (err error) {
	cfg := LoadConfig(ctx)
	fmt.Printf("server: %s, contract: %s\n", cfg.Node.Server, cfg.Node.Contract)

	client, err := ethclient.Dial(cfg.Node.Server)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(cfg.Node.Contract)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{address},
	}
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)

	if err != nil {
		log.Fatal(err)
	}

	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			fmt.Println(err)
		case vLog := <-logs:
			err := process.ParseOTLog(vLog, db)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
	return
}
