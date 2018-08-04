package shape

import (
	"encoding/json"
	"io/ioutil"
	"regexp"

	"github.com/urfave/cli"

	log "github.com/sirupsen/logrus"
)

type file = map[string]interface{}

var input file
var path string
var mutable bool

func Action(c *cli.Context) error {
	input = file{}
	path = c.String("path")
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
	raw, err := ioutil.ReadFile(path)
	if err != nil {
		return err
	}

	if err = json.Unmarshal(raw, &input); err != nil {
		return err
	}

	return nil
}

func update(args cli.Args) error {
	r := regexp.MustCompile("(?P<key>.+)=(?P<value>[[:alnum:]]+)")
	for _, arg := range args {
		if matches := r.FindStringSubmatch(arg); len(matches) > 0 {
			key := matches[1]
			value := matches[2]
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
	log.WithField("File", input).Info("File updated")

	return nil
}
