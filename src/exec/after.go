package exec

import (
	"github.com/0nedark/shape/src/config"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func after(c *cli.Context) error {
	for _, arg := range arguments {
		delete(input, arg.Key)
	}

	if err := config.SaveSilent(c.String("shape"), input); err != nil {
		log.WithError(err).Fatal("Failed to save shape file")
	}

	return nil
}
