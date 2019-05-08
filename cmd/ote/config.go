package main

import (
	"encoding/json"
	"fmt"
	"github.com/open-task/ot-engine/engine"
	"gopkg.in/urfave/cli.v2"
	"log"
	"math/big"
	"os"

	"github.com/open-task/ot-engine/cmd/utils"
	"github.com/open-task/ot-engine/node"
)

type Config struct {
	Engine  node.Engine  `json:"engine"`
	Node    node.Config  `json:"node"`
	Backend node.Backend `json:"backend"`
}

func (c Config) Address() string {
	return ":" + c.Engine.Http.Port
}

func DefaultConfig() (config Config) {
	config.Node = node.DefaultConfig
	return config
}

func (c *Config) Load(file string) (err error) {
	configFile, err := os.Open(file)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	err = json.NewDecoder(configFile).Decode(c)
	if err != nil {
		return err
	}
	if c.Node.FromBlock.Cmp(big.NewInt(0)) == 1 {
		c.Node.FromFlag = true
	}
	if c.Node.ToBlock.Cmp(big.NewInt(0)) == 1 {
		c.Node.ToFlag = true
	}
	return err
}

func LoadConfig(ctx *cli.Context) Config {

	// Load defaults.
	cfg := DefaultConfig()

	// Load config file.
	if file := ctx.String(utils.ConfigFlag.Name); file != "" {
		if err := cfg.Load(file); err != nil {
			log.Fatalf("%v", err)
		}
	}
	return cfg
}

func makeConfigNode(ctx *cli.Context) *node.Node {
	cfg := LoadConfig(ctx)
	stack := &node.Node{
		EthConfig:     &cfg.Node,
		EngineConfig:  &cfg.Engine,
		BackendConfig: &cfg.Backend,
	}
	return stack
}

func makeConfigEngine(ctx *cli.Context) *engine.OtEngineServer {
	cfg := LoadConfig(ctx)
	eCfg := engine.Config{
		Address: cfg.Address(),
		DSN:     cfg.Engine.Database.DSN(),
	}
	stack, err := engine.New(&eCfg)
	if err != nil {
		log.Fatalf("Failed to create the stack: %v", err)
	}

	return stack
}
