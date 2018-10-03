package main

import (
	"os"

	"github.com/0nedark/shape/src/change"
	"github.com/0nedark/shape/src/exec"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "shape"
	app.Usage = "modify config files with one command line tool."
	app.Version = "1.2.0"

	app.Commands = []cli.Command{
		change.GetCommand(),
		exec.GetCommand(),
	}

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}
