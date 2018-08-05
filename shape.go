package main

import (
	"os"

	"github.com/0nedark/shape/src/shape"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "shape"
	app.Usage = "modify config files with one command line tool."
	app.Version = "1.1.0"
	app.Flags = []cli.Flag{
		cli.StringFlag{
			Name:  "file, f",
			Usage: "file to be shaped",
		},
		cli.BoolFlag{
			Name:  "mutable, m",
			Usage: "allow mutation of the original file",
		},
	}
	app.Action = shape.Action

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
