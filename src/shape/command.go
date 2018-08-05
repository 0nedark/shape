package shape

import (
	"encoding/json"
	"io/ioutil"
	"regexp"
	"strconv"

	"github.com/urfave/cli"

	log "github.com/sirupsen/logrus"
)

type keyValuePair = map[string]interface{}

var input keyValuePair
var file string
var mutable bool

// Action defines the shape command
func Action(c *cli.Context) error {
	input = keyValuePair{}
	file = c.String("file")
	mutable = c.Bool("mutable")
	if err := load(); err != nil {
		return err
	}

	if err := update(c.Args()); err != nil {
		return err
	}

	return save()
}

func load() error {
	raw, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(raw, &input); err != nil {
		return err
	}

	return nil
}

func parseValue(value string) interface{} {
	number, err := strconv.Atoi(value)
	if err != nil {
		return value
	}

	return number
}

func update(args cli.Args) error {
	r := regexp.MustCompile("(?P<key>.+)=(?P<value>[[:alnum:]]+)")
	for _, arg := range args {
		if matches := r.FindStringSubmatch(arg); len(matches) > 0 {
			key := matches[1]
			value := parseValue(matches[2])
			if _, ok := input[key]; ok || mutable {
				input[key] = value
			} else {
				log.WithFields(log.Fields{
					"Key":   key,
					"Value": value,
				}).Warn("Argument not in the original file ignored")
			}
		} else {
			log.WithField("Argument", arg).Warn("Invalid argument provided")
		}
	}

	return nil
}

func save() error {
	raw, err := json.Marshal(input)
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile(file, raw, 0); err != nil {
		return err
	}

	log.WithField("File", input).Info("File updated")
	return nil
}
