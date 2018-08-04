package main

import (
	"log"
	"os"

	"github.com/urfave/cli"
)

func main() {
	app := cli.NewApp()
	app.Name = "shape"
	app.Usage = "modify config files with one command line tool."
	app.Action = action

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func action(c *cli.Context) error {
	log.Print("Do stuff")
	return nil
}
