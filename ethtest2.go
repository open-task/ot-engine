package main

import (
	openTask "./contracts"
	"context"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"math/big"
	"strings"
)

func main() {
	client, err := ethclient.Dial("https://rinkeby.infura.io/v3/e17969db9bc94e75a474b3d3c5257a75")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("we have a connection now.")
	}

	_ = client // we'll use this in the upcoming sections

	contractAddress := common.HexToAddress("0xEB6af3deD23E2FA790C2D72B68e21dCC05c439d4") // Rinkeby

	// filter logs
	const FromBlock = 3244562;
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

	contractAbi, err := abi.JSON(strings.NewReader(string(openTask.OpenTaskABI)))
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		fmt.Printf("TxHash: %s\n", vLog.TxHash.Hex())

		for j, vTopic := range vLog.Topics {
			fmt.Printf("\t[Topic%d]: %s\n", j, vTopic.String())
		}
		fmt.Printf("Data: %v\n", vLog.Data)

		event := struct {
			missionId string
			rewardInWei *big.Int
		}{}
		err := contractAbi.Unpack(&event, "Publish", vLog.Data)
		if err != nil {
			log.Fatal(err)
		}

		fmt.Printf("missionId: %s\n", event.missionId)
		fmt.Printf("rewardInWei: %v\n", event.rewardInWei)
	}
}