package main

import (
	"encoding/json"
	"fmt"
	"github.com/xyths/ot-engine/engine"
	"gopkg.in/urfave/cli.v2"
	"log"
	"math/big"
	"os"

	"github.com/xyths/ot-engine/cmd/utils"
	"github.com/xyths/ot-engine/node"
)

type HTTPConfig struct {
	Port string `json:port`
}

type DatabaseConfig struct {
	Host     string `json:"host"`
	Port     string `json:port`
	User     string `json:user`
	Password string `json:"password"`
	Database string `json:"database"`
}

type Config struct {
	Database DatabaseConfig `json:"database"`
	Http     HTTPConfig     `json:http`
	Node     node.Config    `json:"node"`
}

func (c Config) DSN() (dsn string) {
	dsn = c.Database.User + ":" + c.Database.Password + "@"
	if c.Database.Host != "" {
		dsn += "tcp(" + c.Database.Host
		if c.Database.Port != "" {
			dsn += ":" + c.Database.Port
		}
		dsn += ")"
	}

	dsn += "/" + c.Database.Database
	return dsn
}

func (c Config) Address() string {
	return ":" + c.Http.Port
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
	stack, err := node.New(&cfg.Node)
	if err != nil {
		log.Fatalf("Failed to create the stack: %v", err)
	}

	return stack
}

func makeConfigEngine(ctx *cli.Context) *engine.OtEngineServer {
	cfg := LoadConfig(ctx)
	eCfg := engine.Config{Address: cfg.Address(), DSN: cfg.DSN()}
	stack, err := engine.New(&eCfg)
	if err != nil {
		log.Fatalf("Failed to create the stack: %v", err)
	}

	return stack
}
