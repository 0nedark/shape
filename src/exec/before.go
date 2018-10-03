package exec

import (
	"github.com/0nedark/shape/src/config"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var input config.Structure
var arguments []config.Argument

func before(c *cli.Context) error {
	input = config.Structure{}
	if err := config.Load(c.String("shape"), input); err != nil {
		log.WithError(err).Fatal("Failed to load shape file")
	}

	arguments = config.ParseArguments(
		c.Args(),
		"(?P<key>.+)=(?P<value>[[:graph:]]+)",
		toStructure,
	)

	for _, arg := range arguments {
		config.UpdateInput(input, arg.Key, arg.Value)
	}

	if err := config.SaveSilent(c.String("shape"), input); err != nil {
		log.WithError(err).Fatal("Failed to save shape file")
	}

	return nil
}

func toStructure(v string) interface{} {
	input := config.Structure{}
	if err := config.Load(v, input); err != nil {
		log.WithError(err).Warn("Failed to load specified file")
	}

	return input
}
