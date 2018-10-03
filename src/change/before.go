package change

import (
	"github.com/0nedark/shape/src/config"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

var input config.Structure

func before(c *cli.Context) error {
	input = config.Structure{}
	if err := config.Load(c.String("shape"), input); err != nil {
		log.WithError(err).Fatal("Failed to load shape file")
	}

	return nil
}
