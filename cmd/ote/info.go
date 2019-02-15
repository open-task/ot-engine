package main

import "gopkg.in/urfave/cli.v2"

var (
	alex = &cli.Author{
		Name:  "Alexander Xing",
		Email: "alexanderxing@gmail.com",
	}
	Authors = []*cli.Author{alex,}
)

func Info(app *cli.App) {
	app.Copyright = "©2018 The DEC Foundation"
	app.Authors = Authors
}
