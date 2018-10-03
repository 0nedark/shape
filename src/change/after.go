package change

import (
	"github.com/0nedark/shape/src/config"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func after(c *cli.Context) error {
	if err := config.Save(c.String("shape"), input); err != nil {
		log.WithError(err).Fatal("Failed to save shape file")
	}

	return nil
}
