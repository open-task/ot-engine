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
)

func main() {
	//const Rinkeby= "https://rinkeby.infura.io/v3/e17969db9bc94e75a474b3d3c5257a75"
	//const Rinkeby= "wss://rinkeby.infura.io/ws"
	const Rinkeby= "ws://172.31.24.221:8546"
	const OpenTaskAddress= "0xEB6af3deD23E2FA790C2D72B68e21dCC05c439d4"
	const FromBlock= 3244562;
	const (
		PublishSig = "0xa58a8cad8fe1a4abc834a5050930a0f50051e95838fefcc700ec7eebfa80eaf6" // keccak256("Publish(string,uint256)")
		SolveSig = "0x4eced7034db18ccd56b05d63eabae6e3ef4366544223d65cf6d9a0945969bb30"
		AcceptSig = "0xc356d4f9240e33d406223d0f76d7cbf13e019ab08e37d9598b045223b33685b2" //
		RejectSig = "0x5a074a7aea2ae96d3d465d04e81f21320d03a4eb77c3f298647ff4fbdce6787d" //
	)

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
				case PublishSig:
					fmt.Println("Publish")
				case SolveSig:
					fmt.Println("Solve")
				case AcceptSig:
					fmt.Println("Accept")
				case RejectSig:
					fmt.Println("Reject")
				default:
					//
					fmt.Println("UNKNOWN Event Log")
				}
			}
		}
	}
}
