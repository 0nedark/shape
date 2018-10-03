package exec

import "github.com/urfave/cli"

var command cli.Command

// GetCommand definition
func GetCommand() cli.Command {
	return command
}

func init() {
	command = cli.Command{
		Name:    "exec",
		Aliases: []string{"e"},
		Usage:   "Executes the specified command.",
		Before:  before,
		Action:  action,
		After:   after,
		Flags: []cli.Flag{
			cli.StringFlag{
				Name:  "command, c",
				Usage: "command to be executed",
			},
			cli.StringFlag{
				Name:  "shape, s",
				Usage: "Specifies the shape file",
			},
			cli.StringFlag{
				Name:  "regexp, r",
				Usage: "Regular expression used to capture the output of the command. All named groups will be stored in the captures.json file",
			},
		},
	}
}
