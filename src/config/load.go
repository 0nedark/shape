package config

import (
	"encoding/json"
	"io/ioutil"
	"path/filepath"

	yaml "gopkg.in/yaml.v2"
)

type Structure = map[string]interface{}

func Load(file string, input Structure) error {
	var unmarshal func(data []byte, v interface{}) error
	switch filepath.Ext(file) {
	case ".json":
		unmarshal = json.Unmarshal
	case ".yml", ".yaml":
		unmarshal = yaml.Unmarshal
	}

	raw, err := ioutil.ReadFile(file)
	if err != nil {
		return err
	}

	if err = unmarshal(raw, &input); err != nil {
		return err
	}

	return nil
}
