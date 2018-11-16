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
)

func main() {
	//const Rinkeby= "https://rinkeby.infura.io/v3/e17969db9bc94e75a474b3d3c5257a75"
	//const Rinkeby= "wss://rinkeby.infura.io/ws"
	const Rinkeby= "ws://172.31.24.221:8546"
	const OpenTaskAddress= "0xEB6af3deD23E2FA790C2D72B68e21dCC05c439d4"
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

	for {
		select {
		case err := <-sub.Err():
			log.Fatal(err)
		case vLog := <-logs:
			fmt.Println(vLog) // pointer to event log
			if len(vLog.Topics) >= 1 {
				switch vLog.Topics[0].String() {
				case PublishSigHash.Hex():
					fmt.Println("Publish")
				case SolveSigHash.Hex():
					fmt.Println("Solve")
				case AcceptSigHash.Hex():
					fmt.Println("Accept")
				case RejectSigHash.Hex():
					fmt.Println("Reject")
				case ConfirmSigHash.Hex():
					fmt.Println("Reject")
				default:
					//
					fmt.Println("UNKNOWN Event Log")
				}
			}
		}
	}
}
