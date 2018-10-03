package config

import (
	log "github.com/sirupsen/logrus"
)

func UpdateInput(input Structure, key string, value interface{}) {
	if _, ok := input[key]; ok {
		log.WithFields(log.Fields{
			"Key":   key,
			"Value": value,
		}).Warn("Overwriting argument in the original file")
	}

	switch value {
	case "":
		delete(input, key)
	default:
		input[key] = value
	}
}
