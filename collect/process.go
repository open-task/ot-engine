package collect

import (
	"github.com/ethereum/go-ethereum/common"
	"fmt"
	"math/big"
	"log"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"strings"
	openTask "github.com/xyths/ot-engine/contracts"
)

func Publish(topics []common.Hash, data []byte) {
	event := struct {
		MissionId string
		RewardInWei *big.Int
	}{}

	contractAbi, err := abi.JSON(strings.NewReader(string(openTask.OpenTaskABI)))
	if err != nil {
		log.Fatal(err)
	}
	err = contractAbi.Unpack(&event, "Publish", data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("missionId: %s\n", event.MissionId)
	fmt.Printf("rewardInWei: %v\n", event.RewardInWei.String())
}

func Solve(topics []common.Hash, data []byte) {
	event := struct {
		SolutionId string
		MissionId string
		Data string
	}{}

	contractAbi, err := abi.JSON(strings.NewReader(string(openTask.OpenTaskABI)))
	if err != nil {
		log.Fatal(err)
	}
	err = contractAbi.Unpack(&event, "Solve", data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("solutionId: %s, missionId: %s, rewardInWei: %s\n", event.SolutionId,event.MissionId, event.Data)
}

func Accept(topics []common.Hash, data []byte) {
	event := struct {
		SolutionId string
	}{}

	contractAbi, err := abi.JSON(strings.NewReader(string(openTask.OpenTaskABI)))
	if err != nil {
		log.Fatal(err)
	}
	err = contractAbi.Unpack(&event, "Accept", data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("solutionId: %s\n", event.SolutionId)
}

func Reject(topics []common.Hash, data []byte) {
	event := struct {
		SolutionId string
	}{}

	contractAbi, err := abi.JSON(strings.NewReader(string(openTask.OpenTaskABI)))
	if err != nil {
		log.Fatal(err)
	}
	err = contractAbi.Unpack(&event, "Reject", data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("solutionId: %s\n", event.SolutionId)
}

func Confirm(topics []common.Hash, data []byte) {
	fmt.Println("Confirm")
	event := struct {
		SolutionId string
		ArbitrationId string
	}{}

	contractAbi, err := abi.JSON(strings.NewReader(string(openTask.OpenTaskABI)))
	if err != nil {
		log.Fatal(err)
	}
	err = contractAbi.Unpack(&event, "Confirm", data)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("solutionId: %s, missionId: %s\n", event.SolutionId,event.ArbitrationId)
}
