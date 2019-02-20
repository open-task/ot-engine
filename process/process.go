package process

import (
	"database/sql"
	"fmt"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/xyths/ot-engine/contracts"
	"github.com/xyths/ot-engine/database"
	. "github.com/xyths/ot-engine/decode"
	otTypes "github.com/xyths/ot-engine/types"
	"log"
	"math/big"
	"strings"
)

func ParseOTLog(vLog types.Log, txTime string, sender string, db *sql.DB) (err error) {
	switch vLog.Topics[0].String() {
	case PublishSigHash.Hex():
		fmt.Println("Publish")
		row, err1 := Publish(vLog, txTime, sender)
		if err1 != nil {
			return err1
		}
		fmt.Println(row)
		err1 = database.Publish(db, row)
		if err1 != nil {
			return err1
		}
	case SolveSigHash.Hex():
		fmt.Println("Solve")
		row, err1 := Solve(vLog, txTime, sender)
		if err1 != nil {
			return err1
		}
		fmt.Println(row)
		err1 = database.Solve(db, row)
		if err1 != nil {
			return err1
		}
	case AcceptSigHash.Hex():
		fmt.Println("Accept")
		row, err1 := Accept(vLog, txTime, sender)
		if err1 != nil {
			return err1
		}
		fmt.Println(row)
		err1 = database.Accept(db, row)
		if err1 != nil {
			return err1
		}
	case RejectSigHash.Hex():
		fmt.Println("Reject")
		row, err1 := Reject(vLog, txTime, sender)
		if err1 != nil {
			return err1
		}
		fmt.Println(row)
		err1 = database.Reject(db, row)
		if err1 != nil {
			return err1
		}
	case ConfirmSigHash.Hex():
		fmt.Println("Confirm")
		row, err1 := Confirm(vLog, txTime, sender)
		if err1 != nil {
			return err1
		}
		fmt.Println(row)
		err1 = database.Confirm(db, row)
		if err1 != nil {
			return err1
		}
	default:
		fmt.Println("UNKNOWN Event Log")
	}
	return
}

func Publish(vLog types.Log, txTime string, sender string) (p otTypes.PublishEvent, err error) {
	event := struct {
		MissionId   string
		RewardInWei *big.Int
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

	fmt.Printf("missionId: %s, reward: %s\n\n", event.MissionId, event.RewardInWei.String())

	p.Mission = event.MissionId
	p.Reward = event.RewardInWei
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

	fmt.Printf("solutionId: %s, missionId: %s, data: %s\n", event.SolutionId, event.MissionId, event.Data)

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

	fmt.Printf("solutionId: %s\n", event.SolutionId)

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

	fmt.Printf("solutionId: %s\n", event.SolutionId)

	r.Solution = event.SolutionId
	r.Block = vLog.BlockNumber
	r.Tx = vLog.TxHash.String()
	r.TxTime = txTime
	_ = sender
	return
}

func Confirm(vLog types.Log, txTime string, sender string) (c otTypes.ConfirmEvent, err error) {
	fmt.Println("Confirm")
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

	fmt.Printf("solutionId: %s, missionId: %s\n", event.SolutionId, event.ArbitrationId)

	c.Solution = event.SolutionId
	c.Arbitration = event.ArbitrationId
	c.Block = vLog.BlockNumber
	c.Tx = vLog.TxHash.String()
	c.TxTime = txTime
	return
}
