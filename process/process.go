package process

import (
	"context"
	"database/sql"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/open-task/ot-engine/contracts"
	"github.com/open-task/ot-engine/database"
	. "github.com/open-task/ot-engine/decode"
	otTypes "github.com/open-task/ot-engine/types"
	"log"
	"math/big"
	"strings"
	"time"
)

func ParseOTLog(ctx context.Context, vLog types.Log, txTime string, sender string, pool *sql.DB) (err error) {
	log.Println("enter ParseOTLog")
	ctx, cancel := context.WithTimeout(ctx, 1*time.Second)
	defer cancel()

	if err := pool.PingContext(ctx); err != nil {
		log.Fatalf("unable to connect to database: %v", err)
	}

	switch vLog.Topics[0].String() {
	case PublishSigHash.Hex():
		log.Println("Publish event processing")
		row, err1 := Publish(vLog, txTime, sender)
		if err1 != nil {
			log.Println("Error when Parse Publish event:", err1)
			return err1
		}
		log.Println(row)
		err1 = database.Publish(ctx, pool, row)
		if err1 != nil {
			log.Println("Error when Insert Publish event:", err1)
			return err1
		}
	case SolveSigHash.Hex():
		log.Println("Solve event processing")
		row, err1 := Solve(vLog, txTime, sender)
		if err1 != nil {
			log.Println("Error when Parse Solve event:", err1)
			return err1
		}
		log.Println(row)
		err1 = database.Solve(ctx, pool, row)
		if err1 != nil {
			log.Println("Error when Insert Solve event:", err1)
			return err1
		}
	case AcceptSigHash.Hex():
		log.Println("Accept event processing")
		row, err1 := Accept(vLog, txTime, sender)
		if err1 != nil {
			log.Println("Error when Parse Accept event:", err1)
			return err1
		}
		log.Println(row)
		err1 = database.Accept(ctx, pool, row)
		if err1 != nil {
			log.Println("Error when Insert Accept event:", err1)
			return err1
		}
	case RejectSigHash.Hex():
		log.Println("Reject event processing")
		row, err1 := Reject(vLog, txTime, sender)
		if err1 != nil {
			log.Println("Error when Parse Reject event:", err1)
			return err1
		}
		log.Println(row)
		err1 = database.Reject(ctx, pool, row)
		if err1 != nil {
			log.Println("Error when Insert Reject event:", err1)
			return err1
		}
	case ConfirmSigHash.Hex():
		log.Println("Confirm event processing")
		row, err1 := Confirm(vLog, txTime, sender)
		if err1 != nil {
			log.Println("Error when Parse Confirm event:", err1)
			return err1
		}
		log.Println(row)
		err1 = database.Confirm(ctx, pool, row)
		if err1 != nil {
			log.Println("Error when Insert Confirm event:", err1)
			return err1
		}
	default:
		log.Println("UNKNOWN Event")
	}
	return
}

func Publish(vLog types.Log, txTime string, sender string) (p otTypes.PublishEvent, err error) {
	event := struct {
		MissionId   string
		RewardInWei *big.Int
		Data        string
	}{}

	contractAbi, err := abi.JSON(strings.NewReader(string(contracts.OpenTaskABI)))
	if err != nil {
		log.Fatal(err)
		return p, err
	}
	err = contractAbi.Unpack(&event, "Publish", vLog.Data)
	if err != nil {
		log.Fatal(err)
		return p, err
	}

	log.Printf("missionId: %s, reward: %s\n\n", event.MissionId, event.RewardInWei.String())

	p.Mission = event.MissionId
	p.Reward = event.RewardInWei
	p.Data = event.Data
	p.Block = vLog.BlockNumber
	p.Tx = vLog.TxHash.String()
	p.TxTime = txTime
	p.Publisher = sender
	return
}

func Solve(vLog types.Log, txTime string, sender string) (s otTypes.SolveEvent, err error) {
	event := struct {
		SolutionId string
		MissionId  string
		Data       string
	}{}

	contractAbi, err := abi.JSON(strings.NewReader(string(contracts.OpenTaskABI)))
	if err != nil {
		log.Fatal(err)
		return s, err
	}
	err = contractAbi.Unpack(&event, "Solve", vLog.Data)
	if err != nil {
		log.Fatal(err)
		return s, err
	}

	log.Printf("solutionId: %s, missionId: %s, data: %s\n", event.SolutionId, event.MissionId, event.Data)

	s.Solution = event.SolutionId
	s.Mission = event.MissionId
	s.Data = event.Data
	s.Block = vLog.BlockNumber
	s.Tx = vLog.TxHash.String()
	s.TxTime = txTime
	s.Solver = sender
	return
}

func Accept(vLog types.Log, txTime string, sender string) (a otTypes.AcceptEvent, err error) {
	event := struct {
		SolutionId string
	}{}

	contractAbi, err := abi.JSON(strings.NewReader(string(contracts.OpenTaskABI)))
	if err != nil {
		log.Fatal(err)
		return a, err
	}
	err = contractAbi.Unpack(&event, "Accept", vLog.Data)
	if err != nil {
		log.Fatal(err)
		return a, err
	}

	log.Printf("solutionId: %s\n", event.SolutionId)

	a.Solution = event.SolutionId
	a.Block = vLog.BlockNumber
	a.Tx = vLog.TxHash.String()
	a.TxTime = txTime
	_ = sender // NOT record
	return
}

func Reject(vLog types.Log, txTime string, sender string) (r otTypes.RejectEvent, err error) {
	event := struct {
		SolutionId string
	}{}

	contractAbi, err := abi.JSON(strings.NewReader(string(contracts.OpenTaskABI)))
	if err != nil {
		log.Fatal(err)
		return r, err
	}
	err = contractAbi.Unpack(&event, "Reject", vLog.Data)
	if err != nil {
		log.Fatal(err)
		return r, err
	}

	log.Printf("solutionId: %s\n", event.SolutionId)

	r.Solution = event.SolutionId
	r.Block = vLog.BlockNumber
	r.Tx = vLog.TxHash.String()
	r.TxTime = txTime
	_ = sender
	return
}

func Confirm(vLog types.Log, txTime string, sender string) (c otTypes.ConfirmEvent, err error) {
	log.Println("Confirm")
	event := struct {
		SolutionId    string
		ArbitrationId string
	}{}

	contractAbi, err := abi.JSON(strings.NewReader(string(contracts.OpenTaskABI)))
	if err != nil {
		log.Fatal(err)
		return c, err
	}
	err = contractAbi.Unpack(&event, "Confirm", vLog.Data)
	if err != nil {
		log.Fatal(err)
		return c, err
	}

	log.Printf("solutionId: %s, missionId: %s\n", event.SolutionId, event.ArbitrationId)

	c.Solution = event.SolutionId
	c.Arbitration = event.ArbitrationId
	c.Block = vLog.BlockNumber
	c.Tx = vLog.TxHash.String()
	c.TxTime = txTime
	return
}
