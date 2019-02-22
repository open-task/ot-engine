package main

import (
	openTask "github.com/open-task/ot-engine/contracts"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
)

func main() {
	client, err := ethclient.Dial("https://rinkeby.infura.io/v3/e17969db9bc94e75a474b3d3c5257a75")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("we have a connection now.")
	}

	_ = client // we'll use this in the upcoming sections

	contractAddress := common.HexToAddress("0xF562a7c51a158ae6E6170Ef7905af5d1cE43d24A") // Rinkeby
	instance, err := openTask.NewOpenTask(contractAddress, client)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("contract is loaded")
	_ = instance

	owner, err := instance.Owner(nil);
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("contract owner is %s\n", owner.Hex()); // should be 0x1c635f4756ED1dD9Ed615dD0A0Ff10E3015cFa7b


	// filter logs
	const FromBlock = 3146781;
	const ToBlock = FromBlock + 1;
	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(FromBlock),
		ToBlock:   big.NewInt(ToBlock),
		Addresses: []common.Address{
			contractAddress,
		},
	}

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		log.Fatal(err)
	}

	for i, vLog := range logs {
		fmt.Printf("Event: %d\n", i)
		// OwnershipTransferred
		// [topic0] 0x8be0079c531659141344cd1fd0a4f28419497f9722a3daafe3b4186f6b6457e0
		// [topic1] 0x0000000000000000000000000000000000000000000000000000000000000000
		// [topic2] 0x0000000000000000000000001c635f4756ed1dd9ed615dd0a0ff10e3015cfa7b
		// https://rinkeby.etherscan.io/tx/0xb3cfac13487637c36a6bbfb9ee0d9a0096e7355e6cb443c80e55cfa1377569b5#eventlog
		for j, vTopic := range vLog.Topics {
			fmt.Printf("\t[Topic%d] %s\n", j, vTopic.String())
		}
	}
}