package collect

import (
	coreTypes "github.com/ethereum/go-ethereum/core/types"
	"fmt"
	"math/big"
	"log"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
	openTask "github.com/xyths/ot-engine/contracts"
	otTypes "github.com/xyths/ot-engine/types"
)

func Publish(vLog coreTypes.Log) (p otTypes.PublishEvent, err error) {
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

func Solve(vLog coreTypes.Log) (s otTypes.SolveEvent, err error) {
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
	s.Solution=event.SolutionId
	s.Mission = event.MissionId
	s.Data = event.Data
	return s, err
}

func Accept(vLog coreTypes.Log) (a otTypes.AcceptEvent, err error) {
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
	return a, err
}

func Reject(vLog coreTypes.Log) (r otTypes.RejectEvent, err error) {
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
	return r, err
}

func Confirm(vLog coreTypes.Log) (c otTypes.ConfirmEvent, err error) {
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
	return c, err
}
