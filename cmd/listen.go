package main

import (
	"context"
	"fmt"
	"log"
	"math/big"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	. "../decode"
	"github.com/xyths/ot-engine/database"
	"github.com/xyths/ot-engine/collect"
	"database/sql"
	"flag"
	"../config"
)

const version string = "0.1.0"

var (
	h bool
	v bool

	configfile string
)

func init() {
	flag.BoolVar(&h, "h", false, "this help")
	flag.BoolVar(&v, "v", false, "show version and exit")

	flag.StringVar(&configfile, "c", "config.json", "`config`: config file")

	// 改变默认的 Usage
	flag.Usage = usage
}

func usage() {
	fmt.Fprintf(flag.CommandLine.Output(), `otl: OpenTask Listener, listen event log from Ethereum Network, version: %s
Usage: otl [-hv] [-c config]

Options:
`, version)
	flag.PrintDefaults()
}

func main() {
	flag.Parse()

	if h {
		flag.Usage()
		return
	}

	if v {
		fmt.Println("otl(OpenTask Listener):", version)
		return
	}

	cf, err := config.LoadConfig(configfile)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("data source name: ", cf.Dsn())
	//const Rinkeby= "https://rinkeby.infura.io/v3/e17969db9bc94e75a474b3d3c5257a75"
	const Rinkeby= "wss://rinkeby.infura.io/ws/v3/e17969db9bc94e75a474b3d3c5257a75"
	//const Rinkeby= "ws://172.31.24.221:8546"
	const OpenTaskAddress= "0xF562a7c51a158ae6E6170Ef7905af5d1cE43d24A"
	const FromBlock= 3244562;

	client, err := ethclient.Dial(Rinkeby)
	if err != nil {
		log.Fatal(err)
	}

	contractAddress := common.HexToAddress(OpenTaskAddress)
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(FromBlock),
		Addresses: []common.Address{contractAddress},
	}

	logs := make(chan types.Log)
	sub, err := client.SubscribeFilterLogs(context.Background(), query, logs)
	if err != nil {
		log.Fatal(err)
	}
	db, err := sql.Open("mysql", "engine:decopentask@/ot_local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for {
		select {
		case err := <-sub.Err():
			fmt.Println(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
			if len(vLog.Topics) >= 1 {
				switch vLog.Topics[0].String() {
				case PublishSigHash.Hex():
					fmt.Println("Publish")
					row, err := collect.Publish(vLog)
					if err != nil {
						continue
					}
					err = database.Publish(db, row)
					if err != nil {
						log.Println("Got error when insert to database.")
					}
				case SolveSigHash.Hex():
					fmt.Println("Solve")
					row, err := collect.Solve(vLog)
					if err != nil {
						continue
					}
					err = database.Solve(db, row)
					if err != nil {
						log.Println("Got error when insert to database.")
					}
				case AcceptSigHash.Hex():
					fmt.Println("Accept")
					row, err := collect.Accept(vLog)
					if err != nil {
						continue
					}
					err = database.Accept(db, row)
					if err != nil {
						log.Println("Got error when insert to database.")
					}
				case RejectSigHash.Hex():
					fmt.Println("Reject")
					row, err := collect.Reject(vLog)
					if err != nil {
						continue
					}
					err = database.Reject(db, row)
					if err != nil {
						log.Println("Got error when insert to database.")
					}
				case ConfirmSigHash.Hex():
					fmt.Println("Confirm")
					row, err := collect.Confirm(vLog)
					if err != nil {
						continue
					}
					err = database.Confirm(db, row)
					if err != nil {
						log.Println("Got error when insert to database.")
					}
				default:
					//
					fmt.Println("UNKNOWN Event Log")
				}
			}
		}
	}
}
