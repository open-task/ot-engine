package main

import (
	"fmt"
	"github.com/open-task/ot-engine/cmd/utils"
	"gopkg.in/urfave/cli.v2"
	"os"
	"path/filepath"
)

var app *cli.App

func init() {
	app = &cli.App{
		Name:    filepath.Base(os.Args[0]),
		Usage:   "The ote(OpenTask Engine) command line interface",
		Version: "0.4.12",
		Action:  ote,
	}
	Info(app)
	app.Commands = []*cli.Command{
		serveCommand,
		downloadCommand,
		listenCommand,
		timerCommand,
	}
	app.Flags = []cli.Flag{
		utils.ConfigFlag,
	}
}

func main() {
	if err := app.Run(os.Args); err != nil {
		fmt.Fprintln(os.Stderr, err)
		os.Exit(1)
	}
}

func ote(ctx *cli.Context) error {
	return nil
}
