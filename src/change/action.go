package change

import (
	"strconv"

	"github.com/0nedark/shape/src/config"
	"github.com/urfave/cli"
)

func action(c *cli.Context) error {
	arguments := config.ParseArguments(
		c.Args(),
		"(?P<key>.+)=(?P<value>[[:alnum:]]*)",
		parseValue,
	)

	for _, arg := range arguments {
		config.UpdateInput(input, arg.Key, arg.Value)
	}

	return nil
}

func parseValue(value string) interface{} {
	switch value {
	case "true":
		return true
	case "false":
		return false
	}

	number, err := strconv.Atoi(value)
	if err != nil {
		return value
	}

	return number
}
