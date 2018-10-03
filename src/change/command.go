package change

import "github.com/urfave/cli"

var command cli.Command

// GetCommand definition
func GetCommand() cli.Command {
	return command
}

func init() {
	command = cli.Command{
		Name:    "change",
		Aliases: []string{"c"},
		Usage:   "modify config files.",
		Before:  before,
		Action:  action,
		After:   after,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "shape, s",
				Usage: "Specifies the shape file",
			},
		},
	}
}
