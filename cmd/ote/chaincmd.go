package main

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/open-task/ot-engine/cmd/utils"
	"github.com/open-task/ot-engine/process"
	"gopkg.in/urfave/cli.v2"
	"log"
	"math/big"
	"time"
)

var (
	downloadCommand = &cli.Command{
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
	timerCommand = &cli.Command{
		Action:  timer,
		Name:    "timer",
		Aliases: []string{"t"},
		Usage:   "Timely download event log from blockchain",
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

	shanghai, _ := time.LoadLocation("Asia/Shanghai")

	for i, vLog := range logs {
		fmt.Printf("TxHash[%d]: %s\n", i, vLog.TxHash.Hex())
		header, err := client.HeaderByHash(context.Background(), vLog.BlockHash)
		if err != nil {
			fmt.Println(err)
			continue
		}
		txTime := header.Time
		timeStr := time.Unix(txTime.Int64(), 0).In(shanghai).String()

		tx, _, err := client.TransactionByHash(context.Background(), vLog.TxHash)
		if err != nil {
			fmt.Println(err)
			continue
		}
		signer := types.NewEIP155Signer(tx.ChainId())
		sender, err := signer.Sender(tx)

		err = process.ParseOTLog(vLog, timeStr, sender.String(), db)
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
	fmt.Println("Listen process started")
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

	shanghai, _ := time.LoadLocation("Asia/Shanghai")

	for {
		select {
		case err := <-sub.Err():
			fmt.Errorf("Got Err when listen: %s", err)
			break // this will exit the program, and auto restart by supervisor in PRODUCTION
		case vLog := <-logs:
			header, err := client.HeaderByHash(context.Background(), vLog.BlockHash)
			if err != nil {
				fmt.Println(err)
				continue
			}
			txTime := header.Time
			timeStr := time.Unix(txTime.Int64(), 0).In(shanghai).String()

			tx, _, err := client.TransactionByHash(context.Background(), vLog.TxHash)
			if err != nil {
				fmt.Println(err)
				continue
			}
			signer := types.NewEIP155Signer(tx.ChainId())
			sender, err := signer.Sender(tx)

			if err != nil {
				fmt.Println(err)
				continue
			}

			err = process.ParseOTLog(vLog, timeStr, sender.String(), db)
			if err != nil {
				fmt.Println(err)
				continue
			}
		}
	}
	fmt.Println("Listen process finished")
	return
}

func timer(ctx *cli.Context) (err error) {
	cfg := LoadConfig(ctx)
	fmt.Printf("server: %s, contract: %s\n", cfg.Node.Server, cfg.Node.Contract)

	db, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		log.Fatal(err)
	}

	duration, _ := time.ParseDuration("10s")
	timer := time.NewTimer(duration)

	utils.Download(cfg.Node, db)
	timer.Reset(duration)
	for {
		select {
		case <-timer.C:
			utils.Download(cfg.Node, db)
			timer.Reset(duration)
		}
	}

	return err
}
