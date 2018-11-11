package collect

import (
	"github.com/ethereum/go-ethereum/common"
	"fmt"
)

func Publish(topics []common.Hash) {
	fmt.Println("Publish")

}

func Solve(topics []common.Hash) {
	fmt.Println("Solve")

}

func Accept(topics []common.Hash) {
	fmt.Println("Accept")

}

func Reject(topics []common.Hash) {
	fmt.Println("Reject")

}
