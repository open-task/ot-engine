package main

import (
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/xyths/ot-engine/process"
	"gopkg.in/urfave/cli.v2"
	"log"
)

var (
	deCommand = &cli.Command{
		Action:  de,
		Name:    "de",
		Aliases: []string{"d"},
		Usage:   "Download event log from blockchain",
		Flags: []cli.Flag{
		},
	}
	serveCommand = &cli.Command{
		Action:  serve,
		Name:    "serve",
		Aliases: []string{"s"},
		Usage:   "serve as a HTTP server",
		Flags: []cli.Flag{
		},
	}
	listenCommand = &cli.Command{
		Action:  listen,
		Name:    "listen",
		Aliases: []string{"l"},
		Usage:   "Listen to blockchain and process event log",
		Flags: []cli.Flag{
		},
	}
)

func de(ctx *cli.Context) (err error) {
	stack := makeConfigNode(ctx)
	fmt.Printf("server: %s, contract: %s\n", stack.Config.Server, stack.Config.Contract)

	client, err := ethclient.Dial(stack.Config.Server)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("we have a connection now.")
	}

	address := common.HexToAddress(stack.Config.Contract)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{address},
	}
	if stack.Config.FromFlag {
		query.FromBlock = stack.Config.FromBlock
	}
	if stack.Config.ToFlag {
		query.ToBlock = stack.Config.ToBlock
	}
	logs, err := client.FilterLogs(context.Background(), query)

	if err != nil {
		log.Fatal(err)
	}

	for i, vLog := range logs {
		fmt.Printf("TxHash[%d]: %s\n", i, vLog.TxHash.Hex())
		err := process.ParseOTLog(vLog)
		if err != nil {
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
	stack := makeConfigNode(ctx)
	fmt.Printf("server: %s, contract: %s\n", stack.Config.Server, stack.Config.Contract)

	client, err := ethclient.Dial(stack.Config.Server)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("we have a connection now.")
	}

	address := common.HexToAddress(stack.Config.Contract)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{address},
	}
	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)

	if err != nil {
		log.Fatal(err)
	}

	for {
		select {
		case err := <-sub.Err():
			fmt.Println(err)
		case vLog := <-logs:
			err := process.ParseOTLog(vLog)
			if err != nil {
				continue
			}
		}
	}
	return
}
