package utils

import (
	"context"
	"database/sql"
	"fmt"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/ethclient"
	"github.com/open-task/ot-engine/database"
	"github.com/open-task/ot-engine/process"
	"log"
	"math/big"
	"time"

	"github.com/open-task/ot-engine/node"
)

func Download(cfg node.Config, db *sql.DB) {
	// redail every time
	client, err := ethclient.Dial(cfg.Server)
	if err != nil {
		log.Fatal(err)
	}

	address := common.HexToAddress(cfg.Contract)
	query := ethereum.FilterQuery{
		Addresses: []common.Address{address},
	}

	from := getFrom(db)
	//from.Add(from, big.NewInt(1))
	query.FromBlock = from

	logs, err := client.FilterLogs(context.Background(), query)
	if err != nil {
		log.Fatal(err)
	}

	shanghai, _ := time.LoadLocation("Asia/Shanghai")

	for i, vLog := range logs {
		fmt.Printf("TxHash[%d]: %s\n", i, vLog.TxHash.Hex())
		header, err := client.HeaderByHash(context.Background(), vLog.BlockHash)
		if err != nil {
			fmt.Println(err)
			continue
		}
		txTime := header.Time
		timeStr := time.Unix(txTime.Int64(), 0).In(shanghai).String()

		tx, _, err := client.TransactionByHash(context.Background(), vLog.TxHash)
		if err != nil {
			fmt.Println(err)
			continue
		}

		if from.Cmp(header.Number) == -1 {
			from = header.Number
		}
		signer := types.NewEIP155Signer(tx.ChainId())
		sender, err := signer.Sender(tx)

		err = process.ParseOTLog(vLog, timeStr, sender.String(), db)
		if err != nil {
			fmt.Println(err)
			continue
		}
	}
	setFrom(db, from)
}

func getFrom(db *sql.DB) *big.Int {
	from, err := database.GetFrom(db)
	if err != nil {
		return big.NewInt(0)
	}
	fmt.Printf("Load config from: %s\n", from)
	return from
}

func setFrom(db *sql.DB, from *big.Int) bool {
	fmt.Printf("Update config from: %s\n", from)
	err := database.SetFrom(db, from)
	if err != nil {
		return false
	}
	return true
}
