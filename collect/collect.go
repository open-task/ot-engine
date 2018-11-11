package collect

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum"
	"math/big"
	"context"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
	openTask "github.com/xyths/ot-engine/contracts"
	"github.com/xyths/ot-engine/decode"
)

//func Collect(client *Client, query FilterQuery) {
func Collect(server string, address string, from int, to int, eventType string) {
	fmt.Println("in collect now")

	client, err := ethclient.Dial(server)
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("we have a connection now.")
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

	contractAbi, err := abi.JSON(strings.NewReader(string(openTask.OpenTaskABI)))
	if err != nil {
		log.Fatal(err)
	}

	for _, vLog := range logs {
		fmt.Printf("TxHash: %s\n", vLog.TxHash.Hex())

		if len(vLog.Topics) >= 1 {
			switch vLog.Topics[0].String() {
			case decode.PublishSig:
				fmt.Println("Publish")
			case decode.SolveSig:
				fmt.Println("Solve")
			case decode.AcceptSig:
				fmt.Println("Accept")
			case decode.RejectSig:
				fmt.Println("Reject")
			default:
				//
				fmt.Println("UNKNOWN Event Log")
			}
		}

		for j, vTopic := range vLog.Topics {
			fmt.Printf("\t[Topic%d]: %s\n", j, vTopic.String())
		}
		fmt.Printf("Data: %v\n", vLog.Data)

		event := struct {
			missionId   string
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
