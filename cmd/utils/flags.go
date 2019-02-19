package utils

import "gopkg.in/urfave/cli.v2"

var (
	ConfigFlag = &cli.StringFlag{
		Name:    "config",
		Aliases: []string{"c"},
		Value:   "config.json",
		Usage:   "Load configuration from `file`",
	}

	FromFlag = &cli.Int64Flag{
		Name:    "from",
		Aliases: []string{"f"},
		Usage:   "download `FROM` this block number",
	}
	ToFlag = &cli.Int64Flag{
		Name:    "to",
		Aliases: []string{"t"},
		Usage:   "download `TO` this block number",
	}
)
