package config

import (
	"regexp"

	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

type Argument struct {
	Key   string
	Value interface{}
}

func ParseArguments(args cli.Args, expression string, value func(v string) interface{}) []Argument {
	var arguments []Argument
	r := regexp.MustCompile(expression)
	for _, arg := range args {
		if matches := r.FindStringSubmatch(arg); len(matches) > 0 {
			arguments = append(arguments, Argument{
				Key:   matches[1],
				Value: value(matches[2]),
			})
		} else {
			log.WithField("Argument", arg).Warn("Invalid argument provided")
		}
	}

	return arguments
}
