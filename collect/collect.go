package collect

import (
	"fmt"
	"github.com/ethereum/go-ethereum/ethclient"
	"log"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum"
	"math/big"
	"context"
	"database/sql"
	_ "github.com/go-sql-driver/mysql"
	"github.com/xyths/ot-engine/database"
	. "github.com/xyths/ot-engine/decode"
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

	db, err := sql.Open("mysql", "engine:decopentask@/ot_local")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	for i, vLog := range logs {
		fmt.Printf("TxHash[%d]: %s\n", i, vLog.TxHash.Hex())

		if len(vLog.Topics) >= 1 {
			fmt.Printf("Sig(Topic0): %s\n", vLog.Topics[0].String())
			switch vLog.Topics[0].String() {
			case PublishSigHash.Hex():
				fmt.Println("Publish")
				row, err := Publish(vLog)
				if err != nil {
					continue
				}
				err = database.Publish(db, row)
				if err != nil {
					log.Println("Got error when insert to database.")
				}
			case SolveSigHash.Hex():
				fmt.Println("Solve")
				row, err := Solve(vLog)
				if err != nil {
					continue
				}
				err = database.Solve(db, row)
				if err != nil {
					log.Println("Got error when insert to database.")
				}
			case AcceptSigHash.Hex():
				fmt.Println("Accept")
				row, err := Accept(vLog)
				if err != nil {
					continue
				}
				err = database.Accept(db, row)
				if err != nil {
					log.Println("Got error when insert to database.")
				}
			case RejectSigHash.Hex():
				fmt.Println("Reject")
				row, err := Reject(vLog)
				if err != nil {
					continue
				}
				err = database.Reject(db, row)
				if err != nil {
					log.Println("Got error when insert to database.")
				}
			case ConfirmSigHash.Hex():
				fmt.Println("Confirm")
				row, err := Confirm(vLog)
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

		for j, vTopic := range vLog.Topics {
			fmt.Printf("\t[Topic%d]: %s\n", j, vTopic.String())
		}
		//fmt.Printf("Data: %v\n", vLog.Data)
	}
}
