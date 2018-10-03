package exec

import (
	"bytes"
	"os"
	"os/exec"
	"regexp"

	"github.com/0nedark/shape/src/config"
	log "github.com/sirupsen/logrus"
	"github.com/urfave/cli"
)

func action(c *cli.Context) error {
	cmd := exec.Command("sh", "-c", c.String("command"))
	var out bytes.Buffer
	if c.IsSet("regexp") {
		cmd.Stdout = &out
	} else {
		cmd.Stdout = os.Stdout
	}

	err := cmd.Run()
	if err != nil {
		log.WithError(err).Error("Failed to run the specified command")
		return nil
	}

	result := make(map[string]interface{})
	r := regexp.MustCompile(c.String("regexp"))
	matches := r.FindStringSubmatch(out.String())
	for i, name := range r.SubexpNames() {
		if i != 0 && name != "" {
			result[name] = matches[i]
		}
	}

	if err := config.SaveSilent("/captures.json", result); err != nil {
		log.WithError(err).Error("Failed to create the regexp captures file")
	}

	return nil
}
