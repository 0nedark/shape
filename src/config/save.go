package config

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	log "github.com/sirupsen/logrus"
	yaml "gopkg.in/yaml.v2"
)

func SaveSilent(file string, input Structure) error {
	var marshal func(v interface{}) ([]byte, error)
	switch filepath.Ext(file) {
	case ".json":
		marshal = json.Marshal
	case ".yml", ".yaml":
		marshal = yaml.Marshal
	}

	raw, err := marshal(input)
	if err != nil {
		return err
	}

	if err = ioutil.WriteFile(file, raw, 0); err != nil {
		return err
	}

	return nil
}

func Save(file string, input Structure) error {
	err := SaveSilent(file, input)
	if err == nil {
		log.WithField("File", input).Info("File updated")
	}

	return err
}
