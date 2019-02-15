package node

import (
	"math/big"
)

type Config struct {
	Server    string   `json:"server"`
	Contract  string   `json:"contract""`
	FromFlag  bool     `json:"from_flag"`
	FromBlock *big.Int `json:"from_block"`
	ToFlag    bool     `json:"to_flag"`
	ToBlock   *big.Int `json:"to_block"`
}

type Node struct {
	Config *Config
}

var DefaultConfig = Config{
	Server:    "http://localhost",
	Contract:  "",
	FromFlag:  false,
	FromBlock: big.NewInt(0),
	ToFlag:    false,
	ToBlock:   big.NewInt(0),
}

func New(conf *Config) (*Node, error) {
	return &Node{
		Config: conf,
	}, nil
}
