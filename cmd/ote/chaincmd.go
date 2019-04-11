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
	"os"
	"os/signal"
	"time"
)

const interval = "1m"

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

	pool, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		log.Fatal(err)
	}
	pool.SetConnMaxLifetime(0)
	pool.SetMaxIdleConns(2)
	pool.SetMaxOpenConns(2)
	fmt.Println("Database connection opened.")

	dbctx, stop := context.WithCancel(context.Background())
	defer stop()
	appSignal := make(chan os.Signal, 3)
	signal.Notify(appSignal, os.Interrupt)

	go func() {
		select {
		case <-appSignal:
			log.Println("download got stop signal, shutdown context now ...")
			stop()
		}
	}()

	shanghai, _ := time.LoadLocation("Asia/Shanghai")

	for i, vLog := range logs {
		fmt.Printf("TxHash[%d]: %s\n", i, vLog.TxHash.Hex())
		header, err := client.HeaderByHash(context.Background(), vLog.BlockHash)
		if err != nil {
			fmt.Println(err)
			continue
		}
		txTime := header.Time
		timeStr := time.Unix(int64(txTime), 0).In(shanghai).String()

		tx, _, err := client.TransactionByHash(context.Background(), vLog.TxHash)
		if err != nil {
			fmt.Println(err)
			continue
		}
		signer := types.NewEIP155Signer(tx.ChainId())
		sender, err := signer.Sender(tx)

		err = process.ParseOTLog(dbctx, vLog, timeStr, sender.String(), pool)
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
		log.Fatal("Error when dail:", err)
	}

	address := common.HexToAddress(cfg.Node.Contract)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{address},
	}
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)

	if err != nil {
		log.Fatal("Error when subscribe:", err)
	}

	pool, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		log.Fatal(err)
	}
	pool.SetConnMaxLifetime(0)
	pool.SetMaxIdleConns(3)
	pool.SetMaxOpenConns(3)
	fmt.Println("Database connection opened.")
	dbctx, stop := context.WithCancel(context.Background())
	defer stop()

	shanghai, _ := time.LoadLocation("Asia/Shanghai")

	appSignal := make(chan os.Signal, 3)
	signal.Notify(appSignal, os.Interrupt)
	for {
		select {
		case <-appSignal:
			log.Println("listen got stop signal, shutdown context now ...")
			stop()
			os.Exit(-1)
		case err := <-sub.Err():
			fmt.Errorf("Got Err when listen: %s", err)
			break // this will exit the program, and auto restart by supervisor in PRODUCTION
		case vLog := <-logs:
			log.Println("listen got a event log, try to process it now ...")
			header, err := client.HeaderByHash(context.Background(), vLog.BlockHash)
			if err != nil {
				log.Println("get header by hash error:", err)
				continue
			}
			txTime := header.Time
			timeStr := time.Unix(int64(txTime), 0).In(shanghai).String()

			tx, _, err := client.TransactionByHash(context.Background(), vLog.TxHash)
			if err != nil {
				log.Println("get tx by hash error:", err)
				continue
			}
			signer := types.NewEIP155Signer(tx.ChainId())
			sender, err := signer.Sender(tx)

			if err != nil {
				log.Println("get sender from tx error:", err)
				continue
			}

			err = process.ParseOTLog(dbctx, vLog, timeStr, sender.String(), pool)
			if err != nil {
				log.Println("ParseOTLog error:", err)
				continue
			}
			log.Println("event log process done.")
		}
	}
	fmt.Println("Listen process finished")
	return
}

func timer(ctx *cli.Context) (err error) {
	cfg := LoadConfig(ctx)
	fmt.Printf("server: %s, contract: %s\n", cfg.Node.Server, cfg.Node.Contract)

	pool, err := sql.Open("mysql", cfg.DSN())
	if err != nil {
		log.Fatal(err)
	}
	pool.SetConnMaxLifetime(0)
	pool.SetMaxIdleConns(3)
	pool.SetMaxOpenConns(3)
	fmt.Println("Database connection opened.")

	dbctx, stop := context.WithCancel(context.Background())
	defer stop()
	appSignal := make(chan os.Signal, 3)
	signal.Notify(appSignal, os.Interrupt)

	go func() {
		select {
		case <-appSignal:
			log.Println("timer got stop signal, shutdown context now ...")
			stop()
		}
	}()

	duration, _ := time.ParseDuration(interval)
	timer := time.NewTimer(duration)

	utils.Download(dbctx, cfg.Node, pool)
	fmt.Printf("Sleep %s...\n", interval)
	timer.Reset(duration)
	for {
		select {
		case <-timer.C:
			utils.Download(dbctx, cfg.Node, pool)
			fmt.Printf("Sleep %s...\n", interval)
			timer.Reset(duration)
		}
	}

	return err
}
