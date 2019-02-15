package utils

import "gopkg.in/urfave/cli.v2"

var (
	ConfigFlag = &cli.StringFlag{
		Name:    "config",
		Aliases: []string{"c"},
		Value:   "config.json",
		Usage:   "Load configuration from `file`",
	}
)
