package collect

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum"
	"math/big"
	"context"
	"github.com/ethereum/go-ethereum/crypto"
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

	publishSig := []byte("Publish(string,uint256)")
	solveSig := []byte("Solve(string,string,string)")
	acceptSig := []byte("Accept(string)")
	rejectSig := []byte("Reject(string)")
	confirmSig := []byte("Confirm(string,string)")

	publishSigHash := crypto.Keccak256Hash(publishSig)
	solveSigHash := crypto.Keccak256Hash(solveSig)
	acceptSigHash := crypto.Keccak256Hash(acceptSig)
	rejectSigHash := crypto.Keccak256Hash(rejectSig)
	confirmSigHash := crypto.Keccak256Hash(confirmSig)

	fmt.Printf(`publishSigHash: %s
solveSigHash: %s
acceptSigHash: %s
rejectSigHash: %s
confirmSigHash: %s
`, publishSigHash.Hex(), solveSigHash.Hex(), acceptSigHash.Hex(), rejectSigHash.Hex(), confirmSigHash.Hex())

	for i, vLog := range logs {
		fmt.Printf("TxHash[%d]: %s\n", i, vLog.TxHash.Hex())

		if len(vLog.Topics) >= 1 {
			fmt.Printf("Sig(Topic0): %s\n", vLog.Topics[0].String())
			switch vLog.Topics[0].String() {
			case publishSigHash.Hex():
				fmt.Println("Publish")
				Publish(vLog.Topics, vLog.Data)
			case solveSigHash.Hex():
				fmt.Println("Solve")
				Solve(vLog.Topics, vLog.Data)
			case acceptSigHash.Hex():
				fmt.Println("Accept")
				Accept(vLog.Topics, vLog.Data)
			case rejectSigHash.Hex():
				fmt.Println("Reject")
				Reject(vLog.Topics, vLog.Data)
			case confirmSigHash.Hex():
				fmt.Println("Confirm")
				Confirm(vLog.Topics, vLog.Data)
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
