package collect

import (
	"github.com/ethereum/go-ethereum/core/types"
	"fmt"
	"math/big"
	"log"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
	openTask "github.com/xyths/ot-engine/contracts"
	. "github.com/xyths/ot-engine/types"
)

// 解析Publish日志。
func Publish(vLog types.Log) (p PublishEvent, err error) {
	event := struct {
		MissionId   string
		RewardInWei *big.Int
	}{}

	contractAbi, err := abi.JSON(strings.NewReader(string(openTask.OpenTaskABI)))
	if err != nil {
		log.Fatal(err)
		return p, err
	}
	err = contractAbi.Unpack(&event, "Publish", vLog.Data)
	if err != nil {
		log.Fatal(err)
		return p, err
	}

	fmt.Printf("missionId: %s, rewardInWei: %s\n\n", event.MissionId, event.RewardInWei.String())

	p.Mission = event.MissionId
	p.Reward = event.RewardInWei
	p.Block = vLog.BlockNumber
	p.Tx = vLog.TxHash.String()
	return p, err
}

func Solve(vLog types.Log) (s SolveEvent, err error) {
	event := struct {
		SolutionId string
		MissionId  string
		Data       string
	}{}

	contractAbi, err := abi.JSON(strings.NewReader(string(openTask.OpenTaskABI)))
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
	return s, err
}

func Accept(vLog types.Log) (a AcceptEvent, err error) {
	event := struct {
		SolutionId string
	}{}

	contractAbi, err := abi.JSON(strings.NewReader(string(openTask.OpenTaskABI)))
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
	return a, err
}

func Reject(vLog types.Log) (r RejectEvent, err error) {
	event := struct {
		SolutionId string
	}{}

	contractAbi, err := abi.JSON(strings.NewReader(string(openTask.OpenTaskABI)))
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
	return r, err
}

func Confirm(vLog types.Log) (c ConfirmEvent, err error) {
	fmt.Println("Confirm")
	event := struct {
		SolutionId    string
		ArbitrationId string
	}{}

	contractAbi, err := abi.JSON(strings.NewReader(string(openTask.OpenTaskABI)))
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
	return c, err
}
