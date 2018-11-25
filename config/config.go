package config

import (
	"os"
	"fmt"
	"encoding/json"
)

type Config struct {
	Database struct {
		Host     string `json:"host"`
		Port     string `json:port`
		User     string `json:user`
		Password string `json:"password"`
		Database string `json:"database"`
	} `json:"database"`
	Network  string `json:"network"`
	Contract string `json:"contract"`
}

func LoadConfig(filename string) (config Config, err error) {
	configFile, err := os.Open(filename)
	defer configFile.Close()
	if err != nil {
		fmt.Println(err.Error())
	}
	jsonParser := json.NewDecoder(configFile)
	jsonParser.Decode(&config)
	return config, err
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
