package collect

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum"
	"math/big"
	"context"
	"github.com/xyths/ot-engine/decode"
)

func Collect(server string, address string, from int, to int, eventType string) {
	client, err := ethclient.Dial(server)
	if err != nil {
		log.Fatal(err)
	}
	contractAddress := common.HexToAddress(address)

	query := ethereum.FilterQuery{
		FromBlock: big.NewInt(int64(from)),
		ToBlock:   big.NewInt(int64(to)),
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
		fmt.Printf("TxHash[%d]: %s\n", i, vLog.TxHash.Hex())

		if len(vLog.Topics) >= 1 {
			fmt.Printf("Sig(Topic0): %s\n", vLog.Topics[0].String())
			switch vLog.Topics[0].String() {
			case decode.PublishSig:
				fmt.Println("Publish")
				Publish(vLog.Topics)
			case decode.SolveSig:
				fmt.Println("Solve")
				Solve(vLog.Topics)
			case decode.AcceptSig:
				fmt.Println("Accept")
				Accept(vLog.Topics)
			case decode.RejectSig:
				fmt.Println("Reject")
				Reject(vLog.Topics)
			default:
				//
				fmt.Println("UNKNOWN Event Log")
			}
		}

		for j, vTopic := range vLog.Topics {
			fmt.Printf("\t[Topic%d]: %s\n", j, vTopic.String())
		}
		//fmt.Printf("Data: %v\n", vLog.Data)
	}
}
