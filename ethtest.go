package main

import (
	"fmt"
	"log"

	"github.com/ethereum/go-ethereum/ethclient"
)

func main() {
	client, err := ethclient.Dial("https://rinkeby.infura.io/v3/e17969db9bc94e75a474b3d3c5257a75")
	if err != nil {
		log.Fatal(err)
	} else {
		fmt.Println("we have a connection now.")
	}

	_ = client // we'll use this in the upcoming sections

	//address := common.HexToAddress("0xF562a7c51a158ae6E6170Ef7905af5d1cE43d24A") // Rinkeby
	//instance, err := openTask.NewOpenTask(address, client)
	//if err != nil {
	//	log.Fatal(err)
	//}
	//
	//fmt.Println("contract is loaded")
	//_ = instance
}